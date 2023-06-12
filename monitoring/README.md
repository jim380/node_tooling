# Instructions

## Prerequisites

- [Docker](https://docs.docker.com/engine/install/ubuntu/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Deployment

### Node

#### Configuration
1. Initialize the node's working directory. In this tutorial we will use `gaiad` as an example:
    ```
    $ docker run -it -v ~/.gaia:/root/.gaia <img> gaiad init <moniker> --chain-id <chain_id>
    $ sudo chown $(whoami) ~/.gaia -R
    ```
2. Download genesis, add peers and tweak other configs as needed in `~/.gaia/config/config.toml` and `~/.gaia/config/app.toml`
3. Change the image name, ports, binding volume or whatever you need to change in `docker-compose.yml`

#### Fire Up
```
$ docker-compose up -d node
```

### Monitoring Stack

#### Configuration

##### Prometheus & Node Exporter
1. Create a `prometheus.yml` under `/prometheus`
  A couple of sample files have been provided for your reference.
2. Replace `URL_METRIC` with your `:26660` reverse-proxied by NGINX
3. Replace `URL_NODE` with your `:9100` reverse-proxied by NGINX.

##### Cendermint
Learn more about the project [here](https://github.com/jim380/Cendermint).

1. In `/prometheus/prometheus.yml` replace `SERVER_IP` with your server IP
2. In `/Cendermint/config.env` fill out the env variables according to your setup

##### Alert Manager & Telegram Bot

Only if you need alerts.

1. In `docker-compose.yml` replace `URL_ALERT` with your `:9093` reverse-proxied by NGINX
2. Fill in `TELEGRAM_ADMIN` and `TELEGRAM_TOKEN`
3. Tweak the rest as you see fit
4. In `/prometheus/alert_manager/alertmanager.yml` replace `URL_BOT` with your `:8080` reverse-proxied by NGINX

#### Fire Up
- Deploy the whole stack
    ```
    $ docker-compose --profile monitor --profile proxy --profile bot up -d
    ```
- Deploy only the proxy stack
    ```
    $ docker-compose --profile proxy up -d
    ```
- Deploy individual containers 
    ```
    $ docker-compose up -d app db alertmanager-bot cendermint
    ```

## Handy Commands
- View all running containers
    ```
    $ docker-compose ps
    ```
- Stop all containers
    ```
    $ docker-compose down
    ```
- Manage individual containers
    ```
    $ docker-compose stop <container>
    $ docker-compose start <container>
    $ docker-compose restart <container>
    ```