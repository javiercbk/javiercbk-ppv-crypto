#!/bin/bash

set -euo pipefail

BITCOIN_CONF=/root/.bitcoin/bitcoin.conf
IP="$(ifconfig | grep -A 1 'eth0' | tail -1 | cut -d ':' -f 2 | cut -d ' ' -f 1)"

if [ ! -e "${BITCOIN_CONF}" ]; then
  mv /usr/local/bin/bitcoin.conf ${BITCOIN_CONF}
fi

if [ $# -eq 0 ]; then
  exec bitcoind --rpcbind=${IP}
else
  exec "$@"
fi