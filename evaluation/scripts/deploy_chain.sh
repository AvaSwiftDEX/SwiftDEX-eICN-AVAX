# !/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status

# get chain name, --num-nodes, --num-local-nodes, --http-port
chain_name=$1
num_nodes=$2
num_local_nodes=$3
http_port=$4

# deploy chain
# # `--num-nodes` node number of P-Chain
# # `--num-local-nodes` node number of C-Chain
# # `--http-port` http port of C-Chain, and it is also `ws`
avalanche blockchain deploy $chain_name --local --num-nodes $num_nodes --num-local-nodes $num_local_nodes --http-port $http_port --skip-update-check 

# check if the blockchain is deployed successfully
if [ $? -ne 0 ]; then
    echo "Failed to deploy the blockchain: $chain_name"
    exit 1
fi