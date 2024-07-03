package main

import (
	"context"
	"log"

	"cudos-task/cmd/cudos-task/cmd"
	"cudos-task/contract"
	"cudos-task/internal/shutdown"

	"github.com/spf13/viper"
)

func main() {
	// prepare dependencies
	vp := viper.New()
	outChannel := make(chan string, 1)
	ctx, cancel := context.WithCancel(context.Background())

	// handle graceful shutdown of the application
	shutdownReady := shutdown.NewShutdown(cancel)

	// build and run Cobra commands
	rootCmd := cmd.InitRootCmd(ctx, vp, shutdownReady, outChannel)
	executedCMD, err := rootCmd.ExecuteC()
	if err != nil {
		log.Printf("error executing command: %v\nUse --help for details\n", err)
		cancel()
	}

	// report back the scheduled operations to the user
	for flg := executedCMD.Flags().Lookup(contract.ReportBackFlagName); flg != nil && flg.Changed; {
		select {
		// or wait for Ctrl+C
		case <-ctx.Done():
			return
		case msg := <-outChannel:
			log.Printf("[schedule] %v\n", msg)
		}
	}
}

/*func withdrawc() (string, error) {
	sdk.DefaultPowerReduction = sdk.NewIntFromUint64(1000000000000000000)
	startup.SetConfig()

	cobraCmd := &cobra.Command{}
	cobraCmd.SetContext(context.Background())

	flagSet := cobraCmd.Flags()
	//	flagSet.String(flags.FlagFrom, keyringAccount, "")
	flagSet.Bool(flags.FlagSkipConfirmation, true, "")
	flagSet.String(cli.OutputFlag, "json", "")
	flagSet.String(flags.FlagHome, dataHomePath, "")
	// flagSet.String(flags.FlagKeyringDir, dataHomePath, "")
	//	flagSet.String(flags.FlagChainID, chainID, "")
	//	flagSet.String(flags.FlagKeyringBackend, keyringBackend, "")
	//	flagSet.String(flags.FlagNode, nodeAddress, "")
	flagSet.String(flags.FlagNote, memo, "")
	//	flagSet.Float64(flags.FlagGasAdjustment, gasAdj, "")
	//	flagSet.String(flags.FlagGas, gas, "")
	//	flagSet.String(flags.FlagGasPrices, gasPrices, "")

	aminoCodec := codec.NewLegacyAmino()
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txConfig := authtx.NewTxConfig(marshaler, authtx.DefaultSignModes)

	std.RegisterLegacyAminoCodec(aminoCodec)
	std.RegisterInterfaces(interfaceRegistry)
	simapp.ModuleBasics.RegisterLegacyAminoCodec(aminoCodec)
	simapp.ModuleBasics.RegisterInterfaces(interfaceRegistry)

	var outBuffer bytes.Buffer
	outWriter := bufio.NewWriter(&outBuffer)

	clientCtx, err := client.GetClientTxContext(cobraCmd)
	if err != nil {
		log.Fatalf("failed to obtain client context: %v", err)
	}

	clientCtx = clientCtx.
		WithCodec(marshaler).
		WithJSONCodec(marshaler).
		WithInterfaceRegistry(interfaceRegistry).
		WithTxConfig(txConfig).
		WithLegacyAmino(aminoCodec).
		WithInput(os.Stdin).
		WithOutput(outWriter).
		WithBroadcastMode(flags.BroadcastBlock).
		WithAccountRetriever(types.AccountRetriever{})
	//WithBroadcastMode(flags.BroadcastBlock).
	//WithHomeDir(dataHomePath)

	delAddr := clientCtx.GetFromAddress()
	queryClient := distribtypes.NewQueryClient(clientCtx)
	delValsRes, err := queryClient.DelegatorValidators(context.Background(), &distribtypes.QueryDelegatorValidatorsRequest{DelegatorAddress: delAddr.String()})
	if err != nil {
		log.Fatalf("failed to obtain the staking validators: %v", err)
	}

	validators := delValsRes.Validators
	// build multi-message transaction
	msgs := make([]sdk.Msg, 0, len(validators))
	for _, valAddr := range validators {
		val, err := sdk.ValAddressFromBech32(valAddr)
		if err != nil {
			log.Fatalf("failed to check the validator address: %v", err)
		}

		msg := distribtypes.NewMsgWithdrawDelegatorReward(delAddr, val)
		if err := msg.ValidateBasic(); err != nil {
			log.Fatalf("failed to check the withdraw message: %v", err)
		}
		msgs = append(msgs, msg)
	}

	err = tx.GenerateOrBroadcastTxCLI(clientCtx, flagSet, msgs...)
	res := sdk.TxResponse{}
	if err != nil {
		log.Fatalf("failed to broadcast withdraw rewards tx: %v", err)
	}
	// bts := `{"height":"10632092","txhash":"0E88782E3B247D3555D26091A299C0B0401AF2E0F810E4F3674C1431357E9BD7","codespace":"","code":0,"data":"0A1E0A1C2F636F736D6F732E62616E6B2E763162657461312E4D736753656E64","raw_log":"[{\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"cudos1fefgllqh9qpjn3laz3lmhl68kvlfhwzmq6clfq\"},{\"key\":\"amount\",\"value\":\"10000000000000000acudos\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng\"},{\"key\":\"amount\",\"value\":\"10000000000000000acudos\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.bank.v1beta1.MsgSend\"},{\"key\":\"sender\",\"value\":\"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng\"},{\"key\":\"module\",\"value\":\"bank\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"cudos1fefgllqh9qpjn3laz3lmhl68kvlfhwzmq6clfq\"},{\"key\":\"sender\",\"value\":\"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng\"},{\"key\":\"amount\",\"value\":\"10000000000000000acudos\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"coin_received","attributes":[{"key":"receiver","value":"cudos1fefgllqh9qpjn3laz3lmhl68kvlfhwzmq6clfq"},{"key":"amount","value":"10000000000000000acudos"}]},{"type":"coin_spent","attributes":[{"key":"spender","value":"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng"},{"key":"amount","value":"10000000000000000acudos"}]},{"type":"message","attributes":[{"key":"action","value":"/cosmos.bank.v1beta1.MsgSend"},{"key":"sender","value":"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng"},{"key":"module","value":"bank"}]},{"type":"transfer","attributes":[{"key":"recipient","value":"cudos1fefgllqh9qpjn3laz3lmhl68kvlfhwzmq6clfq"},{"key":"sender","value":"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng"},{"key":"amount","value":"10000000000000000acudos"}]}]}],"info":"","gas_wanted":"74869","gas_used":"69392","tx":null,"timestamp":"","events":[{"type":"coin_spent","attributes":[{"key":"c3BlbmRlcg==","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true},{"key":"YW1vdW50","value":"Mzc0MzQ1MDAwMDAwMDAwMDAwYWN1ZG9z","index":true}]},{"type":"coin_received","attributes":[{"key":"cmVjZWl2ZXI=","value":"Y3Vkb3MxN3hwZnZha20yYW1nOTYyeWxzNmY4NHoza2VsbDhjNWwzZzJsNGc=","index":true},{"key":"YW1vdW50","value":"Mzc0MzQ1MDAwMDAwMDAwMDAwYWN1ZG9z","index":true}]},{"type":"transfer","attributes":[{"key":"cmVjaXBpZW50","value":"Y3Vkb3MxN3hwZnZha20yYW1nOTYyeWxzNmY4NHoza2VsbDhjNWwzZzJsNGc=","index":true},{"key":"c2VuZGVy","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true},{"key":"YW1vdW50","value":"Mzc0MzQ1MDAwMDAwMDAwMDAwYWN1ZG9z","index":true}]},{"type":"message","attributes":[{"key":"c2VuZGVy","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true}]},{"type":"tx","attributes":[{"key":"ZmVl","value":"Mzc0MzQ1MDAwMDAwMDAwMDAwYWN1ZG9z","index":true}]},{"type":"tx","attributes":[{"key":"YWNjX3NlcQ==","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmcvNQ==","index":true}]},{"type":"tx","attributes":[{"key":"c2lnbmF0dXJl","value":"VHlscjhraGQyUnFuaDA4VWhvdEFQNStCMWV4ODl5M3daWWNHbzRqNHg5ODFLVXdEVE80b2hhU1RKSktyeGdsa0YzYkh6OGFacW4vVHV1QW4wRE5ZWlE9PQ==","index":true}]},{"type":"message","attributes":[{"key":"YWN0aW9u","value":"L2Nvc21vcy5iYW5rLnYxYmV0YTEuTXNnU2VuZA==","index":true}]},{"type":"coin_spent","attributes":[{"key":"c3BlbmRlcg==","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true},{"key":"YW1vdW50","value":"MTAwMDAwMDAwMDAwMDAwMDBhY3Vkb3M=","index":true}]},{"type":"coin_received","attributes":[{"key":"cmVjZWl2ZXI=","value":"Y3Vkb3MxZmVmZ2xscWg5cXBqbjNsYXozbG1obDY4a3ZsZmh3em1xNmNsZnE=","index":true},{"key":"YW1vdW50","value":"MTAwMDAwMDAwMDAwMDAwMDBhY3Vkb3M=","index":true}]},{"type":"transfer","attributes":[{"key":"cmVjaXBpZW50","value":"Y3Vkb3MxZmVmZ2xscWg5cXBqbjNsYXozbG1obDY4a3ZsZmh3em1xNmNsZnE=","index":true},{"key":"c2VuZGVy","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true},{"key":"YW1vdW50","value":"MTAwMDAwMDAwMDAwMDAwMDBhY3Vkb3M=","index":true}]},{"type":"message","attributes":[{"key":"c2VuZGVy","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"YmFuaw==","index":true}]}]}`
	// err = clientCtx.Codec.UnmarshalJSON([]byte(bts), &res)
	_ = outWriter.Flush()
	// bts := `{"height":"0","txhash":"272DAAADCC886944AC75ABDB18185D59C5E4451EFFC052300F509C608CAF7C5A","codespace":"sdk","code":13,"data":"","raw_log":"insufficient fees; got: 37434500000000000acudos required: 374345000000000000acudos: insufficient fee","logs":[],"info":"","gas_wanted":"74869","gas_used":"0","tx":null,"timestamp":"","events":[]}`
	err = clientCtx.Codec.UnmarshalJSON([]byte(outBuffer.Bytes()), &res)
	if err != nil {
		log.Fatalf("failed to unmarshal the result of send tx: %v", err)
	}
	if res.Code != 0 {
		log.Fatalf("widthdraw rewards tx faild: %s", res.RawLog)
	}
	fmt.Printf("widthdraw rewards success: %#v\n", res)

	return "", nil
}

func send(amount string) {
	sdk.DefaultPowerReduction = sdk.NewIntFromUint64(1000000000000000000)
	startup.SetConfig()

	cobraCmd := &cobra.Command{}
	cobraCmd.SetContext(context.Background())

	flagSet := cobraCmd.Flags()
	//	flagSet.String(flags.FlagFrom, keyringAccount, "")
	flagSet.Bool(flags.FlagSkipConfirmation, true, "")
	flagSet.String(cli.OutputFlag, "json", "")
	flagSet.String(flags.FlagHome, dataHomePath, "")
	// flagSet.String(flags.FlagKeyringDir, dataHomePath, "")
	//	flagSet.String(flags.FlagChainID, chainID, "")
	//	flagSet.String(flags.FlagKeyringBackend, keyringBackend, "")
	//	flagSet.String(flags.FlagNode, nodeAddress, "")
	flagSet.String(flags.FlagNote, memo, "")
	//	flagSet.Float64(flags.FlagGasAdjustment, gasAdj, "")
	//	flagSet.String(flags.FlagGas, gas, "")
	//	flagSet.String(flags.FlagGasPrices, gasPrices, "")

	aminoCodec := codec.NewLegacyAmino()
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txConfig := authtx.NewTxConfig(marshaler, authtx.DefaultSignModes)

	std.RegisterLegacyAminoCodec(aminoCodec)
	std.RegisterInterfaces(interfaceRegistry)
	simapp.ModuleBasics.RegisterLegacyAminoCodec(aminoCodec)
	simapp.ModuleBasics.RegisterInterfaces(interfaceRegistry)

	var outBuffer bytes.Buffer
	outWriter := bufio.NewWriter(&outBuffer)

	clientCtx, err := client.GetClientTxContext(cobraCmd)
	if err != nil {
		log.Fatalf("failed to obtain client context: %v", err)
	}

	clientCtx = clientCtx.
		WithCodec(marshaler).
		WithJSONCodec(marshaler).
		WithInterfaceRegistry(interfaceRegistry).
		WithTxConfig(txConfig).
		WithLegacyAmino(aminoCodec).
		WithInput(os.Stdin).
		WithOutput(outWriter).
		WithBroadcastMode(flags.BroadcastBlock).
		WithAccountRetriever(types.AccountRetriever{})
	//WithBroadcastMode(flags.BroadcastBlock).
	//WithHomeDir(dataHomePath)

	coins, err := sdk.ParseCoinsNormalized(amount)
	if err != nil {
		log.Fatalf("failed to parse the amount to send: %v", err)
	}

	msg := &banktypes.MsgSend{
		FromAddress: clientCtx.GetFromAddress().String(),
		ToAddress:   toAddress,
		Amount:      coins,
	}
	if err := msg.ValidateBasic(); err != nil {
		log.Fatalf("failed to parse the amount to send: %v", err)
	}
	err = tx.GenerateOrBroadcastTxCLI(clientCtx, flagSet, msg)
	res := sdk.TxResponse{}
	if err != nil {
		log.Fatalf("failed to broadcast send tx: %v", err)
	}
	// bts := `{"height":"10632092","txhash":"0E88782E3B247D3555D26091A299C0B0401AF2E0F810E4F3674C1431357E9BD7","codespace":"","code":0,"data":"0A1E0A1C2F636F736D6F732E62616E6B2E763162657461312E4D736753656E64","raw_log":"[{\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"cudos1fefgllqh9qpjn3laz3lmhl68kvlfhwzmq6clfq\"},{\"key\":\"amount\",\"value\":\"10000000000000000acudos\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng\"},{\"key\":\"amount\",\"value\":\"10000000000000000acudos\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.bank.v1beta1.MsgSend\"},{\"key\":\"sender\",\"value\":\"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng\"},{\"key\":\"module\",\"value\":\"bank\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"cudos1fefgllqh9qpjn3laz3lmhl68kvlfhwzmq6clfq\"},{\"key\":\"sender\",\"value\":\"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng\"},{\"key\":\"amount\",\"value\":\"10000000000000000acudos\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"coin_received","attributes":[{"key":"receiver","value":"cudos1fefgllqh9qpjn3laz3lmhl68kvlfhwzmq6clfq"},{"key":"amount","value":"10000000000000000acudos"}]},{"type":"coin_spent","attributes":[{"key":"spender","value":"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng"},{"key":"amount","value":"10000000000000000acudos"}]},{"type":"message","attributes":[{"key":"action","value":"/cosmos.bank.v1beta1.MsgSend"},{"key":"sender","value":"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng"},{"key":"module","value":"bank"}]},{"type":"transfer","attributes":[{"key":"recipient","value":"cudos1fefgllqh9qpjn3laz3lmhl68kvlfhwzmq6clfq"},{"key":"sender","value":"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng"},{"key":"amount","value":"10000000000000000acudos"}]}]}],"info":"","gas_wanted":"74869","gas_used":"69392","tx":null,"timestamp":"","events":[{"type":"coin_spent","attributes":[{"key":"c3BlbmRlcg==","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true},{"key":"YW1vdW50","value":"Mzc0MzQ1MDAwMDAwMDAwMDAwYWN1ZG9z","index":true}]},{"type":"coin_received","attributes":[{"key":"cmVjZWl2ZXI=","value":"Y3Vkb3MxN3hwZnZha20yYW1nOTYyeWxzNmY4NHoza2VsbDhjNWwzZzJsNGc=","index":true},{"key":"YW1vdW50","value":"Mzc0MzQ1MDAwMDAwMDAwMDAwYWN1ZG9z","index":true}]},{"type":"transfer","attributes":[{"key":"cmVjaXBpZW50","value":"Y3Vkb3MxN3hwZnZha20yYW1nOTYyeWxzNmY4NHoza2VsbDhjNWwzZzJsNGc=","index":true},{"key":"c2VuZGVy","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true},{"key":"YW1vdW50","value":"Mzc0MzQ1MDAwMDAwMDAwMDAwYWN1ZG9z","index":true}]},{"type":"message","attributes":[{"key":"c2VuZGVy","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true}]},{"type":"tx","attributes":[{"key":"ZmVl","value":"Mzc0MzQ1MDAwMDAwMDAwMDAwYWN1ZG9z","index":true}]},{"type":"tx","attributes":[{"key":"YWNjX3NlcQ==","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmcvNQ==","index":true}]},{"type":"tx","attributes":[{"key":"c2lnbmF0dXJl","value":"VHlscjhraGQyUnFuaDA4VWhvdEFQNStCMWV4ODl5M3daWWNHbzRqNHg5ODFLVXdEVE80b2hhU1RKSktyeGdsa0YzYkh6OGFacW4vVHV1QW4wRE5ZWlE9PQ==","index":true}]},{"type":"message","attributes":[{"key":"YWN0aW9u","value":"L2Nvc21vcy5iYW5rLnYxYmV0YTEuTXNnU2VuZA==","index":true}]},{"type":"coin_spent","attributes":[{"key":"c3BlbmRlcg==","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true},{"key":"YW1vdW50","value":"MTAwMDAwMDAwMDAwMDAwMDBhY3Vkb3M=","index":true}]},{"type":"coin_received","attributes":[{"key":"cmVjZWl2ZXI=","value":"Y3Vkb3MxZmVmZ2xscWg5cXBqbjNsYXozbG1obDY4a3ZsZmh3em1xNmNsZnE=","index":true},{"key":"YW1vdW50","value":"MTAwMDAwMDAwMDAwMDAwMDBhY3Vkb3M=","index":true}]},{"type":"transfer","attributes":[{"key":"cmVjaXBpZW50","value":"Y3Vkb3MxZmVmZ2xscWg5cXBqbjNsYXozbG1obDY4a3ZsZmh3em1xNmNsZnE=","index":true},{"key":"c2VuZGVy","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true},{"key":"YW1vdW50","value":"MTAwMDAwMDAwMDAwMDAwMDBhY3Vkb3M=","index":true}]},{"type":"message","attributes":[{"key":"c2VuZGVy","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"YmFuaw==","index":true}]}]}`
	// err = clientCtx.Codec.UnmarshalJSON([]byte(bts), &res)
	_ = outWriter.Flush()
	// bts := `{"height":"0","txhash":"272DAAADCC886944AC75ABDB18185D59C5E4451EFFC052300F509C608CAF7C5A","codespace":"sdk","code":13,"data":"","raw_log":"insufficient fees; got: 37434500000000000acudos required: 374345000000000000acudos: insufficient fee","logs":[],"info":"","gas_wanted":"74869","gas_used":"0","tx":null,"timestamp":"","events":[]}`
	err = clientCtx.Codec.UnmarshalJSON([]byte(outBuffer.Bytes()), &res)
	if err != nil {
		log.Fatalf("failed to unmarshal the result of send tx: %v", err)
	}
	if res.Code != 0 {
		log.Fatalf("send tx faild: %s", res.RawLog)
	}
	fmt.Printf("send success: %#v\n", res)
}*/
