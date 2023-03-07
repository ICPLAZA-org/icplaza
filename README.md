<!--
parent:
  order: false
-->

<div align="center">
  <h1> ICPlaza </h1>
</div>

ICPlaza is a scalable, high-throughput Proof-of-Stake blockchain built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk/) which runs on top of [Tendermint Core](https://github.com/tendermint/tendermint) consensus engine.And Integrate [Ethermint](https://github.com/evmos/ethermint) whitch is a scalable and interoperable Ethereum library.

**Note**: Requires [Go 1.18+](https://golang.org/dl/)

## Installation
 run:

```bash
make install
```
1. Server recommended configuration:

CPU: 4H; memory 8G; hard disk: 200G ssd; system: centos7/ubuntu18

2. Create the /data/chain/icplaza/bin directory:
```bash
mkdir -p /data/chain/icplaza/bin
```
3. Download the icplaza executable file:
```bash
wget https://github.com/ICPLAZA-org/icplaza/releases/download/v6.0.2/icplazad
```
4. Go to the /data/chain/icplaza directory and execute the initialization command:
```bash
bin/icplazad --home data init icplaza --chain-id=icplaza_9000-1
```
5. Download the mainnet configuration file:

(1) Back up the genesis.json config.toml app.toml in the /data/chain/icplaza/data/config/ directory

(2) Download the mainnet configuration file and save it to /data/chain/icplaza/data/config/
```bash
wget https://github.com/ICPLAZA-org/mainnet/raw/main/config.tar.gz
```
6. Download the block data image to the /data/chain/icplaza/data/ directory
```bash
wget http://image.icplaza.pro/download/icplaza_chain_data.tar.gz
```
After the download is complete, directly replace the /data/chain/icplaza/data/data/ directory
7. Start the node /data/chain/icplaza/ directory and execute:
```bash
bin/icplazad --home=data start --pruning=default --json-rpc.api eth,txpool,personal,net,debug,web3 > log 2>&1 &
```

