package erc721metadata

import (
	"encoding/json"
	"testing"

	"github.com/cheekybits/is"
)

var desiredJSONString = `
{
  "name": "Dave McPufflestein",
  "image": "https://storage.googleapis.com/opensea-prod.appspot.com/puffs/3.png",
  "description": "Generic puff description. This really should be customized.",
  "external_url": "https://cryptopuff.io/3",
  "background_color": "00FFFF"
}
`

func TestERC721Metadata(t *testing.T) {

	is := is.New(t)
	var err error

	e, err := NewERC721Metadata()
	is.NoErr(err)

	e.Name = "Dave McPufflestein"
	e.Image = "https://storage.googleapis.com/opensea-prod.appspot.com/puffs/3.png"
	e.Description = "Generic puff description. This really should be customized."
	e.ExternalURL = "https://cryptopuff.io/3"
	e.BackgroundColor = "00FFFF"

	desired := new(ERC721Metadata)
	err = json.Unmarshal(([]byte)(desiredJSONString), desired)
	is.NoErr(err)

	is.Equal(e, desired)
}
