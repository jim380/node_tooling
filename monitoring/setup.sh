#!/usr/bin/env bash

set -e

unset NODE_IP
unset YAML_PATH

NODE_IP=$(ip route get 8.8.8.8 | sed -n '/src/{s/.*src *\([^ ]*\).*/\1/p;q}')
YAML_PATH='./prometheus/prometheus.yml'


if [ ! -f "$YAML_PATH" ]; then
    echo "prometheus.yml is missing at $YAML_PATH."
    exit 1
fi

cp "$YAML_PATH" './prometheus/prometheus.yml.edited'

sed -i "s/NODE_IP/$NODE_IP/g" './prometheus/prometheus.yml.edited'