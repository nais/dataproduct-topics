#!/usr/bin/env python

import argparse
import glob
import os
import re
import subprocess
import sys
import tempfile
from collections import defaultdict

REQUIREMENTS = [
    "requests==2.28.1",
    "pyyaml==6.0",
]

try:
    import requests
    import yaml
except ImportError:
    print("Missing requirements! Run the following command (inside a virtualenv):")
    print("pip install {}".format(" ".join(REQUIREMENTS)))
    sys.exit(100)

KAFKA_ADMIN_URL = "https://kafka-adminrest.{}.nais.io"
TEAM_CATALOG_URL = "https://teamkatalog-api.intern.nav.no"
COMMON_NAME = re.compile(r"CN=(.*?)(,|$)")
PERSON_ID = re.compile(r"[A-Z]\d{6}")
VALID_CLUSTERS = ("dev-fss", "prod-fss")  # Update TEMPLATE below when changing this
SESSION = requests.Session()


def init_vault_checkout():
    tmpdir = tempfile.gettempdir()
    workdir = os.path.join(tmpdir, "inspect_onprem_topic")
    vault_iac_dir = os.path.join(workdir, "vault-iac")
    if not os.path.isdir(vault_iac_dir):
        os.makedirs(workdir, exist_ok=True)
        subprocess.check_call(["git", "clone", "git@github.com:navikt/vault-iac.git"], cwd=workdir)
    else:
        subprocess.check_call(["git", "pull"], cwd=vault_iac_dir)
    teamsdir = os.path.join(workdir, "vault-iac", "terraform", "teams")
    return teamsdir


def generate_service_user_mappings(teamsdir, env):
    mapping = defaultdict(set)
    for team in os.listdir(teamsdir):
        app_files = os.path.join(teamsdir, team, "apps", "*.y*ml")
        for app in glob.glob(app_files):
            with open(app) as fobj:
                app_data = yaml.safe_load(fobj)
                for cluster_name, cluster_data in app_data.get("clusters", {}).items():
                    if cluster_name != env:
                        continue
                    for service_users in cluster_data.get("serviceuser", {}).values():
                        for svc_name in service_users:
                            mapping[svc_name].add(team)
    return mapping


def get_rest_data(url):
    resp = SESSION.get(url)
    resp.raise_for_status()
    return resp.json()


def get_topic_data(topic, env, service_user_mapping):
    get_groups(env, topic, service_user_mapping)
    get_offsets(env, topic)
    get_consumer_groups(env, topic)


def get_consumer_groups(env, topic):
    print(f"Getting consumer groups for {topic}")
    consumer_data = get_rest_data(KAFKA_ADMIN_URL.format(env) + "/api/v1/topics/{}/consumergroups".format(topic))
    for consumer in consumer_data.get("groups", []):
        line = [consumer["groupId"]]
        try:
            for offset_data in consumer["offsets"]:
                partition = offset_data["partition"]
                offset = offset_data["offset"]
                line.append(f"{partition}: {offset}")
        except KeyError:
            pass
        print("\n\t".join(line))


def get_offsets(env, topic):
    print(f"Getting offsets for {topic}")
    offset_data = get_rest_data(KAFKA_ADMIN_URL.format(env) + "/api/v1/topics/{}/offsets".format(topic))
    for num, partition in offset_data.get("partitions", {}).items():
        try:
            latest_offset = partition["latest"]["offset"]
        except KeyError:
            latest_offset = -1
        print(f"Latest offset in partition {num}: {latest_offset}")


def print_user_info(user_id):
    line = []
    try:
        user_data = get_rest_data(TEAM_CATALOG_URL + "/resource/{}".format(user_id))
        line.append("{} ({})".format(user_data.get("fullName"), user_id))
        teams_data = get_rest_data(TEAM_CATALOG_URL + "/member/membership/{}".format(user_id))
        for team in teams_data.get("teams", []):
            team_names = [team.get("name")]
            team_names.extend(t for t in team.get("naisTeams", []))
            line.append(", ".join(team_names))
        print("\n\t* ".join(line))
    except requests.RequestException:
        print(f"{user_id} - Unknown ident")


def get_groups(env, topic, service_user_mapping):
    print(f"Getting groups for {topic}")
    group_data = get_rest_data(KAFKA_ADMIN_URL.format(env) + "/api/v1/topics/{}/groups".format(topic))
    for group in group_data.get("groups", []):
        print("{name} ({type})".format(**group))
        for member in group.get("members", []):
            if m := COMMON_NAME.search(member):
                common_name = m.group(1)
                if m := PERSON_ID.match(common_name):
                    print_user_info(m.group(0))
                else:
                    teams = ", ".join(service_user_mapping.get(common_name, ["<unknown>"]))
                    print(f"{common_name}\n\t* {teams}")


def main(topic, env):
    teamsdir = init_vault_checkout()
    service_user_mapping = generate_service_user_mappings(teamsdir, env)
    get_topic_data(topic, env, service_user_mapping)


if __name__ == '__main__':
    parser = argparse.ArgumentParser(formatter_class=argparse.ArgumentDefaultsHelpFormatter)
    parser.add_argument("topic", help="Name of topic to inspect")
    parser.add_argument("env", help="Environment to inspect", choices=VALID_CLUSTERS, default="prod-fss", nargs="?")
    options = parser.parse_args()
    main(options.topic, options.env)
