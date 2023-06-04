#!/bin/sh

(
    cd nginx-http3/
    docker build -t nginx-http3 .
)
