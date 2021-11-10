#!/usr/bin/env sh
if [ "$ONPREM" = "true" ]; then
    KAFKA_USERNAME="$(cat /var/run/secrets/nais.io/kafka_serviceuser/username)"
    KAFKA_PASSWORD="$(cat /var/run/secrets/nais.io/kafka_serviceuser/password)"
    export KAFKA_USERNAME KAFKA_PASSWORD
fi

ls /var/run/secrets/gcp
head -n1 /var/run/secrets/gcp/sa.json

/app/dataproduct-topics
