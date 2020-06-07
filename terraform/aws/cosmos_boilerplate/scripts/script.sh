#!/bin/bash

# sleep until instance is ready
until [[ -f /var/lib/cloud/instance/boot-finished ]]; do
  sleep 1
done

# install nginx
apt-get update
apt-get upgrade -y
sed -E -i 's/persistent_peers = \".*\"/persistent_peers = \"cf65f9b5f290e3eef62dbf721397bc2e3fd47ecd@wenchang-testnet2-alice.node.bandchain.org:26656,c5d042cca13c34ee5318a4e5fc49dd7950f0a1da@wenchang-testnet2-bob.node.bandchain.org:26656\"/' $HOME/.bandd/config/config.toml

# make sure nginx is started
# service nginx start
