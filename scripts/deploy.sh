#!/usr/bin/env bash

ln -sf ./build/package/Dockerfile .
ln -sf ./deployments/docker-compose.mysql.yaml docker-compose.yaml

docker-compose down
docker-compose up --build