#!/bin/sh

xcaddy build \
    --with github.com/lucas-clemente/quic-go@v0.18.0 \
    --with github.com/rocketreferrals/caddy2-proxyprotocol=$(pwd)/../../caddy2-proxyprotocol \
    --with github.com/pires/go-proxyproto=$(pwd)/../../proxyproto

./caddy stop
./caddy start -config=./caddy.json
