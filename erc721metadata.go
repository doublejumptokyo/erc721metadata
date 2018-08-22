package erc721metadata

// ERC721Metadata is ERC721 Metadata JSON Schema
type ERC721Metadata struct {
	Name            string `json:"name"`             // ERC721
	Description     string `json:"description"`      // ERC721
	Image           string `json:"image"`            // ERC721
	ExternalURL     string `json:"external_url"`     // OpenSea
	BackgroundColor string `json:"background_color"` // OpenSea
}

// NewERC721ObjectMetadata returns *NewERC721Metadata (type: object)
func NewERC721Metadata() (*ERC721Metadata, error) {
	return &ERC721Metadata{}, nil
}
