package erc721metadata

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/cheekybits/is"
)

var desiredJSONString = `
{
  "name": "Dave McPufflestein",
  "image": "https://storage.googleapis.com/opensea-prod.appspot.com/puffs/3.png",
  "description": "Generic puff description. This really should be customized.",

  "background_color": "00FFFF",
  "external_url": "https://cryptopuff.io/3",

  "attributes": {
    "level": 3,
    "weapon_power": 55,
    "personality": "sad",
    "stamina": 11.7
  },
  "tags": [
    "origin",
    "shiny"
  ],
  "ID": 9223372036854775808,
  "token_created_timestamp": 1534914870

}
`

// type attributes struct {
// 	Level       int64   `json:"level"`
// 	WeaponPower int64   `json:"weapon_power"`
// 	Personality string  `json:"personality"`
// 	Stamina     float64 `json:"stamina"`
// }

func TestERC721Metadata(t *testing.T) {

	is := is.New(t)
	var err error

	e, err := NewERC721Metadata()
	is.NoErr(err)

	// string values
	e.Name = "Dave McPufflestein"
	e.Image = "https://storage.googleapis.com/opensea-prod.appspot.com/puffs/3.png"
	e.Description = "Generic puff description. This really should be customized."
	e.ExternalURL = "https://cryptopuff.io/3"
	e.BackgroundColor = "00FFFF"

	// attributes object
	e.Attributes["level"] = 3
	e.Attributes["weapon_power"] = 55
	e.Attributes["personality"] = "sad"
	e.Attributes["stamina"] = 11.7

	// list of string
	e.Tags = append(e.Tags, "origin")
	e.Tags = append(e.Tags, "shiny")

	// big int
	var ok bool
	e.ID, ok = new(big.Int).SetString("9223372036854775808", 10)
	is.OK(ok)
	e.TokenCreatedTimestamp = 1534914870

	ej, err := json.Marshal(e)
	is.NoErr(err)

	desired := new(ERC721Metadata)
	err = json.Unmarshal(([]byte)(desiredJSONString), desired)
	is.NoErr(err)

	dj, err := json.Marshal(desired)
	is.NoErr(err)

	is.Equal(string(dj), string(ej))
	fmt.Println(string(ej))
}
