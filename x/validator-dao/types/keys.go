package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// ModuleName defines the module name
	ModuleName = "validatordao"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_validatordao"
)

const (
	IssueErc20BizKey = "issue-erc20"
)

var (
	AuthorizerBizsKey  = []byte{0x10}
	GranteeAuthBizsKey = []byte{0x11}
)

func GetAuthorizerBizsKey(authorizer sdk.AccAddress) []byte {
	return append(AuthorizerBizsKey, address.MustLengthPrefix(authorizer)...)
}

func GetGranteeAuthBizsKey(grantee sdk.AccAddress) []byte {
	return append(GranteeAuthBizsKey, address.MustLengthPrefix(grantee)...)
}

func ParseGranteeAuthBizsKey(key []byte) sdk.AccAddress {
	addrs := key[1:] // remove prefix bytes
	// addrLen := addrs[0]

	granteeAddr := addrs[1:]
	return granteeAddr
}
