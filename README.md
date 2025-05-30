# SwiftDEX: A Swift, Atomic and Secure Cross-Chain Exchange Protocol

<!-- 

Abstract: what is SwiftDEX, core technologies, and efficiency, features (especially reproducibility, practicality)

A figure (finally to do it)

 -->

## Introduction

<!-- <u>**Slow Interoperations:**</u> -->

### Motivation: Slow Interoperation and Why

Traditional cross-chain protocols (Avalanche, Chainlink, Cosmos, etc.) face the High-Latency bottleneck.

<!-- <u>**Why Slow:**</u> -->

<!-- ### Why Slow: -->

It is a fact that, only when the cross-chain message is finalized (or with a extreme-high probability) on the source chain, the target chain could accept it as a valid message by verification mechnisms (like Avalanche's BLS-Aggregation, IBC's Light Client, and Zero-Knowledge, etc.). The finalization process spends at least one consensus epoch (more epochs for Bitcoin, Ethereum2.0, etc.). Essentially, message finalization means that the message could be verified by outside entities (also called Verifier). Missing finalization means missing security guaranteed by underlying decentralized blockchains. **In one word, the message is first finalized, then transmitted and finally verified.** It is the traditional cross-chain protocol abstract paradigm.

<!-- <u>**Core Idea with One-Word:**</u> -->

### SwiftDEX: Not Wait Finalization

**Core Idea:** Any message is always generated before finalization (It is the First Principle). Since the message has generated, then <u>*DO NOT wait finalization*</u>. Just asynchronize finalization, transmission, and verification.

<!-- <u>**SwiftDEX:**</u> -->

We propose <u>*SwiftDEX*</u> protocol, achieving <u>*swift*</u>, <u>*atomic*</u> cross-chain exchange without any <u>*security*</u> loss. It introduces the following cutting-edge technologies:

- *Post-Finality:* It is the fundamental sub-protocol of SwiftDEX. Post-Finality achieves the core idea of asynchronizing finalization, transmission and verification. With it, we could build the following Unstable Cross-chain message and Dual Lock.
- *UNSTABLE Cross-chain Message:* It is the **Pivotal, Critical, Core, Key** component to reduce atomic interoperation latency. However, it is also extremly easy while ignored by most researchers and developers. The message's unstability could reduce latency about 36%.
- *Dual Lock:* For some Non-Fungible Token/Asset/State, the Dual Lock could avoid dirty/repeatable/phantom read or write operations.

<!-- 
**Challenges:**

- (*Secure*) How to guarantee message validity?
  - We propose the one-way protocol 
- (*Atomic*) How to keep atomicity?
  - SwiftDEX, pro 
-->

<!-- <u>**Practical Considerations:**</u> -->

### Practical Considerations

- All in smart contracts without introducing any other trusted third-parties.
- Suitable for all kinds of blockchains (First Principle), without any hard-fork invasion.
- Compatible with existing verification mechanisms, like BLS-Aggregation, Light Client, Zero-Knowledge, etc.
  
<!-- <u>**Efficiency:**</u> -->

### Efficiency

Until now, we have conducted some evaluations in a local cross-chain network (including Ethereum, and Avalanche). The result shows that the latency has reduced 36%.

Next, we will conduct more evaluations in a geo-distributed cross-chain network with more complex DApps (like defi, nft, supply chain, etc.) and scenarios. And, we will also adapt SwiftDEX to other homogeneous/heterogeneous blockchains to emphasize its practicality.

## Implementation on Avalanche

We implement SwiftDEX in smart contracts by solidity for on-chain logic.

And the off-chain relayer (also called eICN) connecting different Avalanche C-Chains is also customized by golang.

Because this project is still a Proof-of-Concept which focuses on the core functionalities and technology, some interactive components (like UI, beautiful interactive operations, etc.) are listed as future work.

During research and development, we also produce some necessary avalanche tools, like multi-chain creation, deployment, and SDK (by golang), contract compiling&deployment (by golang) etc.

### Prerequisites

The contract is:

- written by solidity:v0.8.2,
- compiled by solc:v0.8.2, abigen:v1.15.0, protoc:v24.3,
- depoyed by ethclient:v1.15.0 with golang:v1.22.12。

The off-chain relayer (also called eICN), and metrics collector (for evaluation) are written by golang:v1.22.12。

The evaluation script is written by python:v3.8.10, and managed by pipenv.

#### pipenv

```bash
pip install pipenv
pipenv shell --python 'your python:v3.8.10 path'
pipenv install
```

#### solc

```bash
pip install solc-select
solc-select install 0.8.2
solc-select use 0.8.2
```

#### protoc

```bash
# Download latest release
PROTOC_ZIP=protoc-24.3-linux-x86_64.zip  # Change version and platform accordingly
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v24.3/$PROTOC_ZIP

# Unzip and move files
# sudo apt install zip unzip -y
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'

# Verify installation
protoc --version

# install the Go plugin for Protocol Buffers
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# For gRPC support:
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# add $PATH
export PATH="$PATH:$(go env GOPATH)/bin"
```

#### abigen

```bash
# install abigen in go-ethereum
git clone --depth 1 --branch v1.15.0 https://github.com/ethereum/go-ethereum.git
cd go-ethereum
make
make devtools
```

### Deploy and Execute

We create a basic local cross-chain network (based on Avalanche) with 3 C-Chains.

The blockchain creation&deployment, and contract compiling&deployment, and off-chain relayers setup all are in following scripts.

To achieve auto-mining in avalanche local-net, we also start a background process to send transactions continuously. (We do not find any command for auto-mining).

#### creation

```shell
bash evaluation/onestep/3/create_avalanche.sh
```

#### deploy C-Chains

In this script, we:

- deploy a avalanche C-Chain for each config
- update config with C-Chain's BlockchainId by querying "platform.getBlockchains" api of P-Chain
- start background process to send transaction continuously for auto-mining

```shell
bash evaluation/onestep/3/start_avalanche.sh
```

#### contracts, relayers and metrics collector

In this script, we:

- compile, deploy and initilize contracts by solc, abigen and ethClient
- start metrics collector for collectiong system metrics
- start a customized relayer for each C-Chain
- build a mini-network for all C-Chains' relayers

```shell
bash evaluation/onestep/3/start.sh 
```

> NOTE: at present, we can not open-source SwiftDEX's on-chain contract code, because the related reseach paper has not been published. Once the paper accepted, we will open-source the core contracts.

#### watching interoperation

Start another shell (named Shell-B), and run the following command to watch the whole cross-chain process.

```shell
bash ./analysis/run_analysis.sh 1 test evaluation/configs/config1.yaml,evaluation/configs/config2.yaml,evaluation/configs/config3.yaml false 
```

#### send cross-chain transaction

Start another shell, and run the following command to send a cross-chain tranasaction.

```shell
bash ./scripts/cross_send.sh --chain-ids="2,3" --value="10" --app-identifier="State" --app-value-id="0" --config=evaluation/configs/config1.yaml
```

After execution, in Shell-B, you will see the intermediate data is output continuously.
