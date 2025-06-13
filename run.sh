#!/bin/bash
export $(grep -v '^#' env.properties | xargs)

if [ "$1" = "prod" ]; then
    # Start the adapter in production mode
    docker-compose -f ./docker-compose.prod.yaml up --build
elif [ "$1" = "nolive" ]; then
    # Start the adapter in nolive mode
    docker-compose -f ./docker-compose.nolive.yaml up --build
else
    # Start the adapter in development mode
    docker-compose -f ./docker-compose.dev.yaml up $1 $2
fi
