# Analysis

## Setup

### Baseline

APP:

- Fungilbe(ERC20)
  - 2PL:
    - No
    - Only influenced by SDL
  - Prepare:
    - move some value from source account to intermediate account
  - Abort:
    - if value is not enough
  - Commit:
    - move value from intermediate account to the target account
  - Rollback:
    - move value from intermediate account back to source account
- Non-Fungible(State, ERC721)
  - 2PL:
    - Yes
    - Influenced by both SDL and 2PL
  - Prepare:
    - the state satisfies the app requirement
    - own 2PL
  - Abort:
    - if the state not satisfies the app requirement
    - if 2PC is occupied by other owner
  - Commit:
    - change state
    - release 2PL
  - ROLLBACK:
    - release 2PL

Evaluation Protocols:

- SuperRunner-ERC20
- SuperRunner-State
- 2PC-ERC20
- 2PC-State

Factors:

- First Level:
  - WCR (Write Conflict Rate): transaction number of writing on the same state
  - FP (Forged Probability): change forged block number for each height
  - FL (Finalization Latency): the finalization latency in an overall block epoch
- Second Level:
  - WCN (Worker Chain Number): change number of worker blockchains
  - TN (Transaction Number): change issuance number of transactions
  - SDL (Shadow Lock): change transaction number in each block

Observation:

- Latency
  - Overall latency
  - Each phase latency
- TPS
- gas
  - Each part gas
  - Amortized gas
- storage
- retry
  - Each phase average retry numbers
- abort rate
  - Q: improve or decrease?

## 1. How does it perform in different scales?

### 1.1. latency for each phase

- proc1.1: issue, worker prepare (uncfm)
- proc1.2: issue, worker prepare (cfm)
- proc2.1: worker prepare (uncfm), master receive (uncfm)
- proc2.2: worker prepare (uncfm), master receive (cfm)
- proc3.1: master commit, worker commit (uncfm)
- proc3.2: master commit, worker commit (cfm)

### 1.2. latency for overall process

- issue, worker commit (cfm)

### 1.3. tps

- tps

## 2. What is the influence of fork probability?

change fp=[0.1, 0.3, 0.5]

evaluate only one transaction

### 2.1. gas

- worker prepare (uncfm/cfm)
- master receive/commit
- worker commit

## 3. How does ShadowLock Perform?

change write conflict rate. the number of cross-chain messages which write the same variable

### 3.1. overall latency

- issue, worker commit (cfm)

### 3.2. overall tps

- tps
