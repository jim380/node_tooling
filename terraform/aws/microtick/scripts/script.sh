#!/bin/bash

# sleep until instance is ready
until [[ -f /var/lib/cloud/instance/boot-finished ]]; do
  sleep 1
done

apt-get update
apt-get upgrade -y
sed -E -i 's/persistent_peers = \".*\"/persistent_peers = \"922043cd83af759dd5a0605b32991667e8fd4977@45.79.207.112:26656,48ba6c0308f8687083ede2012d2ad2c969d2ead8@microtick.spanish-node.es:6868\"/' $HOME/.microtick/mtd/config/config.toml
