package erc721metadata

import (
	"encoding/json"
)

// ERC721Metadata is ERC721 Metadata JSON
type ERC721Metadata struct {
	// ERC721 EIP1047
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image,omitempty"`

	// OpenSea
	Attributes      *json.RawMessage `json:"attributes,omitempty"`       // EIPs#1071
	ExternalURL     string           `json:"external_url,omitempty"`     // OpenSea
	BackgroundColor string           `json:"background_color,omitempty"` // OpenSea

	// Rare Bits
	ImageURL   string              `json:"image_url,omitempty"`
	HomeURL    string              `json:"home_url,omitempty"`
	Tags       []string            `json:"tags,omitempty"`
	Properties []*RareBitsProperty `json:"properties,omitempty"`
}

// OpenSeaAttributes is attrebute object defined by OpenSea
type OpenSeaAttributes struct {
	TraitType   string      `json:"trait_type"`
	Value       interface{} `json:"value"`
	DisplayType string      `json:"display_type,omitempty"`
	MaxValue    int64       `json:"max_value,omitempty"`
}

// RareBitsProperty is property object defined by RareBits
type RareBitsProperty struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}

// NewERC721Metadata returns *NewERC721Metadata
func NewERC721Metadata() (*ERC721Metadata, error) {
	ret := new(ERC721Metadata)
	ret.Attributes = nil
	return ret, nil
}
