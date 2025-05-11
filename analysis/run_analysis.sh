#! /bin/bash

# Store the integer parameter
total_number=$1
echo "total number: $total_number"

go run analysis/*.go --metrics-server-url=127.0.0.1:8090 --total-number=$total_number