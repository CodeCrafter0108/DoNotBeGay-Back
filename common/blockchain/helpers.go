package blockchain

import (
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetContract(client *ethclient.Client, address common.Address, abiName string) *bind.BoundContract {
	var contractABIString string

	switch abiName {
	case "launchpad":
		contractABIString = launchpadABI
	default:
		log.Fatalf("ABI not found for contract: %s", abiName)
	}

	contractABI, err := abi.JSON(strings.NewReader(contractABIString))
	if err != nil {
		panic(err)
	}

	contract := bind.NewBoundContract(address, contractABI, client, client, client)
	return contract
}

func CallContractMethod(contract *bind.BoundContract, methodName string, result ...interface{}) error {
	return contract.Call(&bind.CallOpts{}, &result, methodName)
}
