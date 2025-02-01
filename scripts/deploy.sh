#!/usr/bin/env bash

ln -sf ./build/package/Dockerfile .
ln -sf ./deployments/compose.yaml .

docker-compose down -v
docker-compose up --build
