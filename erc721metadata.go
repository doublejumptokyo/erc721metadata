package erc721metadata

// ERC721Metadata is ERC721 Metadata JSON
type ERC721Metadata struct {
	Name            string                 `json:"name"`             // ERC721 EIP1047
	Description     string                 `json:"description"`      // ERC721 EIP1047
	Image           string                 `json:"image"`            // ERC721 EIP1047
	Attributes      map[string]interface{} `json:"attributes"`       // PR#1071
	ExternalURL     string                 `json:"external_url"`     // OpenSea
	BackgroundColor string                 `json:"background_color"` // OpenSea
}

// NewERC721ObjectMetadata returns *NewERC721Metadata
func NewERC721Metadata() (*ERC721Metadata, error) {
	ret := new(ERC721Metadata)
	ret.Attributes = map[string]interface{}{}
	return ret, nil
}
