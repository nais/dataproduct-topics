#!/usr/bin/env sh

if [ "$ONPREM" = "true" ]; then
    . /var/run/secrets/nais.io/kafka_serviceuser
fi

/app/dataproduct-topics
