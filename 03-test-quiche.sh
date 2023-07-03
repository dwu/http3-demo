#!/bin/sh

QLOGDIR=qlog SSLKEYLOGFILE=keys ./http3-client-rust/target/debug/http3-client http://localhost:4443
