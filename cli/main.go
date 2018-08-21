package main

import (
	"encoding/json"
	"fmt"

	"github.com/rmanzoku/erc721metadata"
)

func main() {
	erc, _ := erc721metadata.NewERC721Metadata("Hello")

	erc.AddProperty("name", "Identifies the asset to which this NFT represents")

	outputJSON, _ := json.Marshal(erc)
	fmt.Println(string(outputJSON))
}
