#! /bin/bash

set -e

# Define an array where each element contains chain name, number of nodes, number of local nodes and port number
declare -a chains=(
    "ChainSwiftDEX01 2 4 60010 evaluation/configs/config1.yaml"
    "ChainSwiftDEX02 2 4 60020 evaluation/configs/config2.yaml"
    "ChainSwiftDEX03 2 4 60030 evaluation/configs/config3.yaml"
)

P_CHAIN_HOST="127.0.0.1"
P_CHAIN_PORT=9650
C_CHAIN_HOST="127.0.0.1"

for chain_config in "${chains[@]}"; do
    read -r chain_name _ _ _ _ <<< "$chain_config"
    if ! avalanche blockchain list | grep -q "$chain_name"; then
        echo "Blockchain $chain_name not found. Please create the blockchain($chain_name) first."
        exit 1
    fi
done

# Loop through each chain configuration and deploy
for chain_config in "${chains[@]}"; do
    # Split the chain config string into individual parameters
    read -r chain_name num_nodes num_local_nodes port config_path <<< "$chain_config"

    echo "=============================================="
    echo "Deploying blockchain: $chain_name"
    echo "Number of P-Chain nodes: $num_nodes"
    echo "Number of C-Chain nodes: $num_local_nodes" 
    echo "HTTP/WS Port: $port"
    echo "=============================================="
    
    # Execute deploy_chain.sh with the parameters
    bash evaluation/scripts/deploy_chain.sh "$chain_name" "$num_nodes" "$num_local_nodes" "$port"

    # Check if the blockchain is deployed successfully
    if [ $? -ne 0 ]; then
        echo "Failed to deploy the blockchain: $chain_name"
        exit 1
    fi

    # Update the config file
    go run evaluation/scripts/update_config/update_config.go --p_host "$P_CHAIN_HOST" --p_port "$P_CHAIN_PORT" --c_host "$C_CHAIN_HOST" --c_port "$port" --c_name "$chain_name" --config_path "$config_path"

    echo "Deployed the blockchain: $chain_name"
done

# python evaluation/scripts/update_config.py --p_host "127.0.0.1" --p_port 9650 --c_host "127.0.0.1" --c_port 60010 --c_name "ChainSwiftDEX01" --config_path "evaluation/test_configs/config1.yaml"

# Create an array to store config paths
declare -a config_paths=()

# Extract config paths from chains array
for chain_config in "${chains[@]}"; do
    read -r _ _ _ _ config_path <<< "$chain_config"
    config_paths+=("$config_path")
done

# Join config paths with semicolons
CONFIG_PATHS=$(IFS=';'; echo "${config_paths[*]}")
echo "CONFIG_PATHS: $CONFIG_PATHS"
screen -dmLS background_workload go run evaluation/scripts/background/background.go --config_paths "$CONFIG_PATHS"