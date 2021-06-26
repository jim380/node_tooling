# Instructions

## Prerequisites

- [Docker](https://docs.docker.com/engine/install/ubuntu/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Deployment

1. Create a `prometheus.yml` under `/prometheus`. Some sample yaml files are included in the folder. Remember to replace `NODE_IP` with your server IP

2. Fill in `NODE_IP:9093`, `TELEGRAM_ADMIN:` and `TELEGRAM_TOKEN` under section `alertmanager-bot:` in `docker-compose.yml`

3. Tweak `docker-compose.yml` if need be

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