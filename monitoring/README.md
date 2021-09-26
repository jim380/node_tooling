# Instructions

## Prerequisites

- [Docker](https://docs.docker.com/engine/install/ubuntu/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Deployment

1. Create a `prometheus.yml` under `/prometheus`. Some sample yaml files are included in the folder. Remember to replace `NODE_IP` with your server IP and `PROXY_HOST` with your `:9100` reverse-proxied by NGINX   

2. (**If alerts are needed**) Replace `PROXY_HOST` with your `:9093` reverse-proxied by NGINX , and fill in `TELEGRAM_ADMIN:` and `TELEGRAM_TOKEN` under section `alertmanager-bot:` in `docker-compose.yml`. Tweak the rest of `docker-compose.yml` as you see fit

3. (**If alerts are needed**) Fill in `http://NODE_IP:8080` under section `webhook_configs:` in `/prometheus/alert_manager/alertmanager.yml` 

4. Start the contrainers
    Deploy the monitoring stack (Prometheus + Node Exporter + Alert Manager) and proxy (Nginx + MariaDB) stack
    ```
    $ docker-compose up -d --profile monitor --profile proxy 
    ```
    Deploy the monitoring stack, the proxy stack and the telegram bot 
    ```
    $ docker-compose up -d --profile monitor --profile proxy --profile bot
    ```
    Deploy individual containers 
    ```
    $ docker-compose up -d prometheus node-exporter alertmanager nginx-proxy-manager mariadb
    ```

5. Stop the Containers
    ```
    $ docker-compose down
    ```