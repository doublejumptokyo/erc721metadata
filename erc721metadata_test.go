package erc721metadata

import (
	"encoding/json"
	"testing"

	"github.com/cheekybits/is"
)

var desiredJSONString = `
{
    "title": "Asset Metadata",
    "type": "object",
    "properties": {
        "name": {
            "type": "string",
            "description": "Identifies the asset to which this NFT represents"
        },
        "description": {
            "type": "string",
            "description": "Describes the asset to which this NFT represents"
        },
        "image": {
            "type": "string",
            "description": "A URI pointing to a resource with mime type image/* representing the asset to which this NFT represents. Consider making any images at a width between 320 and 1080 pixels and aspect ratio between 1.91:1 and 4:5 inclusive."
        }
    }
}
`

func TestERC721Metadata(t *testing.T) {

	is := is.New(t)
	var err error

	e, err := NewERC721Metadata("Asset Metadata")
	is.NoErr(err)

	e.AddStringProperty("name", "Identifies the asset to which this NFT represents")
	e.AddStringProperty("description", "Describes the asset to which this NFT represents")
	e.AddStringProperty("image", "A URI pointing to a resource with mime type image/* representing the asset to which this NFT represents. Consider making any images at a width between 320 and 1080 pixels and aspect ratio between 1.91:1 and 4:5 inclusive.")

	desired := new(ERC721Metadata)
	err = json.Unmarshal(([]byte)(desiredJSONString), desired)
	is.NoErr(err)

	is.Equal(e, desired)
}
