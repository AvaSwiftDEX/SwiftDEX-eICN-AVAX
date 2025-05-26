#!/bin/bash

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)/c

bash ./scripts/cross_send_workload.sh \
--chain-ids="2,3" \
--app-identifier="State" \
--write-conflict-rate=0 \
--transaction-number=50 \
--worker-cfg-files=$SCRIPT_DIR/config2.yaml,$SCRIPT_DIR/config3.yaml \
--config=$SCRIPT_DIR/config1.yaml