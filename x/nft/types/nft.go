package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sort"
	"strings"
)

// NewNft creates a new NFT instance
func NewNft(tokenId, cateId, owner, creator, companyUri, tokenUri string, valueAddedTax, copyrightTax sdk.Dec, tokenStatus uint32) Nft {
	return Nft{
		TokenId:       tokenId,
		CateId:        cateId,
		Owner:         owner,
		Creator:       creator,
		CompanyUri:    companyUri,
		TokenUri:      tokenUri,
		ValueAddedTax: &valueAddedTax,
		CopyrightTax:  &copyrightTax,
		TokenStatus:   tokenStatus,
	}
}

// ----------------------------------------------------------------------------
// NFT

// NFTs define a list of NFT
type NFTs []Nft

// NewNFTs creates a new set of NFTs
func NewNFTs(nfts ...Nft) NFTs {
	if len(nfts) == 0 {
		return NFTs{}
	}
	return NFTs(nfts)
}

// Find returns the searched collection from the set
func (nfts NFTs) Find(id string) (nft Nft, found bool) {
	index := nfts.find(id)
	if index == -1 {
		return nft, false
	}
	return nfts[index], true
}

// Update removes and replaces an NFT from the set
func (nfts NFTs) Update(id string, nft Nft) (NFTs, bool) {
	index := nfts.find(id)
	if index == -1 {
		return nfts, false
	}

	return append(append(nfts[:index], nft), nfts[index+1:]...), true
}

func (nfts NFTs) find(id string) int {
	return FindUtil(nfts, id)
}

// Findable and Sort interfaces
func (nfts NFTs) ElAtIndex(index int) string { return nfts[index].GetTokenId() }
func (nfts NFTs) Len() int                   { return len(nfts) }
func (nfts NFTs) Less(i, j int) bool {
	return strings.Compare(nfts[i].GetTokenId(), nfts[j].GetTokenId()) == -1
}
func (nfts NFTs) Swap(i, j int) { nfts[i], nfts[j] = nfts[j], nfts[i] }

var _ sort.Interface = NFTs{}

// Sort is a helper function to sort the set of coins in place
func (nfts NFTs) Sort() NFTs {
	sort.Sort(nfts)
	return nfts
}

func ValidateTokenID(tokenID string) error {
	//tokenID = strings.TrimSpace(tokenID)
	//if len(tokenID) < MinDenomLen || len(tokenID) > MaxDenomLen {
	//	return sdkerrors.Wrapf(ErrInvalidTokenID, "invalid tokenID %s, only accepts value [%d, %d]", tokenID, MinDenomLen, MaxDenomLen)
	//}
	//if !IsBeginWithAlpha(tokenID) || !IsAlphaNumeric(tokenID) {
	//	return sdkerrors.Wrapf(ErrInvalidTokenID, "invalid tokenID %s, only accepts alphanumeric characters,and begin with an english letter", tokenID)
	//}
	return nil
}

func ValidateTokenURI(tokenURI string) error {
	//if len(tokenURI) > MaxTokenURILen {
	//	return sdkerrors.Wrapf(ErrInvalidTokenURI, "invalid tokenURI %s, only accepts value [0, %d]", tokenURI, MaxTokenURILen)
	//}
	return nil
}
