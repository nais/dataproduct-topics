#!/usr/bin/env sh
if [ "$ONPREM" = "true" ]; then
    KAFKA_USERNAME="$(cat /var/run/secrets/nais.io/kafka_serviceuser/username)"
    KAFKA_PASSWORD="$(cat /var/run/secrets/nais.io/kafka_serviceuser/password)"
    export KAFKA_USERNAME KAFKA_PASSWORD
fi

/app/dataproduct-topics
