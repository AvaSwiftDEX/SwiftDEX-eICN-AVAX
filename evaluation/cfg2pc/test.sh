#!/bin/bash

# 获取脚本所在绝对路径
SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)/c

# 使用示例
echo "脚本目录：$SCRIPT_DIR"