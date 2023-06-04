#!/bin/sh

docker run -it --volume $(pwd)/qlog:/qlog --env QLOGDIR=/qlog --env SSLKEYLOGFILE=/qlog/key --rm ymuski/curl-http3 curl -kILv --http3 https://172.17.0.1:4443
