#!/bin/sh

echo "Building nginx-http3 docker container..."
(
    cd nginx-http3/
    docker build -t nginx-http3 .
)
