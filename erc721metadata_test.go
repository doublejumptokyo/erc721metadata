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
	err = e.SetAttributes(a)
	is.NoErr(err)

	desired := new(ERC721Metadata)
	err = json.Unmarshal(([]byte)(desiredJSONString), desired)
	is.NoErr(err)

	dj, err := json.Marshal(desired)
	is.NoErr(err)

	ej, err := json.Marshal(desired)
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
  "external_url": "https://cryptopuff.io/1",
  "image": "https://storage.googleapis.com/opensea-prod.appspot.com/puffs/1.png",
  "name": "Jessica McDonald"
}
`

func TestOpenSea(t *testing.T) {
	is := is.New(t)
	var err error

	e, err := NewERC721Metadata()
	is.NoErr(err)
	tmp := *e
	desired := &tmp

	// string values
	e.Name = "Jessica McDonald"
	e.Image = "https://storage.googleapis.com/opensea-prod.appspot.com/puffs/1.png"
	e.Description = "Generic puff description. Check out this link to ethmoji.io\nHave a new line or two\n\n This really should be customized."
	e.ExternalURL = "https://cryptopuff.io/1"

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

	err = e.SetAttributes(a)
	is.NoErr(err)

	ej, err := json.Marshal(e)
	is.NoErr(err)

	err = json.Unmarshal(([]byte)(desiredOpenSea), desired)
	is.NoErr(err)

	da := new([]OpenSeaAttributes)
	err = json.Unmarshal(([]byte)(*desired.Attributes), da)
	is.NoErr(err)

	err = desired.SetAttributes(da)
	is.NoErr(err)

	dj, err := json.Marshal(desired)
	is.NoErr(err)

	is.Equal(string(dj), string(ej))
	//fmt.Println(string(ej))
}

var desiredRareBits = `
{
  "name": "Robot token #14",
  "image_url": "https://www.robotgame.com/images/14.png",
  "home_url": "https://www.robotgame.com/robots/14.html",
  "description": "This is the amazing Robot #14, please buy me!",
  "properties": [
    {"key": "generation", "value": 4, "type": "integer"},
    {"key": "cooldown", "value": "slow", "type": "string"}
  ],
  "tags": ["red","rare","fire"]
}
`

func TestRareBit(t *testing.T) {
	is := is.New(t)
	var err error

	e, err := NewERC721Metadata()
	is.NoErr(err)
	tmp := *e
	desired := &tmp

	e.Name = "Robot token #14"
	e.ImageURL = "https://www.robotgame.com/images/14.png"
	e.HomeURL = "https://www.robotgame.com/robots/14.html"
	e.Description = "This is the amazing Robot #14, please buy me!"

	e.Properties = []*RareBitsProperty{
		&RareBitsProperty{
			Key:   "generation",
			Value: 4,
			Type:  "integer",
		},
		&RareBitsProperty{
			Key:   "cooldown",
			Value: "slow",
			Type:  "string",
		},
	}

	e.Tags = append(e.Tags, "red")
	e.Tags = append(e.Tags, "rare")
	e.Tags = append(e.Tags, "fire")

	ej, err := json.Marshal(e)
	is.NoErr(err)

	err = json.Unmarshal(([]byte)(desiredRareBits), desired)
	is.NoErr(err)

	dj, err := json.Marshal(desired)
	is.NoErr(err)

	is.Equal(string(dj), string(ej))
	//fmt.Println(string(ej))
}
