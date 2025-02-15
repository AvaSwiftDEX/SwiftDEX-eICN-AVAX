# Evaluation

## Latency

### Overall

### Each Phase

#### 1. Cross-Chain Message Issue

The timestamp of cross-chain message issue is included in the coordinator chain's block.
- unfinalized timestamp
  - SendCMHash
- finalized timestamp
  - timestamp of the finalized block of ISSUE CM
  
#### 2. Cross-Chain Message Prepare on the target chain

The timestamp of cross-chain message receive is included in the target chain's block.
- unfinalized timestamp: the issue message is included in the block of the target chain
  - SendCMHash
- finalized timestamp: the issue message is finalized on the target chain
  - Finalized Block of PREPARE CM
- confirmed timestamp: the issue message's header is finalized on the target chain
  - Finalized Block of header (ISSUE CM)

#### 3. Cross-Chain Message Back on the source chain

The timestamp of cross-chain message back is included in the source chain's block.
- finalized timestamp: the back message is finalized in the block
  - Finalized Block of BACK CM
- confirmed timestamp: the header of the back message is finalized in the block
  - Finalized Block of header (BACK CM) 

#### 4. Cross-Chain Message Commit/Rollback on the coordinator chain

The timestamp of cross-chain message commit is included in the coordinator chain's block.
The back messages are all confirmed on the coordinator chain once the block header (including the prepare messages) is finalized on the coordinator chain.
- unfinalized timestamp: the timestamp of the block header is included in the block
  - SendCMHash
- finalized timestamp: the timestamp of the block header is finalized
  - Finalized Block of Commit CM

#### 5. Cross-Chain Message Commit/Rollback on the target chain

The timestamp of cross-chain message commit is finalized on the target chain.
- finalized timestamp: the timestamp of the commit message is finalized
  - Finalized Block of Commit CM
- confirmed timestamp: the header of the commit message is finalized
  - Finalized Block of header (COMMIT CM)
