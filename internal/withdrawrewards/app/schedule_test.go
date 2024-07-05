package app

import (
	"context"
	"errors"
	"testing"
	"time"

	"cudos-task/contract"
	contractmocks "cudos-task/contract/mocks"
	apimocks "cudos-task/internal/withdrawrewards/app/cudos/contract/mocks"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCudosCommand_RunSchedule(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	marshaller := codec.NewProtoCodec(interfaceRegistry)

	type fields struct {
		withdrawRewardAmount string
		withdraw             string
		withdrawErr          error
		send                 string
		sendErr              error
	}
	type expects struct {
		withdrawOk       bool
		withdrawContains string
		sendOk           bool
		sendContains     string
	}
	tests := []struct {
		name    string
		fields  fields
		expects expects
	}{
		{
			name: "withdraws rewards and then send them successfully",
			fields: fields{
				withdrawRewardAmount: "100303843979829000acudos",
				withdraw:             `{"height":"14499078","txhash":"2C23F5DA94A0F4BA5EAD2A9C11B53969A6952E6455637129D723D784C3690955","codespace":"","code":0,"data":"0A390A372F636F736D6F732E646973747269627574696F6E2E763162657461312E4D7367576974686472617744656C656761746F725265776172640A390A372F636F736D6F732E646973747269627574696F6E2E763162657461312E4D7367576974686472617744656C656761746F72526577617264","raw_log":"[{\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng\"},{\"key\":\"amount\",\"value\":\"100303843979829000acudos\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"cudos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8xwdrh2\"},{\"key\":\"amount\",\"value\":\"100303843979829000acudos\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward\"},{\"key\":\"sender\",\"value\":\"cudos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8xwdrh2\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng\"},{\"key\":\"sender\",\"value\":\"cudos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8xwdrh2\"},{\"key\":\"amount\",\"value\":\"100303843979829000acudos\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\",\"value\":\"100303843979829000acudos\"},{\"key\":\"validator\",\"value\":\"cudosvaloper198qaeg4wkf9tn7y345dhk2wyjmm0krdm68jp09\"}]}]},{\"msg_index\":1,\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng\"},{\"key\":\"amount\",\"value\":\"108239694834454299acudos\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"cudos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8xwdrh2\"},{\"key\":\"amount\",\"value\":\"108239694834454299acudos\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward\"},{\"key\":\"sender\",\"value\":\"cudos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8xwdrh2\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng\"},{\"key\":\"sender\",\"value\":\"cudos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8xwdrh2\"},{\"key\":\"amount\",\"value\":\"108239694834454299acudos\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\",\"value\":\"108239694834454299acudos\"},{\"key\":\"validator\",\"value\":\"cudosvaloper1v08h0h7sv3wm0t9uaz6x2q26p0g63tyzw68ynj\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"coin_received","attributes":[{"key":"receiver","value":"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng"},{"key":"amount","value":"100303843979829000acudos"}]},{"type":"coin_spent","attributes":[{"key":"spender","value":"cudos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8xwdrh2"},{"key":"amount","value":"100303843979829000acudos"}]},{"type":"message","attributes":[{"key":"action","value":"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward"},{"key":"sender","value":"cudos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8xwdrh2"},{"key":"module","value":"distribution"},{"key":"sender","value":"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng"}]},{"type":"transfer","attributes":[{"key":"recipient","value":"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng"},{"key":"sender","value":"cudos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8xwdrh2"},{"key":"amount","value":"100303843979829000acudos"}]},{"type":"withdraw_rewards","attributes":[{"key":"amount","value":"100303843979829000acudos"},{"key":"validator","value":"cudosvaloper198qaeg4wkf9tn7y345dhk2wyjmm0krdm68jp09"}]}]},{"msg_index":1,"log":"","events":[{"type":"coin_received","attributes":[{"key":"receiver","value":"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng"},{"key":"amount","value":"108239694834454299acudos"}]},{"type":"coin_spent","attributes":[{"key":"spender","value":"cudos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8xwdrh2"},{"key":"amount","value":"108239694834454299acudos"}]},{"type":"message","attributes":[{"key":"action","value":"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward"},{"key":"sender","value":"cudos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8xwdrh2"},{"key":"module","value":"distribution"},{"key":"sender","value":"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng"}]},{"type":"transfer","attributes":[{"key":"recipient","value":"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng"},{"key":"sender","value":"cudos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8xwdrh2"},{"key":"amount","value":"108239694834454299acudos"}]},{"type":"withdraw_rewards","attributes":[{"key":"amount","value":"108239694834454299acudos"},{"key":"validator","value":"cudosvaloper1v08h0h7sv3wm0t9uaz6x2q26p0g63tyzw68ynj"}]}]}],"info":"","gas_wanted":"229885","gas_used":"188708","tx":null,"timestamp":"","events":[{"type":"coin_spent","attributes":[{"key":"c3BlbmRlcg==","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true},{"key":"YW1vdW50","value":"MTE0OTQyNTAwMDAwMDAwMDAwMGFjdWRvcw==","index":true}]},{"type":"coin_received","attributes":[{"key":"cmVjZWl2ZXI=","value":"Y3Vkb3MxN3hwZnZha20yYW1nOTYyeWxzNmY4NHoza2VsbDhjNWwzZzJsNGc=","index":true},{"key":"YW1vdW50","value":"MTE0OTQyNTAwMDAwMDAwMDAwMGFjdWRvcw==","index":true}]},{"type":"transfer","attributes":[{"key":"cmVjaXBpZW50","value":"Y3Vkb3MxN3hwZnZha20yYW1nOTYyeWxzNmY4NHoza2VsbDhjNWwzZzJsNGc=","index":true},{"key":"c2VuZGVy","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true},{"key":"YW1vdW50","value":"MTE0OTQyNTAwMDAwMDAwMDAwMGFjdWRvcw==","index":true}]},{"type":"message","attributes":[{"key":"c2VuZGVy","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true}]},{"type":"tx","attributes":[{"key":"ZmVl","value":"MTE0OTQyNTAwMDAwMDAwMDAwMGFjdWRvcw==","index":true},{"key":"ZmVlX3BheWVy","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true}]},{"type":"tx","attributes":[{"key":"YWNjX3NlcQ==","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmcvMg==","index":true}]},{"type":"tx","attributes":[{"key":"c2lnbmF0dXJl","value":"ZWV5OFVDSDR1QU9HdFFwNktERWJTVHdESHBPMkkrNkFPRHA5MUtxQzJaWkhjTCtFcjZpemhUTkRISjhjNlpzQWpBVXBqWHFzV0NpdHpQbVZaUURkbEE9PQ==","index":true}]},{"type":"message","attributes":[{"key":"YWN0aW9u","value":"L2Nvc21vcy5kaXN0cmlidXRpb24udjFiZXRhMS5Nc2dXaXRoZHJhd0RlbGVnYXRvclJld2FyZA==","index":true}]},{"type":"coin_spent","attributes":[{"key":"c3BlbmRlcg==","value":"Y3Vkb3MxanY2NXMzZ3JxZjZ2NmpsM2RwNHQ2Yzl0OXJrOTljZDh4d2RyaDI=","index":true},{"key":"YW1vdW50","value":"MTAwMzAzODQzOTc5ODI5MDAwYWN1ZG9z","index":true}]},{"type":"coin_received","attributes":[{"key":"cmVjZWl2ZXI=","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true},{"key":"YW1vdW50","value":"MTAwMzAzODQzOTc5ODI5MDAwYWN1ZG9z","index":true}]},{"type":"transfer","attributes":[{"key":"cmVjaXBpZW50","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true},{"key":"c2VuZGVy","value":"Y3Vkb3MxanY2NXMzZ3JxZjZ2NmpsM2RwNHQ2Yzl0OXJrOTljZDh4d2RyaDI=","index":true},{"key":"YW1vdW50","value":"MTAwMzAzODQzOTc5ODI5MDAwYWN1ZG9z","index":true}]},{"type":"message","attributes":[{"key":"c2VuZGVy","value":"Y3Vkb3MxanY2NXMzZ3JxZjZ2NmpsM2RwNHQ2Yzl0OXJrOTljZDh4d2RyaDI=","index":true}]},{"type":"withdraw_rewards","attributes":[{"key":"YW1vdW50","value":"MTAwMzAzODQzOTc5ODI5MDAwYWN1ZG9z","index":true},{"key":"dmFsaWRhdG9y","value":"Y3Vkb3N2YWxvcGVyMTk4cWFlZzR3a2Y5dG43eTM0NWRoazJ3eWptbTBrcmRtNjhqcDA5","index":true}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"ZGlzdHJpYnV0aW9u","index":true},{"key":"c2VuZGVy","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true}]},{"type":"message","attributes":[{"key":"YWN0aW9u","value":"L2Nvc21vcy5kaXN0cmlidXRpb24udjFiZXRhMS5Nc2dXaXRoZHJhd0RlbGVnYXRvclJld2FyZA==","index":true}]},{"type":"coin_spent","attributes":[{"key":"c3BlbmRlcg==","value":"Y3Vkb3MxanY2NXMzZ3JxZjZ2NmpsM2RwNHQ2Yzl0OXJrOTljZDh4d2RyaDI=","index":true},{"key":"YW1vdW50","value":"MTA4MjM5Njk0ODM0NDU0Mjk5YWN1ZG9z","index":true}]},{"type":"coin_received","attributes":[{"key":"cmVjZWl2ZXI=","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true},{"key":"YW1vdW50","value":"MTA4MjM5Njk0ODM0NDU0Mjk5YWN1ZG9z","index":true}]},{"type":"transfer","attributes":[{"key":"cmVjaXBpZW50","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true},{"key":"c2VuZGVy","value":"Y3Vkb3MxanY2NXMzZ3JxZjZ2NmpsM2RwNHQ2Yzl0OXJrOTljZDh4d2RyaDI=","index":true},{"key":"YW1vdW50","value":"MTA4MjM5Njk0ODM0NDU0Mjk5YWN1ZG9z","index":true}]},{"type":"message","attributes":[{"key":"c2VuZGVy","value":"Y3Vkb3MxanY2NXMzZ3JxZjZ2NmpsM2RwNHQ2Yzl0OXJrOTljZDh4d2RyaDI=","index":true}]},{"type":"withdraw_rewards","attributes":[{"key":"YW1vdW50","value":"MTA4MjM5Njk0ODM0NDU0Mjk5YWN1ZG9z","index":true},{"key":"dmFsaWRhdG9y","value":"Y3Vkb3N2YWxvcGVyMXYwOGgwaDdzdjN3bTB0OXVhejZ4MnEyNnAwZzYzdHl6dzY4eW5q","index":true}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"ZGlzdHJpYnV0aW9u","index":true},{"key":"c2VuZGVy","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true}]}]}`,
				withdrawErr:          nil,
				send:                 `{"height":"10632092","txhash":"0E88782E3B247D3555D26091A299C0B0401AF2E0F810E4F3674C1431357E9BD7","codespace":"","code":0,"data":"0A1E0A1C2F636F736D6F732E62616E6B2E763162657461312E4D736753656E64","raw_log":"[{\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"cudos1fefgllqh9qpjn3laz3lmhl68kvlfhwzmq6clfq\"},{\"key\":\"amount\",\"value\":\"10000000000000000acudos\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng\"},{\"key\":\"amount\",\"value\":\"10000000000000000acudos\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.bank.v1beta1.MsgSend\"},{\"key\":\"sender\",\"value\":\"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng\"},{\"key\":\"module\",\"value\":\"bank\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"cudos1fefgllqh9qpjn3laz3lmhl68kvlfhwzmq6clfq\"},{\"key\":\"sender\",\"value\":\"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng\"},{\"key\":\"amount\",\"value\":\"10000000000000000acudos\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"coin_received","attributes":[{"key":"receiver","value":"cudos1fefgllqh9qpjn3laz3lmhl68kvlfhwzmq6clfq"},{"key":"amount","value":"10000000000000000acudos"}]},{"type":"coin_spent","attributes":[{"key":"spender","value":"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng"},{"key":"amount","value":"10000000000000000acudos"}]},{"type":"message","attributes":[{"key":"action","value":"/cosmos.bank.v1beta1.MsgSend"},{"key":"sender","value":"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng"},{"key":"module","value":"bank"}]},{"type":"transfer","attributes":[{"key":"recipient","value":"cudos1fefgllqh9qpjn3laz3lmhl68kvlfhwzmq6clfq"},{"key":"sender","value":"cudos1px780rkpusl7rfcs56v2dfqgv04hfqv4jt9rng"},{"key":"amount","value":"10000000000000000acudos"}]}]}],"info":"","gas_wanted":"74869","gas_used":"69392","tx":null,"timestamp":"","events":[{"type":"coin_spent","attributes":[{"key":"c3BlbmRlcg==","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true},{"key":"YW1vdW50","value":"Mzc0MzQ1MDAwMDAwMDAwMDAwYWN1ZG9z","index":true}]},{"type":"coin_received","attributes":[{"key":"cmVjZWl2ZXI=","value":"Y3Vkb3MxN3hwZnZha20yYW1nOTYyeWxzNmY4NHoza2VsbDhjNWwzZzJsNGc=","index":true},{"key":"YW1vdW50","value":"Mzc0MzQ1MDAwMDAwMDAwMDAwYWN1ZG9z","index":true}]},{"type":"transfer","attributes":[{"key":"cmVjaXBpZW50","value":"Y3Vkb3MxN3hwZnZha20yYW1nOTYyeWxzNmY4NHoza2VsbDhjNWwzZzJsNGc=","index":true},{"key":"c2VuZGVy","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true},{"key":"YW1vdW50","value":"Mzc0MzQ1MDAwMDAwMDAwMDAwYWN1ZG9z","index":true}]},{"type":"message","attributes":[{"key":"c2VuZGVy","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true}]},{"type":"tx","attributes":[{"key":"ZmVl","value":"Mzc0MzQ1MDAwMDAwMDAwMDAwYWN1ZG9z","index":true}]},{"type":"tx","attributes":[{"key":"YWNjX3NlcQ==","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmcvNQ==","index":true}]},{"type":"tx","attributes":[{"key":"c2lnbmF0dXJl","value":"VHlscjhraGQyUnFuaDA4VWhvdEFQNStCMWV4ODl5M3daWWNHbzRqNHg5ODFLVXdEVE80b2hhU1RKSktyeGdsa0YzYkh6OGFacW4vVHV1QW4wRE5ZWlE9PQ==","index":true}]},{"type":"message","attributes":[{"key":"YWN0aW9u","value":"L2Nvc21vcy5iYW5rLnYxYmV0YTEuTXNnU2VuZA==","index":true}]},{"type":"coin_spent","attributes":[{"key":"c3BlbmRlcg==","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true},{"key":"YW1vdW50","value":"MTAwMDAwMDAwMDAwMDAwMDBhY3Vkb3M=","index":true}]},{"type":"coin_received","attributes":[{"key":"cmVjZWl2ZXI=","value":"Y3Vkb3MxZmVmZ2xscWg5cXBqbjNsYXozbG1obDY4a3ZsZmh3em1xNmNsZnE=","index":true},{"key":"YW1vdW50","value":"MTAwMDAwMDAwMDAwMDAwMDBhY3Vkb3M=","index":true}]},{"type":"transfer","attributes":[{"key":"cmVjaXBpZW50","value":"Y3Vkb3MxZmVmZ2xscWg5cXBqbjNsYXozbG1obDY4a3ZsZmh3em1xNmNsZnE=","index":true},{"key":"c2VuZGVy","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true},{"key":"YW1vdW50","value":"MTAwMDAwMDAwMDAwMDAwMDBhY3Vkb3M=","index":true}]},{"type":"message","attributes":[{"key":"c2VuZGVy","value":"Y3Vkb3MxcHg3ODBya3B1c2w3cmZjczU2djJkZnFndjA0aGZxdjRqdDlybmc=","index":true}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"YmFuaw==","index":true}]}]}`,
				sendErr:              nil,
			},
			expects: expects{
				withdrawOk:       true,
				withdrawContains: "withdraw rewards collected",
				sendOk:           true,
				sendContains:     "sent coins",
			},
		},
		{
			name: "withdraws rewards with insufficient fees and fails",
			fields: fields{
				withdrawRewardAmount: "",
				withdraw:             `{"height":"0","txhash":"272DAAADCC886944AC75ABDB18185D59C5E4451EFFC052300F509C608CAF7C5A","codespace":"sdk","code":13,"data":"","raw_log":"insufficient fees; got: 37434500000000000acudos required: 374345000000000000acudos: insufficient fee","logs":[],"info":"","gas_wanted":"74869","gas_used":"0","tx":null,"timestamp":"","events":[]}`,
				withdrawErr:          errors.New("widthdraw rewards tx faild: insufficient fees; got: 37434500000000000acudos required: 374345000000000000acudos: insufficient fee"),
				send:                 ``,
				sendErr:              nil,
			},
			expects: expects{
				withdrawOk:       false,
				withdrawContains: "widthdraw rewards tx faild: insufficient fees; got: 37434500000000000acudos",
				sendOk:           false,
				sendContains:     "sent coins",
			},
		},
		{
			name: "run withdraw then close the context and fails",
			fields: fields{
				withdrawRewardAmount: "",
				withdraw:             ``,
				withdrawErr:          nil,
				send:                 ``,
				sendErr:              nil,
			},
			expects: expects{
				withdrawOk:       false,
				withdrawContains: "",
				sendOk:           false,
				sendContains:     "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			withdrawSender := new(apimocks.CudosWithdrawSender)

			resWithdraw := types.TxResponse{}
			if tt.expects.withdrawOk {
				err := marshaller.UnmarshalJSON([]byte(tt.fields.withdraw), &resWithdraw)
				require.Nil(t, err)
			}

			resSend := types.TxResponse{}
			if tt.expects.sendOk {
				err := marshaller.UnmarshalJSON([]byte(tt.fields.send), &resSend)
				require.Nil(t, err)
			}

			var err error
			amount := types.NewInt64Coin(contract.Denom, 0)
			if tt.fields.withdrawRewardAmount != "" {
				amount, err = types.ParseCoinNormalized(tt.fields.withdrawRewardAmount)
				require.Nil(t, err)
			}

			withdrawSender.On("Withdraw").Return(amount, &resWithdraw, tt.fields.withdrawErr)
			withdrawSender.On("Send", mock.AnythingOfType("types.Coin")).Return(amount, &resSend, tt.fields.sendErr)

			shutdown := new(contractmocks.ShutdownReady)
			shutdown.On("SetReady", mock.Anything).Return(true)

			ccx := NewCudosCommand(shutdown, withdrawSender)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			out := make(chan string, 20)
			interval := time.Second
			ccx.RunSchedule(ctx, out, interval)

			var withdrawMsg, sendMsg string

			// skip first n messages
			for i := 0; i < 1; i++ {
				select {
				case <-out:
				case <-time.After(interval):
				}
			}

			if tt.expects.withdrawOk || tt.fields.withdrawErr != nil {
				select {
				case withdrawMsg = <-out:
				case <-time.After(interval):
				}
			}

			if tt.expects.sendOk || tt.fields.sendErr != nil {
				select {
				case sendMsg = <-out:
				case <-time.After(interval):
				}
			}
			cancel()

			if tt.expects.withdrawOk || tt.fields.withdrawErr != nil {
				assert.Contains(t, withdrawMsg, tt.expects.withdrawContains)
			}
			if tt.expects.sendOk || tt.fields.sendErr != nil {
				assert.Contains(t, sendMsg, tt.expects.sendContains)
			}
		})
	}
}

func TestCudosCommand_RunSchedule_AndCloseContext(_ *testing.T) {
	withdrawSender := new(apimocks.CudosWithdrawSender)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shutdown := new(contractmocks.ShutdownReady)
	shutdown.On("SetReady", mock.Anything).Return(true).Run(func(_ mock.Arguments) {
		// cancel the context after first call of SetReady
		cancel()
	})

	// Withdraw shouldn't be called at all in this test
	withdrawSender.On("Withdraw").Times(0)

	ccx := NewCudosCommand(shutdown, withdrawSender)

	out := make(chan string, 20)
	interval := 100 * time.Millisecond
	ccx.RunSchedule(ctx, out, interval)

	select {
	case <-out:
	case <-time.After(500 * time.Millisecond):
	}
}
