#!/bin/bash

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)

./analysis/run_analysis.sh 1 test $SCRIPT_DIR/config1.yaml,$SCRIPT_DIR/config2.yaml,$SCRIPT_DIR/config3.yaml false