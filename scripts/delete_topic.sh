#!/usr/bin/bash

# Delete topic from dev-fss and prod-fss
# Requires httpie
# Requires that the users .netrc file contains valid credentials for kafka-adminrest in both environments

topic="${1}"

echo "Deleting ${topic} from dev-fss ..."
http DELETE "https://kafka-adminrest.dev-fss.nais.io/api/v1/topics/${topic}"
echo

echo "Deleting ${topic} from prod-fss ..."
http DELETE "https://kafka-adminrest.prod-fss.nais.io/api/v1/topics/${topic}"
echo
