#!/bin/bash
set -e
set -x

# load environment variables from .env file
source .env

sudo rm -rf $NODE_HOME_DIR_LOCAL

docker-compose up init-node
docker rm node-init
docker-compose up reset-node
docker rm node-reset

# download and replace genesis file
curl -s $GENEISIS_FILE_URL > $NODE_HOME_DIR_LOCAL/config/genesis.json

# configure config.toml
if [[ "$OSTYPE" == "darwin"* ]]; then
  # mac
  sed -i '' "s/seeds = \".*\"/seeds = \"$SEEDS\"/" $NODE_HOME_DIR_LOCAL/config/config.toml
  sed -i '' "s/indexer = \".*\"/indexer = \"$INDEXER\"/" $NODE_HOME_DIR_LOCAL/config/config.toml
else
  # linux
  sed -i "s/seeds = \".*\"/seeds = \"$SEEDS\"/" $NODE_HOME_DIR_LOCAL/config/config.toml
  sed -i "s/indexer = \".*\"/indexer = \"$INDEXER\"/" $NODE_HOME_DIR_LOCAL/config/config.toml
fi
awk -v rpc_laddr="$RPC_LADDR" '
  BEGIN {OFS=FS="="}
  /^\[rpc\]$/ {rpc=1; p2p=0}
  /^\[p2p\]$/ {p2p=1; rpc=0}
  rpc && $1~/^laddr/ {$2=" \"" rpc_laddr "\""}
  {print $0}
' $NODE_HOME_DIR_LOCAL/config/config.toml > temp_config.toml && mv temp_config.toml $NODE_HOME_DIR_LOCAL/config/config.toml

# configure app.toml
if [[ "$OSTYPE" == "darwin"* ]]; then
  # mac
  sed -i '' "s/minimum-gas-prices = \".*\"/minimum-gas-prices = \"$MIN_GAS_PRICE\"/" $NODE_HOME_DIR_LOCAL/config/app.toml
else
  # linux
  sed -i "s/minimum-gas-prices = \".*\"/minimum-gas-prices = \"$MIN_GAS_PRICE\"/" $NODE_HOME_DIR_LOCAL/config/app.toml
fi
awk -v api_enable="$API_ENABLE" -v api_address="$API_ADDR" -v grpc_enable="$GRPC_ENABLE" -v grpc_address="$GRPC_ADDR" -v grpc_web_enable="$GRPC_WEB_ENABLE" '
  BEGIN {OFS=FS="="}
  /^\[api\]$/ {api=1; grpc=0; grpc_web=0}
  /^\[grpc\]$/ {grpc=1; api=0; grpc_web=0}
  /^\[grpc-web\]$/ {grpc_web=1; api=0; grpc=0}
  api && $1~/enable/ {$2=" " api_enable}
  api && $1~/address/ {$2=" \"" api_address "\""}
  grpc && $1~/enable/ {$2=" " grpc_enable}
  grpc && $1~/^address/ {$2=" \"" grpc_address "\""}
  grpc_web && $1~/^enable/ {$2=" " grpc_web_enable}
  {print $0}
' $NODE_HOME_DIR_LOCAL/config/app.toml > temp_app.toml && mv temp_app.toml $NODE_HOME_DIR_LOCAL/config/app.toml

# start node
docker-compose up -d node

