#!/usr/bin/env bash

VOLUME_NAME="bla_data"

docker volume inspect $VOLUME_NAME > /dev/null 2>&1 || docker volume create $VOLUME_NAME

ln -sf ./build/Dockerfile .
ln -sf ./build/Dockerfile.moderation-tui .
ln -sf ./deployments/docker-compose.mysql.moderation.yaml docker-compose.yaml

docker-compose up --build -d
