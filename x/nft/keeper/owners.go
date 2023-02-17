package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gauss/gauss/v6/x/nft/types"
)

// GetOwners returns all the Owners ID Collections
func (k Keeper) GetOwners(ctx sdk.Context) (owners []types.Owner) {
	var foundOwners = make(map[string]bool)
	k.IterateOwners(ctx,
		func(owner types.Owner) (stop bool) {
			if _, ok := foundOwners[owner.Address.String()]; !ok {
				foundOwners[owner.Address.String()] = true
				owners = append(owners, owner)
			}
			return false
		},
	)
	return
}

// GetOwner gets all the ID Collections owned by an address
func (k Keeper) GetOwner(ctx sdk.Context, address sdk.AccAddress) (owner types.Owner) {
	var idCollections []types.IDCollection

	k.IterateIDCollections(ctx, types.GetOwnersKey(address),
		func(_ sdk.AccAddress, idCollection types.IDCollection) (stop bool) {
			idCollections = append(idCollections, idCollection)
			return false
		},
	)
	return types.NewOwner(address, idCollections...)
}

// GetOwnerByDenom gets the ID Collection owned by an address of a specific denom
func (k Keeper) GetOwnerByDenom(ctx sdk.Context, owner sdk.AccAddress, denom string) (idCollection types.IDCollection, found bool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.GetOwnerKey(owner, denom))
	if b == nil {
		return types.NewIDCollection(denom, []string{}), false
	}
	k.cdc.MustUnmarshalLengthPrefixed(b, &idCollection)
	return idCollection, true
}

// SetOwnerByDenom sets a collection of NFT IDs owned by an address
func (k Keeper) SetOwnerByDenom(ctx sdk.Context, owner sdk.AccAddress, cateId string, ids []string) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetOwnerKey(owner, cateId)

	var idCollection types.IDCollection
	idCollection.CateId = cateId
	idCollection.Ids = ids

	store.Set(key, k.cdc.MustMarshalLengthPrefixed(&idCollection))
}

// SetOwner sets an entire Owner
func (k Keeper) SetOwner(ctx sdk.Context, owner types.Owner) {
	for _, idCollection := range owner.IDCollections {
		k.SetOwnerByDenom(ctx, owner.Address, idCollection.CateId, idCollection.Ids)
	}
}

// SetOwners sets all Owners
func (k Keeper) SetOwners(ctx sdk.Context, owners []types.Owner) {
	for _, owner := range owners {
		k.SetOwner(ctx, owner)
	}
}

// IterateIDCollections iterates over the IDCollections by Owner and performs a function
func (k Keeper) IterateIDCollections(ctx sdk.Context, prefix []byte,
	handler func(owner sdk.AccAddress, idCollection types.IDCollection) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var idCollection types.IDCollection
		k.cdc.MustUnmarshalLengthPrefixed(iterator.Value(), &idCollection)

		owner, _ := types.SplitOwnerKey(iterator.Key())
		if handler(owner, idCollection) {
			break
		}
	}
}

// IterateOwners iterates over all Owners and performs a function
func (k Keeper) IterateOwners(ctx sdk.Context, handler func(owner types.Owner) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.OwnersKeyPrefix))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var owner types.Owner

		address, _ := types.SplitOwnerKey(iterator.Key())
		owner = k.GetOwner(ctx, address)

		if handler(owner) {
			break
		}
	}
}

// SwapOwners swaps the owners of a NFT ID
func (k Keeper) SwapOwners(ctx sdk.Context, denom string, id string, oldAddress sdk.AccAddress, newAddress sdk.AccAddress) (err error) {
	oldOwnerIDCollection, found := k.GetOwnerByDenom(ctx, oldAddress, denom)
	if !found {
		return sdkerrors.Wrap(types.ErrUnknownCollection,
			fmt.Sprintf("id collection %s doesn't exist for owner %s", denom, oldAddress),
		)
	}
	oldOwnerIDCollection, err = oldOwnerIDCollection.DeleteID(id)
	if err != nil {
		return err
	}
	k.SetOwnerByDenom(ctx, oldAddress, denom, oldOwnerIDCollection.Ids)

	newOwnerIDCollection, found := k.GetOwnerByDenom(ctx, newAddress, denom)
	if !found {
		newOwnerIDCollection = types.NewIDCollection(denom, []string{})
	}
	newOwnerIDCollection = newOwnerIDCollection.AddID(id)
	k.SetOwnerByDenom(ctx, newAddress, denom, newOwnerIDCollection.Ids)
	return nil
}
