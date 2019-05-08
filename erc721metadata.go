package erc721metadata

import (
	"encoding/json"
	"time"
)

// ERC721Metadata is ERC721 Metadata JSON
type ERC721Metadata struct {
	// ERC721 EIP1047
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
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

	// MCH
	ExtraData map[string]interface{} `json:"extra_data,omitempty"`
	Timestamp int64                  `json:"timestamp"`
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
	ret.Timestamp = time.Now().Unix()
	return ret, nil
}

// UnmarshalERC721Metadata unmarshals ERC721Metadata
func UnmarshalERC721Metadata(data []byte) (*ERC721Metadata, error) {
	var r = &ERC721Metadata{}
	err := json.Unmarshal(data, r)
	return r, err
}

// Marshal marshals ERC721Metadata
func (e *ERC721Metadata) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

// SetRawAttributes convert to JSON.RawMessage and set to Attributes
func (e *ERC721Metadata) SetRawAttributes(attributes map[string]interface{}) error {
	byte, err := json.Marshal(attributes)
	if err != nil {
		return err
	}
	raw := json.RawMessage(byte)
	e.Attributes = &raw
	return nil
}
