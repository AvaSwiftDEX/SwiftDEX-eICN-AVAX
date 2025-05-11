# Analysis

Compare with 2pc

## How does it perform in different scales?

### latency for each phase

- issue, worker prepare (uncfm/cfm)
- worker prepare (uncfm/cfm), master receive (uncfm/cfm) *here, uncfm means shadowStore*
- master commit, worker commit (uncfm/cfm) *here, uncfm means shadowStore*

### latency for overall process

- issue, worker commit (cfm)

### tps

- tps

## What is the influence of fork probability?

change fp=[0.1, 0.3, 0.5]

evaluate only one transaction

### gas

- worker prepare (uncfm/cfm)
- master receive/commit
- worker commit

## How does ShadowLock Perform?

change write conflict rate. the number of cross-chain messages which write the same variable

### overall latency

- issue, worker commit (cfm)

### overall tps

- tps
