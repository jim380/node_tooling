#!/bin/bash
set -e
set -x

# load environment variables from .env file
source .env

docker stop $NODE_CONTAINER_NAME && docker rm $NODE_CONTAINER_NAME
docker-compose up reset-node
docker rm node-reset

# split SNAP_RPC by comma and take the first part
FIRST_RPC=$(echo $SNAP_RPC | cut -d',' -f1)
LATEST_HEIGHT=$(curl -s $FIRST_RPC/block | jq -r .result.block.header.height)
TRUST_HEIGHT=$((LATEST_HEIGHT - 2000))
TRUST_HASH=$(curl -s "$FIRST_RPC/block?height=$TRUST_HEIGHT" | jq -r .result.block_id.hash)

awk -v state_sync_enable="$STATE_SYNC_ENABLE" -v snap_rpc="$SNAP_RPC" -v trust_height="$TRUST_HEIGHT" -v trust_hash="$TRUST_HASH" '
  BEGIN {OFS=FS="="}
  /^\[statesync\]$/ {statesync=1}
  statesync && $1~/^enable/ {$2=" " state_sync_enable}
  statesync && $1~/^rpc_servers/ {$2=" \"" snap_rpc "\""}
  statesync && $1~/^trust_height/ {$2=" " trust_height}
  statesync && $1~/^trust_hash/ {$2=" \"" trust_hash "\""}
  {print $0}
' $NODE_HOME_DIR_LOCAL/config/config.toml > temp_config.toml && mv temp_config.toml $NODE_HOME_DIR_LOCAL/config/config.toml

# start node
docker-compose up -d node