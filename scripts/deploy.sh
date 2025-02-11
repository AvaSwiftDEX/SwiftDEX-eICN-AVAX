#!/bin/bash

# 获取执行该脚本时的所有参数
params="$@"

go run scripts/scripts.go deploy $params