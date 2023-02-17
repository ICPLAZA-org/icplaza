package types

import "fmt"

const (
	// ModuleName defines the module name
	ModuleName = "blindbox"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_blindbox"
)

var (
	PrefixKeyOfPool             = "Pool-%s"
	PrefixKeyOfGroup            = "Group-%s"
	PrefixKeyOfBox              = "Box-%s-%s"
	PrefixKeyOfCreatorToPoolId  = "CreatorToPoolId-%s-%s"
	PrefixKeyOfPoolIdToGroupId  = "PoolIdToGroupId-%s-%s"
	PrefixKeyOfCreatorToGroupId = "CreatorToGroupId-%s-%s"
	PrefixKeyOfExpiredGroupId   = "Expired-group-%s"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func KeyOfPool(poolId string) []byte {
	str := fmt.Sprintf(PrefixKeyOfPool, poolId)
	return KeyPrefix(str)
}

func KeyOfCreatorToPoolId(creator string, poolId string) []byte {
	str := fmt.Sprintf(PrefixKeyOfCreatorToPoolId, creator, poolId)
	return KeyPrefix(str)
}

func KeyOfGroup(groupId string) []byte {
	str := fmt.Sprintf(PrefixKeyOfGroup, groupId)
	return KeyPrefix(str)
}

func KeyOfPoolIdToGroupId(poolId string, groupId string) []byte {
	str := fmt.Sprintf(PrefixKeyOfPoolIdToGroupId, poolId, groupId)
	return KeyPrefix(str)
}

func KeyOfCreatorToGroupId(creator string, groupId string) []byte {
	str := fmt.Sprintf(PrefixKeyOfCreatorToGroupId, creator, groupId)
	return KeyPrefix(str)
}

//func KeyOfGroupCreator(poolId string, creator string) []byte {
//	return KeyPrefix(PrefixKeyOfGroup + poolId + "-" + creator + "-")
//}
//
func KeyOfBox(groupId string, boxId string) []byte {
	str := fmt.Sprintf(PrefixKeyOfBox, groupId, boxId)
	return KeyPrefix(str)
}

func KeyOfExpiredGroupId(groupId string) []byte {
	str := fmt.Sprintf(PrefixKeyOfExpiredGroupId, groupId)
	return KeyPrefix(str)
}
