// package blockchain

// import (
// 	"context"
// 	// "math/big"
// 	"time"

// 	// "github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// func NewClient(chainId ChainId) (*ethclient.Client, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	rpc, err := GetRpc(chainId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	client, err := ethclient.DialContext(ctx, rpc)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return client, nil
// }

// func GetUserBalance(client *ethclient.Client, address common.Address) (*big.Int, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	balance, err := client.BalanceAt(ctx, address, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return balance, nil
// }

// func GetInitialLiquidity(client *ethclient.Client, chainId ChainId) (*big.Int, error) {
// 	launchpadAddr := common.HexToAddress(Launchpads[chainId])
// 	launchpadContract := GetContract(client, launchpadAddr, "launchpad")

// 	initialLiquidityData := &InitialLiquidityData{}

// 	if err := CallContractMethod(launchpadContract, "INITIAL_LIQ", &initialLiquidityData.Result); err != nil {
// 		return nil, err
// 	}

// 	return initialLiquidityData.Result, nil
// }