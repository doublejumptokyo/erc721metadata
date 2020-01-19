package erc721metadata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	Timestamp int64                  `json:"timestamp,omitempty"`

	Language string `json:"language,omitempty"`
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

// FetchERC721Metadata returns *NewERC721Metadata fetch Metadata from tokenURI
func FetchERC721Metadata(tokenURI string) (*ERC721Metadata, error) {
	client := new(http.Client)
	req, err := http.NewRequest("GET", tokenURI, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status: %d, msg: %s", resp.StatusCode, string(body))
	}

	return UnmarshalERC721Metadata(body)
}

// NewERC721Metadata returns *NewERC721Metadata
func NewERC721Metadata() (*ERC721Metadata, error) {
	ret := new(ERC721Metadata)
	ret.Attributes = nil
	ret.Timestamp = time.Now().Unix()
	ret.Language = "en"
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

// SetAttributes convert to JSON.RawMessage and set to Attributes
func (e *ERC721Metadata) SetAttributes(attributes interface{}) error {
	byte, err := json.Marshal(attributes)
	if err != nil {
		return err
	}
	raw := json.RawMessage(byte)
	e.Attributes = &raw
	return nil
}
