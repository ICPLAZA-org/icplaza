package types

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewCollection creates a new NFT Collection
func NewCollection(cateId string, nfts []Nft) (c Collection) {
	c.CateId = cateId
	for _, nft := range nfts {
		c = c.AddNFT(nft)
	}
	return c
}

// AddNFT adds an NFT to the collection
func (c Collection) AddNFT(nft Nft) Collection {
	c.NFTs = append(c.NFTs, nft)
	return c
}
func (c Collection) AddNFTs(nfts []Nft) Collection {
	for _, nft := range nfts {
		c.NFTs = append(c.NFTs, nft)
	}
	return c
}

func (c Collection) Supply() int {
	return len(c.NFTs)
}

// NewCollection creates a new NFT Collection
func NewCollections(c ...Collection) []Collection {
	return append([]Collection{}, c...)
}

// GetNFT gets a NFT from the collection
func (collection Collection) GetNFT(id string) (nft Nft, err error) {
	nft, found := NFTs(collection.NFTs).Find(id)
	if found {
		return nft, nil
	}
	return Nft{}, sdkerrors.Wrap(ErrUnknownNFT,
		fmt.Sprintf("NFT #%s doesn't exist in collection %s", id, collection.CateId),
	)
}

// ContainsNFT returns whether or not a Collection contains an NFT
func (collection Collection) ContainsNFT(id string) bool {
	_, err := collection.GetNFT(id)
	return err == nil
}

// UpdateNFT updates an NFT from a collection
func (collection Collection) UpdateNFT(nft Nft) (Collection, error) {
	nfts, ok := NFTs(collection.NFTs).Update(nft.GetTokenId(), nft)

	if !ok {
		return collection, sdkerrors.Wrap(ErrUnknownNFT,
			fmt.Sprintf("NFT #%s doesn't exist on collection %s", nft.GetTokenId(), collection.CateId),
		)
	}
	collection.NFTs = nfts
	return collection, nil
}
