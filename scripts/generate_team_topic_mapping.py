#!/usr/bin/env python
import argparse
import glob
import os
import re
import shutil
import subprocess
import sys
import tempfile
from collections import defaultdict, Counter

# All third-party imports inside try-except block below
REQUIREMENTS = [
    "requests==2.28.1",
    "pyyaml==6.0",
    "alive-progress==2.4.1",
]

try:
    import requests
    import yaml
    from alive_progress import alive_bar
except ImportError:
    print("Missing requirements! Run the following command (inside a virtualenv):")
    print("pip install {}".format(" ".join(REQUIREMENTS)))
    sys.exit(100)

DEFAULT_OUTPUT = os.path.normpath(os.path.join(os.path.dirname(__file__), "..", "pkg/collector/team_topic_mapping.go"))
KAFKA_ADMIN_URL = "https://kafka-adminrest.{}.nais.io"
STREAM_TOPIC = re.compile(r".*-streams-[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}-.*")
COMMON_NAME = re.compile(r"CN=(.*?)(,|$)")
VALID_CLUSTERS = ("dev-fss", "prod-fss")  # Update TEMPLATE below when changing this

FILE_TEMPLATE = """\
package collector

var teamTopicMapping = map[string]map[string]string{
	"onprem-dev": {
%(dev-fss)s	    
	},
	"onprem-prod": {
%(prod-fss)s	    
	},
}
"""  # NOQA
TOPIC_TEMPLATE = '"{topic}": "{team}",'


def clone_vault_iac():
    workdir = tempfile.mkdtemp(prefix="team-topic-mapping")
    subprocess.check_call(["git", "clone", "git@github.com:navikt/vault-iac.git"], cwd=workdir)
    return workdir


def generate_service_user_mapping():
    print("Generating service user mappings ...")
    rootdir = clone_vault_iac()
    try:
        teamsdir = os.path.join(rootdir, "vault-iac", "terraform", "teams")
        mapping = {cn: {} for cn in VALID_CLUSTERS}
        for team in os.listdir(teamsdir):
            app_files = os.path.join(teamsdir, team, "apps", "*.y*ml")
            for app in glob.glob(app_files):
                with open(app) as fobj:
                    app_data = yaml.safe_load(fobj)
                    for cluster_name, cluster_data in app_data.get("clusters", {}).items():
                        if cluster_name not in VALID_CLUSTERS:
                            continue
                        for service_users in cluster_data.get("serviceuser", {}).values():
                            mapping[cluster_name].update({svc_name: team for svc_name in service_users})
        return mapping
    finally:
        shutil.rmtree(rootdir)


def get_topics(env):
    print(f"Getting topics in {env}...")
    url = KAFKA_ADMIN_URL.format(env) + "/api/v1/topics"
    resp = requests.get(url)
    resp.raise_for_status()
    data = resp.json()
    return [topic for topic in data["topics"] if not STREAM_TOPIC.match(topic)]


def get_service_users(env, topic):
    url = KAFKA_ADMIN_URL.format(env) + "/api/v1/topics/{}/groups".format(topic)
    resp = requests.get(url)
    resp.raise_for_status()
    data = resp.json()
    for group in data["groups"]:
        multiplier = 1
        if group["type"] == "PRODUCER":
            multiplier = 1000
        for member in group["members"]:
            if m := COMMON_NAME.search(member):
                yield multiplier, m.group(1)


def generate_topic_mapping(env):
    """Returns a map of topic -> list of service users for this env"""
    topics = get_topics(env)
    mappings = {}
    with alive_bar(len(topics), title=f"Generating topic mappings in {env} ...") as bar:
        for topic in topics:
            service_users = [t for m, t in sorted(get_service_users(env, topic), reverse=True)]
            mappings[topic] = service_users
            bar()
    return mappings


def generate_mapping(env, topic_mapping, service_user_mapping):
    """Returns a map of topic->team for this env"""
    mapping = {}
    with alive_bar(len(topic_mapping), title=f"Collating results for {env}") as bar:
        for topic, users in topic_mapping.items():
            for user in users:
                if team := service_user_mapping[env].get(user):
                    mapping[topic] = team
                    break
            bar()
    return mapping


def generate_go_code(mappings, output):
    params = {}
    for cluster, topics in mappings.items():
        lines = []
        for topic, team in topics.items():
            lines.append(TOPIC_TEMPLATE.format(topic=topic, team=team))
        params[cluster] = "\n".join(lines)
    generated = FILE_TEMPLATE % params
    output.write(generated)
    subprocess.check_call(["go", "fmt", output.name])


def main(output):
    service_user_mapping = generate_service_user_mapping()
    mappings = {}
    for env in ("dev-fss", "prod-fss"):
        topic_mapping = generate_topic_mapping(env)
        mappings[env] = generate_mapping(env, topic_mapping, service_user_mapping)
    generate_go_code(mappings, output)


if __name__ == '__main__':
    with open(DEFAULT_OUTPUT, "w") as fobj:
        main(fobj)
