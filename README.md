# Executor AVS

## Generate Go code from ABI

`go install github.com/ethereum/go-ethereum/cmd/abigen@latest` - to be able to generate Go code from abi files

## Run

Use Holesky Testnet

```bash
./bin/operator register \
    --node-url wss://silent-tame-seed.ethereum-holesky.quiknode.pro/a09b2aafbc9447b172c9964f3ac40c85edf5fd6a \
    --contract-addr 0xA1554918d4C6Ee86BDfA45a94937565FD3C35D00 \
    --private-key 67690922275186342153160243748991894478440812240329176884047813007980001353474
```

```bash
./bin/operator unregister \
    --node-url wss://silent-tame-seed.ethereum-holesky.quiknode.pro/a09b2aafbc9447b172c9964f3ac40c85edf5fd6a \
    --contract-addr 0xA1554918d4C6Ee86BDfA45a94937565FD3C35D00 \
    --private-key 67690922275186342153160243748991894478440812240329176884047813007980001353474
```

```bash
./bin/operator run \
    --node-url wss://silent-tame-seed.ethereum-holesky.quiknode.pro/a09b2aafbc9447b172c9964f3ac40c85edf5fd6a \
    --contract-addr 0xA1554918d4C6Ee86BDfA45a94937565FD3C35D00 \
    --private-key 67690922275186342153160243748991894478440812240329176884047813007980001353474
```

```bash
./bin/operator arrange \
    --node-url wss://silent-tame-seed.ethereum-holesky.quiknode.pro/a09b2aafbc9447b172c9964f3ac40c85edf5fd6a \
    --contract-addr 0xA1554918d4C6Ee86BDfA45a94937565FD3C35D00 \
    --private-key 67690922275186342153160243748991894478440812240329176884047813007980001353474
```
