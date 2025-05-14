#!/bin/bash

# 获取执行该脚本时的所有参数
params="$@"

mkdir -p SR2PC/Filter

solc --abi --bin --overwrite --optimize --optimize-runs 200 -o output ../SuperRunner-Contracts/contracts/2pc-master/lib/Filter.sol --allow-paths .
abigen --bin=output/Filter.bin --abi=output/Filter.abi --pkg=Filter --out=SR2PC/Filter/Filter.go
go run scripts/scripts.go deploy-lib-Filter $params

solc --abi --bin --overwrite --optimize --optimize-runs 200 -o output ../SuperRunner-Contracts/contracts/2pc-master/SR2PC.sol --allow-paths . --libraries SR2PC/lib
abigen --bin=output/SR2PC.bin --abi=output/SR2PC.abi --pkg=SR2PC --out=SR2PC/SR2PC.go
go run scripts/scripts.go deploy $params

solc --abi --bin --overwrite --optimize --optimize-runs 200 -o output ../SuperRunner-Contracts/contracts/2pc-master/app/State.sol --allow-paths .
abigen --bin=output/State.bin --abi=output/State.abi --pkg=AppState --out=SR2PC/AppState/State.go
go run scripts/scripts.go deploy-app-State $params

go run scripts/scripts.go register-app-State $params