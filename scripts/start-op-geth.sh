#!/bin/bash
#

GIT_BASE_DIR=$(git rev-parse --show-toplevel)
REPOS_DIR=$HOME/workspace/blockchain/blockchain-tools/setup/repos/
L2_DATA_DIR=$HOME/l2-op-node

cd $REPOS_DIR/op-geth

./build/bin/geth \
        --datadir $L2_DATA_DIR \
        --port 30313 \
        --http \
        --http.corsdomain="*" \
        --http.vhosts="*" \
        --http.addr=0.0.0.0 \
        --http.api=web3,debug,eth,txpool,net,engine,personal \
        --http.port=5055 \
        --ws \
        --ws.addr=0.0.0.0 \
        --ws.port=8546 \
        --ws.origins="*" \
        --ws.api=debug,eth,txpool,net,engine,personal \
        --syncmode=full \
        --gcmode=archive \
        --nodiscover \
        --maxpeers=0 \
        --networkid=42069 \
        --authrpc.vhosts="*" \
        --authrpc.addr=0.0.0.0 \
        --authrpc.port=8552 \
        --authrpc.jwtsecret=./jwt.txt \
        --rollup.disabletxpoolgossip=true
