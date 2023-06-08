#!/bin/sh

echo "Building http3-client..."
(
    cd http3-client/
    cargo build
)

echo "Building qpack-decoder..."
(
    cd qpack-decoder/
    cargo build
)