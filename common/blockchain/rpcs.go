package blockchain

// import (
// 	"errors"
// 	"strconv"
// )

// type RpcMap map[ChainId]string

// var Rpcs = RpcMap{
// 	ARBITRUM: "https://rpc.ankr.com/arbitrum",
// }

// const (
// 	ErrInvalidChainId = "no rpc for chain id"
// )

// func GetRpc(chainId ChainId) (string, error) {
// 	rpc := Rpcs[chainId]

// 	if rpc == "" {
// 		return "", errors.New(ErrInvalidChainId + " " + strconv.Itoa(int(chainId)))
// 	}

// 	return rpc, nil
// }
