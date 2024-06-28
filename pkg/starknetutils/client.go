package starknetutils

import (
	"context"
	"errors"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/NethermindEth/starknet.go/utils"
)

func CallReadFunction(rpcAddress string, smartContractAddressHex string, contractMethod string, callData []*felt.Felt) ([]*felt.Felt, error) {

	clientv02, err := rpc.NewProvider(rpcAddress)

	if err != nil {
		return nil, err
	}

	contractAddress, err := utils.HexToFelt(smartContractAddressHex)
	if err != nil {
		return nil, err
	}

	// Make read contract call
	tx := rpc.FunctionCall{
		ContractAddress:    contractAddress,
		EntryPointSelector: utils.GetSelectorFromNameFelt(contractMethod),
		Calldata:           callData,
	}

	callResp, rpcErr := clientv02.Call(context.Background(), tx, rpc.BlockID{Tag: "latest"})

	if callResp == nil {
		return nil, errors.New(rpcErr.Message)
	}
	return callResp, nil
}
