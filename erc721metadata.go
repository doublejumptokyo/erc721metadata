package erc721metadata

// ERC721Metadata is ERC721 Metadata JSON
type ERC721Metadata struct {
	// ERC721 EIP1047
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Image       string                 `json:"image"`
	Attributes  map[string]interface{} `json:"attributes"` // PR#1071

	// OpenSea
	ExternalURL     string `json:"external_url"`     // OpenSea
	BackgroundColor string `json:"background_color"` // OpenSea

	// Rare Bits
	ID                       int64    `json:"id"`
	TokenContractAddress     string   `json:"token_contract_address"`
	TokenID                  string   `json:"token_id"`
	TokenOwnerAddress        string   `json:"token_owner_address"`
	ImageURL                 string   `json:"image_url"`
	HomeURL                  string   `json:"home_url"`
	URL                      string   `json:"url"`
	Color                    string   `json:"color"`
	Tags                     []string `json:"tags"`
	TokenCreatedTimestamp    int64    `json:"token_created_timestamp"`
	LastSoldTimestamp        int64    `json:"last_sold_timestamp"`
	LastSalePrice            string   `json:"last_sale_price"`
	EstimatedValue           string   `json:"estimated_value"`
	TokenContractDisplayName string   `json:"token_contract_display_name"`
	TokenContractID          string   `json:"token_contract_id"`
}

// NewERC721ObjectMetadata returns *NewERC721Metadata
func NewERC721Metadata() (*ERC721Metadata, error) {
	ret := new(ERC721Metadata)
	ret.Attributes = map[string]interface{}{}
	return ret, nil
}
