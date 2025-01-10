package blockchain

type ChainId int

const (
	ARBITRUM ChainId = 42161
)

var SupportedChains = []ChainId{ARBITRUM}

func IsSupportedChain(chainId ChainId) bool {
	for _, supportedChainId := range SupportedChains {
		if chainId == supportedChainId {
			return true
		}
	}
	return false
}
