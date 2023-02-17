package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

const (
	// ModuleName defines the module name
	ModuleName = "nftoken"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_nftoken"

	// this line is used by starport scaffolding # ibc/keys/name
	NFTFree   = 0
	NFTLocked = 1
	NFTFrozen = 2
	NFTRent = 3
	// component type
	ComponentTypeOfNFT = 0
	ComponentTypeOfFT  = 1
)

// NFTs are stored as follow:
//
// - Colections: 0x00<cateId_bytes_key> :<Collection>
//
// - Owners: 0x01<address_bytes_key><dcateId_bytes_key>: <Owner>

var (
	CollectionsKeyPrefix = []byte{0x00} // key for NFT collections
	OwnersKeyPrefix      = []byte{0x01} // key for balance of NFTs held by an address

	FixOrderKeyPrefix    = []byte{0x02}
	ActionOrderKeyPrefix = []byte{0x03}
	PrefixPool           = "Pool-value-"
	// 记录每种类型的pool数量
	PrefixPoolCount       = "Pool-count-"
	PrefixPoolTypeCount   = "PoolType-count-"
	PrefixFixedPriceOrder = "Fixed-order-"
	PrefixAuctionOrder    = "Auction-order-"
	PrefixDelegate        = "Delegate-"
	// 重新设计
	PrefixNFT         = "NFT-"
	PrefixOwner       = "owner-"
	PrefixCate        = "cate-"
	PrefixTokenToCate = "token-"
	PrefixCreator     = "create-"
	// 记录最新成交价格-增值税使用
	PrefixLastPrice = "LastPrice-"
	// Transfered-{tokenId}
	PrefixTranferedCNFT = "Transfered-"
	PrefixTokenByName   = "Name-"
)

// this line is used by starport scaffolding # ibc/keys/port

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// GetCollectionKey gets the key of a collection
func GetCollectionKey(denom string) []byte {
	h := tmhash.New()
	_, err := h.Write([]byte(denom))
	if err != nil {
		panic(err)
	}
	bs := h.Sum(nil)

	return append([]byte(CollectionsKeyPrefix), bs...)
}

// SplitOwnerKey gets an address and denom from an owner key
func SplitOwnerKey(key []byte) (sdk.AccAddress, []byte) {
	if len(key) != 54 {
		panic(fmt.Sprintf("unexpected key length %d", len(key)))
	}
	addrLen := key[1]
	addr := key[2 : addrLen+2]
	cateIdHashBz := key[addrLen+2:]
	return addr, cateIdHashBz
}

// GetOwnersKey gets the key prefix for all the collections owned by an account address
func GetOwnersKey(addr sdk.AccAddress) []byte {
	return append([]byte(OwnersKeyPrefix), address.MustLengthPrefix(addr)...)
}

// GetOwnerKey gets the key of a collection owned by an account address
func GetOwnerKey(addr sdk.AccAddress, cateId string) []byte {
	h := tmhash.New()
	_, err := h.Write([]byte(cateId))
	if err != nil {
		panic(err)
	}
	bs := h.Sum(nil)

	return append(GetOwnersKey(addr), bs...)
}

// get

func GetKeyOfTranferedCNFT(tokenId string) []byte {
	return KeyPrefix(PrefixTranferedCNFT + tokenId)
}

func GetKeyOfTokenName(owner string, name string) []byte {
	return KeyPrefix(PrefixTokenByName + owner + "-" + name)
}
