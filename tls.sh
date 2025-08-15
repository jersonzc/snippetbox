#!/bin/sh
set -e

if [ ! -f /tls/cert.pem ] || [ ! -f /tls/key.pem ]; then
    echo "Generating TLS certificates..."
    mkdir -p /tls
    go run /usr/local/go/src/crypto/tls/generate_cert.go \
        --rsa-bits=2048 --host=localhost
    mv cert.pem key.pem /tls
fi

exec "$@"
