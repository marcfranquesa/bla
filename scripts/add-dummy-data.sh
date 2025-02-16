#!/usr/bin/env bash

source .env

urls=(
  "https://github.com/marcfranquesa/bla"
  "https://marcfranquesa.com"
  "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
)

for url in "${urls[@]}"; do
  curl -X POST ${DOMAIN:-http://localhost:8080} -d "$url"
done
