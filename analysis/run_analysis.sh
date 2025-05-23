#! /bin/bash

# Store the integer parameter
total_number=$1
echo "total number: $total_number"

# Store the identifier parameter
identifier=$2
echo "identifier: $identifier"

# Store the config file paths
config_files=$3
echo "config files: $config_files"

# Store the extract parameter
extract=$4
echo "extract: $extract"

go run analysis/*.go --metrics-server-url=127.0.0.1:8090 --total-number=$total_number --identifier=$identifier --config-files=$config_files --extract=$extract
