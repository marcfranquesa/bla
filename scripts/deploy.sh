#!/usr/bin/env bash

#!/usr/bin/env bash

RESET_VOLUMES=false

for arg in "$@"; do
    case $arg in
        --reset-volumes)
        RESET_VOLUMES=true
        shift
        ;;
    *)
        ;;
    esac
done

ln -sf ./build/package/Dockerfile .
ln -sf ./deployments/compose.yaml .

if [ "$RESET_VOLUMES" = true ]; then
    docker-compose down -v
else
    docker-compose down
fi

docker-compose up --build

