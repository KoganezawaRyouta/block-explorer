# block-explorer

# About Containers
| container  name |  |
|:---|:---|
|btc-node|bitcoin node(bitcoin core)|
|bch-node|bitcoin cash node(bitcoin abc)|
|eth-node|ethereum node(parity ethereum)|
|web3|web3js|
|neo4j|GraphDB|

# GraphDB
## Block of Cypher Queries
### Parameters
```sql
:params { "blockhash": "00000000000003e690288380c9b27443b86e5a5ff0f8ed2473efbfdacb3014f3", "version": 536870912, "prevblock": "000000000000050bc5c1283dceaff83c44d3853c44e004198c59ce153947cbf4", "merkleroot": "64027d8945666017abaf9c1b7dc61c46df63926584bed7efd6ed11a6889b0bac", "timestamp": 1500514748, "bits": "1a0707c7", "nonce": 2919911776, "size": 748959, "txcount": 1926}
```

### Query
```sql
MERGE (block:block {hash:$blockhash})
CREATE UNIQUE (block)-[:coinbase]->(:output:coinbase)
SET
   block.size=$size,
   block.prevblock=$prevblock,
   block.merkleroot=$merkleroot,
   block.time=$timestamp,
   block.bits=$bits,
   block.nonce=$nonce,
   block.txcount=$txcount,
   block.version=$version

MERGE (prevblock:block {hash:$prevblock})
MERGE (block)-[:chain]->(prevblock)
```

### Parameters
```sql
:params {"txid":"2e2c43d9ef2a07f22e77ed30265cc8c3d669b93b7cab7fe462e84c9f40c7fc5c","hash":"00000000000003e690288380c9b27443b86e5a5ff0f8ed2473efbfdacb3014f3","i":1,"tx":{"version":1,"locktime":0,"size":237,"weight":840,"segwit":"0001"},"inputs":[{"vin":0,"index":"0000000000000000000000000000000000000000000000000000000000000000:4294967295","scriptSig":"03779c110004bc097059043fa863360c59306259db5b0100000000000a636b706f6f6c212f6d696e65642062792077656564636f646572206d6f6c69206b656b636f696e2f","sequence":4294967295,"witness":"01200000000000000000000000000000000000000000000000000000000000000000"}],"outputs":[{"vout":0,"index":"2e2c43d9ef2a07f22e77ed30265cc8c3d669b93b7cab7fe462e84c9f40c7fc5c:0","value":166396426,"scriptPubKey":"76a91427f60a3b92e8a92149b18210457cc6bdc14057be88ac","addresses":"14eJ6e2GC4MnQjgutGbJeyGQF195P8GHXY"},{"vout":1,"index":"2e2c43d9ef2a07f22e77ed30265cc8c3d669b93b7cab7fe462e84c9f40c7fc5c:1","value":0,"scriptPubKey":"6a24aa21a9ed98c67ed590e849bccba142a0f1bf5832bc5c094e197827b02211291e135a0c0e","addresses":""}]}
```

## Transaction of Cypher Queries

### Query
```sql
MATCH (block :block {hash:$hash})
MERGE (tx:tx {txid:$txid})
MERGE (tx)-[:inc {i:$i}]->(block)
SET tx += {tx}    

WITH tx
FOREACH (input in $inputs |
         MERGE (in :output {index: input.index}) 
         MERGE (in)-[:in {vin: input.vin, scriptSig: input.scriptSig, sequence: input.sequence, witness: input.witness}]->(tx)
         )
            
FOREACH (output in $outputs |
         MERGE (out :output {index: output.index})
         MERGE (tx)-[:out {vout: output.vout}]->(out)
         SET
             out.value= output.value,
             out.scriptPubKey= output.scriptPubKey,
             out.addresses= output.addresses
         FOREACH(ignoreMe IN CASE WHEN output.addresses <> '' THEN [1] ELSE [] END |
                 MERGE (address :address {address: output.addresses})
                 MERGE (out)-[:locked]->(address)
                 )
        )
```

## Sample Query
```sql
MATCH (inputs)-[:in]->(tx:tx)-[:out]->(outputs)
WHERE tx.txid='2e2c43d9ef2a07f22e77ed30265cc8c3d669b93b7cab7fe462e84c9f40c7fc5c'
OPTIONAL MATCH (inputs)-[:locked]->(inputsaddresses)
OPTIONAL MATCH (outputs)-[:locked]->(outputsaddresses)
OPTIONAL MATCH (tx)-[:inc]->(block)
RETURN inputs, tx, outputs, block, inputsaddresses, outputsaddresses
```

# web3js
## Debug Sample Code
```javascript
W3 = require('web3')
web3 = new W3(new W3.providers.HttpProvider('http://eth-node:18545'))
web3.eth.getNodeInfo().then(result => {console.log(result)})
web3.eth.getBlockNumber().then(result => {console.log(result)})
web3.eth.isSyncing().then(result => {console.log(result)})
```
