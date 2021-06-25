# Instructions

## Prerequisites

- [Docker](https://docs.docker.com/engine/install/ubuntu/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Deployment

1. Create a `prometheus.yml` under `/prometheus`. Some sample yaml files are included in the folder

2. Fill out the `email_configs:` section in `/prometheus/alert_manager/alertmanager.yml`

3. Run the setup script to populate your server IP into `prometheus.yml`
    ```
    $ ./setup.sh
    ```

4. Tweak `docker-compose.yml` if need be

5. Start the Contrainers
    ```
    $ docker-compose up -d
    ```

6. Stop the Containers
    ```
    $ docker-compose down
    ```