#!/bin/sh

docker run -it --rm -p 4443:443/tcp -p 4443:443/udp nginx-http3
