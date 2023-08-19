#!/bin/bash
#
# This script is used to setup the Optimism repos

GIT_BASE_DIR=$(git rev-parse --show-toplevel)
REPOS_DIR=$HOME/workspace/blockchain/blockchain-tools/setup/repos/
L1_DATA_DIR=$HOME/l1-op-node
L1_NODE_URL=http://localhost:8545
L1_BLOCK_TIME=3
L2_DATA_DIR=$HOME/l2-op-node

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd op-node
go run cmd/main.go genesis l2 \
  --deploy-config ../packages/contracts-bedrock/deploy-config/getting-started.json \
  --deployment-dir ../packages/contracts-bedrock/deployments/getting-started/ \
  --outfile.l2 genesis.json \
  --outfile.rollup rollup.json \
  --l1-rpc $L1_NODE_URL

yq -i '.config += {"l2GenesisBlockCronJobGasLimit": "0x5f5e100"}' genesis.json

openssl rand -hex 32 > jwt.txt

cp genesis.json $REPOS_DIR/op-geth/genesis.json
cp jwt.txt $REPOS_DIR/op-geth/jwt.txt


cd $REPOS_DIR/op-geth
rm -rf $L2_DATA_DIR
mkdir -p $L2_DATA_DIR
./build/bin/geth --datadir=$L2_DATA_DIR init genesis.json
