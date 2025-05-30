#!/bin/bash

set -e

echo "============================================================"
echo "                Creating Chain 1                             "
echo "============================================================"
avalanche blockchain create ChainSwiftDEX01 --evm --proof-of-authority --validator-manager-owner 0x8db97C7cEcE249c2b98bDC0226Cc4C2A57BF52FC --test-defaults --evm-chain-id 971006 --evm-token Token01

echo "============================================================"
echo "                Creating Chain 2                             " 
echo "============================================================"
avalanche blockchain create ChainSwiftDEX02 --evm --proof-of-authority --validator-manager-owner 0x8db97C7cEcE249c2b98bDC0226Cc4C2A57BF52FC --test-defaults --evm-chain-id 971008 --evm-token Token02

echo "============================================================"
echo "                Creating Chain 3                             "
echo "============================================================"
avalanche blockchain create ChainSwiftDEX03 --evm --proof-of-authority --validator-manager-owner 0x8db97C7cEcE249c2b98bDC0226Cc4C2A57BF52FC --test-defaults --evm-chain-id 971010 --evm-token Token03


