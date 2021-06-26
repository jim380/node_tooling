# Instructions

## Prerequisites

- [Docker](https://docs.docker.com/engine/install/ubuntu/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Deployment

1. Create a `prometheus.yml` under `/prometheus`. Some sample yaml files are included in the folder. Remember to replace `NODE_IP` with your server IP

2. Fill in `NODE_IP:9093`, `TELEGRAM_ADMIN:` and `TELEGRAM_TOKEN` under section `alertmanager-bot:` in `docker-compose.yml`. Tweak the rest of `docker-compose.yml` as you see fir

3. Fill in `http://NODE_IP:8080` under section `webhook_configs:` in `/prometheus/alert_manager/alertmanager.yml` 

4. Start the contrainers
    Deploy the monitoring stack (Grafana + Prometheus + Node Exporter)
    ```
    $ docker-compose up -d
    ```
    Deploy the monitor stack and the alerting stack (alert manager + alerta + telegram bot)
    ```
    $ docker-compose --profile alert up -d
    ```

5. Stop the Containers
    ```
    $ docker-compose down
    ```

## Firewall Rules

Make sure the following ports are open:

- `9090` (prometheus)
- `9100` (node exporter)
- `9093` (alert manager)
- `8080` (telegram bot)