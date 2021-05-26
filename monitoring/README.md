# Instructions

## Prerequisites

- [Docker](https://docs.docker.com/engine/install/ubuntu/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Deployment

1. Create a `prometheus.yml` under `/prometheus`. Some sample yaml files are included in the folder

2. Run the setup script to populate your server IP into `prometheus.yml`
    ```
    $ ./setup.sh
    ```

3. Tweak `docker-compose.yml` if need be

4. Start the Contrainers
    ```
    $ docker-compose up -d
    ```

5. Stop the Containers
    ```
    $ docker-compose down
    ```