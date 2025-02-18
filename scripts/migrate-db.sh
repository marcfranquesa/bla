#!/usr/bin/env bash

source .env

CONTAINER="bla-db"
MIGRATION_SQL_DIR="db/migrations"

SQL_FILES=$(find "$MIGRATION_SQL_DIR" -type f -name "*.sql" | sort)

for FILE in $SQL_FILES; do
    echo "Running: $FILE"
    docker exec -i "$CONTAINER" mysql --user="$DB_USER" --password="$DB_PASSWORD" "$DB_NAME" < "$FILE"
done
