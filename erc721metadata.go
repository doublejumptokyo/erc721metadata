package erc721metadata

import (
	"encoding/json"
	"math/big"
)

// ERC721Metadata is ERC721 Metadata JSON
type ERC721Metadata struct {
	// ERC721 EIP1047
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`

	// OpenSea
	Attributes       *json.RawMessage `json:"attributes,omitempty"`       // EIPs#1071
	ExternalURL      string           `json:"external_url,omitempty"`     // OpenSea
	ExternalURLCamel string           `json:"externalUrl,omitempty"`      // OpenSea
	BackgroundColor  string           `json:"background_color,omitempty"` // OpenSea

	// Rare Bits
	ID                       *big.Int `json:"id,omitempty"`
	TokenContractAddress     string   `json:"token_contract_address,omitempty"`
	TokenID                  string   `json:"token_id,omitempty"`
	TokenOwnerAddress        string   `json:"token_owner_address,omitempty"`
	ImageURL                 string   `json:"image_url,omitempty"`
	HomeURL                  string   `json:"home_url,omitempty"`
	URL                      string   `json:"url,omitempty"`
	Color                    string   `json:"color,omitempty"`
	Tags                     []string `json:"tags,omitempty"`
	TokenCreatedTimestamp    int64    `json:"token_created_timestamp,omitempty"`
	LastSoldTimestamp        int64    `json:"last_sold_timestamp,omitempty"`
	LastSalePrice            string   `json:"last_sale_price,omitempty"`
	EstimatedValue           string   `json:"estimated_value,omitempty"`
	TokenContractDisplayName string   `json:"token_contract_display_name,omitempty"`
	TokenContractID          string   `json:"token_contract_id,omitempty"`
}

// OpenSeaAttributes is attrebute object defined by OpenSea
type OpenSeaAttributes struct {
	TraitType   string      `json:"trait_type"`
	Value       interface{} `json:"value"`
	DisplayType string      `json:"display_type,omitempty"`
	MaxValue    int64       `json:"max_value,omitempty"`
}

// NewERC721Metadata returns *NewERC721Metadata
func NewERC721Metadata() (*ERC721Metadata, error) {
	ret := new(ERC721Metadata)
	ret.Attributes = nil
	return ret, nil
}
