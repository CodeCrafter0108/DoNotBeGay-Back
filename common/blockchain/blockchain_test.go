// package blockchain

// import (
// 	"math/big"
// 	"testing"
// )

// func TestGetInitialLiquidity(t *testing.T) {
// 	const chainId ChainId = ARBITRUM

// 	client, err := NewClient(chainId)
// 	if err != nil {
// 		t.Skipf("Skipping test: unable to create client for ChainId %d: %v", chainId, err)
// 	}

// 	liquidity, err := GetInitialLiquidity(client, chainId)
// 	if err != nil {
// 		t.Fatalf("Failed to get initial liquidity: %v", err)
// 	}

// 	if liquidity.Cmp(big.NewInt(0)) <= 0 {
// 		t.Errorf("Expected initial liquidity to be greater than 0, got: %v", liquidity)
// 	} else {
// 		t.Logf("Initial liquidity is valid: %v", liquidity)
// 	}
// }
