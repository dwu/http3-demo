#!/bin/sh

echo "Building http3-client-rust..."
(
    cd http3-client-rust/
    cargo build
)

echo "Building qpack-decoder..."
(
    cd qpack-decoder/
    cargo build
)

echo "Building http-client-go..."
(
    cd http3-client-go/
    go build
)