package gmx

import (
	"bytes"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	chainlinkABI       abi.ABI
	fastPriceFeedV1ABI abi.ABI
	fastPriceFeedV2ABI abi.ABI
	pancakePairABI     abi.ABI
	priceFeedABI       abi.ABI
	vaultPriceFeedABI  abi.ABI
	erc20ABI           abi.ABI
)

func init() {
	builder := []struct {
		ABI  *abi.ABI
		data []byte
	}{
		{&chainlinkABI, chainlinkFlagsJson},
		{&fastPriceFeedV1ABI, fastPriceFeedV1Json},
		{&fastPriceFeedV2ABI, fastPriceFeedV2Json},
		{&pancakePairABI, pancakePairJson},
		{&priceFeedABI, priceFeedJson},
		{&vaultPriceFeedABI, vaultPriceFeedJson},
		{&erc20ABI, erc20Json},
	}

	for _, b := range builder {
		var err error
		*b.ABI, err = abi.JSON(bytes.NewReader(b.data))
		if err != nil {
			panic(err)
		}
	}
}
