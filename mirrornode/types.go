package mirrornode

type NetworkType string

const (
	MainnetNetwork    NetworkType = "mainnet"
	TestnetNetwork    NetworkType = "testnet"
	PreviewnetNetwork NetworkType = "previewnet"
)

const (
	mainnetBaseURL    = "https://mainnet-public.mirrornode.hedera.com"
	testnetBaseURL    = "https://testnet.mirrornode.hedera.com"
	previewnetBaseURL = "https://previewnet.mirrornode.hedera.com"
)

type NFT struct {
	AccountID         *string `json:"account_id"`
	CreatedTimestamp  *string `json:"created_timestamp"`
	DelegatingSpender *string `json:"delegating_spender"`
	Deleted           bool    `json:"deleted,omitempty"`
	Metadata          *string `json:"metadata,omitempty"`
	ModifiedTimestamp *string `json:"modified_timestamp"`
	SerialNumber      int64   `json:"serial_number,omitempty"`
	Spender           *string `json:"spender"`
	TokenID           string  `json:"token_id"`
}

type Links struct {
	Next *string `json:"next"`
}

type NFTsResponse struct {
	NFTs  []NFT `json:"nfts"`
	Links Links `json:"links"`
}

type APIError struct {
	Status struct {
		Messages []struct {
			Message string `json:"message"`
		} `json:"messages"`
	} `json:"_status"`
}
