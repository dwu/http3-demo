#!/bin/sh

QLOGDIR=qlog SSLKEYLOGFILE=keys ./http3-client/target/debug/http3-client http://localhost:4443
