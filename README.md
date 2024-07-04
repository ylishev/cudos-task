# Coding Assignment

#### NOTE: Instructions for how to build and run the solution is available at [SETUP.md](SETUP.md)

Create a tool using Golang which automatically withdraws rewards from the user’s Cudos blockchain
account and sends them to a Cudos address in the Cudos chain. These actions should run in the
background/on a schedule, and reports should be sent to the user with the relevant information.

To help on the development of the task, https://rpc.cudos.org:443 is the RPC endpoint of our mainnet,
https://rpc.testnet.cudos.org:443 is for testnet.
The `cudos-node` repo can be found at https://github.com/cudoventures/cudos-node/tree/production/mainnet.
To install it, please install the required dependencies and then run make install and make build.

#### Some useful commands:

Withdraw rewards
```bash
cudos-noded tx distribution withdraw-all-rewards -y --from test --gas auto --keyring-backend test --chain-id cudos-testnet-public-4 --gas-prices 5000000000000acudos --gas-adjustment 1.3 --node https://rpc.testnet.cudos.org:443
```

Send 1 CUDOS from test to <toAddress>
```bash
cudos-noded tx bank send test <toAddress> 1000000000000000000acudos -y --from test --gas auto --keyring-backend test --chain-id cudos-testnet-public-4 --gas-prices 5000000000000acudos --gas-adjustment 1.3 --node https://rpc.testnet.cudos.org:443
```


An example toAddress is cudos198qaeg4wkf9tn7y345dhk2wyjmm0krdm85uqwc.


Note that all flags are shared for all transactions, `-y --from test --gas auto --keyring-backend test --chain-id cudos-testnet-public-4 --gas-prices 5000000000000acudos --gas-adjustment 1.3 --node https://rpc.testnet.cudos.org:443`. The `chain-id` for mainnet is `cudos-1`.


Also note that `--from test` signals that the transaction is sent from an account that has been created using the tool and is named test. This might be one of the parameters that the user would like to customise/configure.


Some extra commands to help with potential testing. Please reach out if you want testnet tokens, or feel free to use the testnet faucet if you have Keplr installed.

Create a new account in the keyring
```bash
cudos-noded keys add test --keyring-backend test
```

Stake 1 CUDOS with the test account
```bash
cudos-noded tx staking delegate cudosvaloper198qaeg4wkf9tn7y345dhk2wyjmm0krdm68jp09 1000000000000000000acudos -y --from test --gas auto --keyring-backend test --chain-id cudos-testnet-public-4 --gas-prices 5000000000000acudos --gas-adjustment 1.3 --node https://rpc.testnet.cudos.org:443
```