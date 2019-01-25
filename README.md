[![CircleCI](https://circleci.com/gh/rmanzoku/erc721metadata.svg?style=svg)](https://circleci.com/gh/rmanzoku/erc721metadata)

# erc721metadata
erc721metadata is ERC721 `tokenURI` metadata JSON library implementation by Go.

## Support
- `name`, `description` and `image` defined by [EIP-721](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-721.md) and [EIP-1047](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-1047.md)
- Enhancements metadata `attributes` PR by [EIPs#1071](https://github.com/ethereum/EIPs/pull/1071)
- OpenSea metadata `external_url` and `background_color` defined by [OpenSea metadata](https://docs.opensea.io/docs/2-adding-metadata)
- OpenSea type `attrubutes`. `trait_type`, `value`, `display_type` and `max_value` [OpenSea metadata](https://docs.opensea.io/docs/2-adding-metadata).
- Rare Bits metadata `image_url`, `home_url`, `tags` and `properties` defined by [Rare Bits](https://docs.rarebits.io/v1.0/docs#section-metadata)
- `extra_data` NFT's specific data which can't compare with other NFTs defined here