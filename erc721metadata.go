package erc721metadata

// ERC721Metadata is ERC721 Metadata JSON Schema
type ERC721Metadata struct {
	Title      string              `json:"title"`
	Type       string              `json:"type"`
	Properties map[string]Property `json:"properties"`
}

// Property is Property define by ERC721 Metadata
type Property struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

// NewERC721Metadata is alias to NewERC721ObjectMetadata
var NewERC721Metadata = NewERC721ObjectMetadata

// NewERC721ObjectMetadata returns *NewERC721Metadata (type: object)
func NewERC721ObjectMetadata(title string) (*ERC721Metadata, error) {
	return &ERC721Metadata{
		Title:      title,
		Type:       "object",
		Properties: make(map[string]Property),
	}, nil
}

// AddProperty is add property to ERC721Metadata.Properties
func (e *ERC721Metadata) AddProperty(propertyName string, description string) error {
	p := Property{
		Type:        "string",
		Description: description,
	}

	e.Properties[propertyName] = p
	return nil
}
