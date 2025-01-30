package blockchain

type ChainId int

const (
	SONIC ChainId = 146
)

var SupportedChains = []ChainId{SONIC}

func IsSupportedChain(chainId ChainId) bool {
	for _, supportedChainId := range SupportedChains {
		if chainId == supportedChainId {
			return true
		}
	}
	return false
}
