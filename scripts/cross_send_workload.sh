#!/bin/bash

# 获取执行该脚本时的所有参数
# 参数示例：
# --chain-ids="1,2,3" --app-identifier="State" --write-conflict-rate=30 --transaction-number=15 --config=evaluation/configs/config1.yaml
params="$@"

go run scripts/scripts.go cross-send-workload $params