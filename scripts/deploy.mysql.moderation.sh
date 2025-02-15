#!/usr/bin/env bash

VOLUME_NAME="mysql_data"

docker volume inspect $VOLUME_NAME > /dev/null 2>&1 || docker volume create $VOLUME_NAME

ln -sf ./build/package/Dockerfile .
ln -sf ./build/package/Dockerfile.moderation-tui .
ln -sf ./deployments/docker-compose.mysql.moderation.yaml docker-compose.yaml

docker-compose up --build -d