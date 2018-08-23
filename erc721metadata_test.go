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

  "background_color": "00FFFF",
  "external_url": "https://cryptopuff.io/3",

  "attributes": {
    "level": 3,
    "weapon_power": 55,
    "personality": "sad",
    "stamina": 11.7
  }
}
`

type attributes struct {
	Level       int64   `json:"level"`
	WeaponPower int64   `json:"weapon_power"`
	Personality string  `json:"personality"`
	Stamina     float64 `json:"stamina"`
}

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

	a := attributes{
		Level:       3,
		WeaponPower: 55,
		Personality: "sad",
		Stamina:     11.7,
	}
	aByte, err := json.Marshal(a)
	is.NoErr(err)
	raw := json.RawMessage(aByte)
	e.Attributes = &raw

	ej, err := json.Marshal(e)
	is.NoErr(err)

	desired := new(ERC721Metadata)
	err = json.Unmarshal(([]byte)(desiredJSONString), desired)
	is.NoErr(err)

	dj, err := json.Marshal(desired)
	is.NoErr(err)

	is.Equal(string(dj), string(ej))
	//fmt.Println(string(ej))
}

var desiredOpenSea = `
{
  "attributes": [
    {
      "trait_type": "level",
      "value": 2
    },
    {
      "max_value": 55,
      "trait_type": "stamina",
      "value": 2.3
    },
    {
      "trait_type": "personality",
      "value": "sad"
    },
    {
      "display_type": "boost_number",
      "trait_type": "puff_power",
      "value": 40
    },
    {
      "display_type": "boost_percentage",
      "trait_type": "stamina_increase",
      "value": 10
    },
    {
      "display_type": "number",
      "trait_type": "generation",
      "value": 2
    }
  ],
  "description": "Generic puff description. Check out this link to ethmoji.io\nHave a new line or two\n\n This really should be customized.",
  "externalUrl": "https://cryptopuff.io/1",
  "image": "https://storage.googleapis.com/opensea-prod.appspot.com/puffs/1.png",
  "name": "Jessica McDonald"
}
`

func TestOpenSea(t *testing.T) {
	is := is.New(t)
	var err error

	e, err := NewERC721Metadata()
	is.NoErr(err)

	// string values
	e.Name = "Jessica McDonald"
	e.Image = "https://storage.googleapis.com/opensea-prod.appspot.com/puffs/1.png"
	e.Description = "Generic puff description. Check out this link to ethmoji.io\nHave a new line or two\n\n This really should be customized."
	e.ExternalURLCamel = "https://cryptopuff.io/1"

	a := []OpenSeaAttributes{
		OpenSeaAttributes{
			TraitType: "level",
			Value:     2,
		},
		OpenSeaAttributes{
			MaxValue:  55,
			TraitType: "stamina",
			Value:     2.3,
		},
		OpenSeaAttributes{
			TraitType: "personality",
			Value:     "sad",
		},
		OpenSeaAttributes{
			DisplayType: "boost_number",
			TraitType:   "puff_power",
			Value:       40,
		},
		OpenSeaAttributes{
			DisplayType: "boost_percentage",
			TraitType:   "stamina_increase",
			Value:       10,
		},
		OpenSeaAttributes{
			DisplayType: "number",
			TraitType:   "generation",
			Value:       2,
		},
	}

	aByte, err := json.Marshal(a)
	is.NoErr(err)
	raw := json.RawMessage(aByte)
	e.Attributes = &raw

	ej, err := json.Marshal(e)
	is.NoErr(err)

	desired := new(ERC721Metadata)
	err = json.Unmarshal(([]byte)(desiredOpenSea), desired)
	is.NoErr(err)

	da := new([]OpenSeaAttributes)
	err = json.Unmarshal(([]byte)(*desired.Attributes), da)
	is.NoErr(err)

	daByte, err := json.Marshal(da)
	is.NoErr(err)
	raw = json.RawMessage(daByte)
	desired.Attributes = &raw

	dj, err := json.Marshal(desired)
	is.NoErr(err)

	is.Equal(string(dj), string(ej))
	//fmt.Println(string(ej))
}
