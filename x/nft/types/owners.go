package types

import (
	"fmt"
	"sort"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// IDCollection defines a set of nft ids that belong to a specific
// collection
// SortedStringArray is an array of strings whose sole purpose is to help with find
type SortedStringArray []string

// String is the string representation
func (sa SortedStringArray) String() string { return strings.Join(sa[:], ",") }

// NewIDCollection creates a new IDCollection instance
func NewIDCollection(cateId string, ids []string) IDCollection {
	return IDCollection{
		CateId: strings.TrimSpace(cateId),
		Ids:    SortedStringArray(ids).Sort(),
	}
}

// Exists determines whether an ID is in the IDCollection
func (idCollection IDCollection) Exists(id string) (exists bool) {
	index := SortedStringArray(idCollection.Ids).find(id)
	return index != -1
}

// AddID adds an ID to the idCollection
func (idCollection IDCollection) AddID(id string) IDCollection {
	idCollection.Ids = SortedStringArray(append(idCollection.Ids, id)).Sort()
	return idCollection
}

// DeleteID deletes an ID from an ID Collection
func (idCollection IDCollection) DeleteID(id string) (IDCollection, error) {
	index := SortedStringArray(idCollection.Ids).find(id)
	if index == -1 {
		return idCollection, sdkerrors.Wrap(ErrUnknownNFT,
			fmt.Sprintf("ID #%s doesn't exist on ID Collection %s", id, idCollection.CateId),
		)
	}

	idCollection.Ids = append(idCollection.Ids[:index], idCollection.Ids[index+1:]...)

	return idCollection, nil
}

// Supply gets the total supply of NFTIDs of a balance
func (idCollection IDCollection) Supply() int {
	return len(idCollection.Ids)
}

// ----------------------------------------------------------------------------
// Owners

// IDCollections is an array of ID Collections whose sole purpose is for find
type IDCollections []IDCollection

// String follows stringer interface
func (idCollections IDCollections) String() string {
	if len(idCollections) == 0 {
		return ""
	}

	out := ""
	for _, idCollection := range idCollections {
		out += fmt.Sprintf("%v\n", idCollection.String())
	}
	return out[:len(out)-1]
}

// Append appends IDCollections to IDCollections
func (idCollections IDCollections) Append(idCollections2 ...IDCollection) IDCollections {
	return append(idCollections, idCollections2...).Sort()
}
func (idCollections IDCollections) find(denom string) int {
	return FindUtil(idCollections, denom)
}

// NewOwner creates a new Owner
func NewOwner(owner sdk.AccAddress, idCollections ...IDCollection) Owner {
	return Owner{
		Address:       owner,
		IDCollections: idCollections,
	}
}

// Supply gets the total supply of an Owner
func (owner Owner) Supply() int {
	total := 0
	for _, idCollection := range owner.IDCollections {
		total += idCollection.Supply()
	}
	return total
}

// GetIDCollection gets the IDCollection from the owner
func (owner Owner) GetIDCollection(denom string) (IDCollection, bool) {
	index := IDCollections(owner.IDCollections).find(denom)
	if index == -1 {
		return IDCollection{}, false
	}
	return owner.IDCollections[index], true
}

// UpdateIDCollection updates the ID Collection of an owner
func (owner Owner) UpdateIDCollection(idCollection IDCollection) (Owner, error) {
	denom := idCollection.CateId
	index := IDCollections(owner.IDCollections).find(denom)
	if index == -1 {
		return owner, sdkerrors.Wrap(ErrUnknownCollection,
			fmt.Sprintf("ID Collection %s doesn't exist for owner %s", denom, owner.Address),
		)
	}

	owner.IDCollections = append(append(owner.IDCollections[:index], idCollection), owner.IDCollections[index+1:]...)

	return owner, nil
}

// DeleteID deletes an ID from an owners ID Collection
func (owner Owner) DeleteID(denom string, id string) (Owner, error) {
	idCollection, found := owner.GetIDCollection(denom)
	if !found {
		return owner, sdkerrors.Wrap(ErrUnknownNFT,
			fmt.Sprintf("ID #%s doesn't exist in ID Collection %s", id, denom),
		)
	}
	idCollection, err := idCollection.DeleteID(id)
	if err != nil {
		return owner, err
	}
	owner, err = owner.UpdateIDCollection(idCollection)
	if err != nil {
		return owner, err
	}
	return owner, nil
}

func (sa SortedStringArray) find(el string) (idx int) {
	return FindUtil(sa, el)
}

//-----------------------------------------------------------------------------
// Sort and Findable interface for SortedStringArray

func (sa SortedStringArray) ElAtIndex(index int) string { return sa[index] }
func (sa SortedStringArray) Len() int                   { return len(sa) }
func (sa SortedStringArray) Less(i, j int) bool {
	return strings.Compare(sa[i], sa[j]) == -1
}
func (sa SortedStringArray) Swap(i, j int) {
	sa[i], sa[j] = sa[j], sa[i]
}

var _ sort.Interface = SortedStringArray{}

// Sort is a helper function to sort the set of strings in place
func (sa SortedStringArray) Sort() SortedStringArray {
	sort.Sort(sa)
	return sa
}

// Sort and Findable interface for IDCollections

func (idCollections IDCollections) ElAtIndex(index int) string { return idCollections[index].CateId }
func (idCollections IDCollections) Len() int                   { return len(idCollections) }
func (idCollections IDCollections) Less(i, j int) bool {
	return strings.Compare(idCollections[i].CateId, idCollections[j].CateId) == -1
}
func (idCollections IDCollections) Swap(i, j int) {
	idCollections[i].CateId, idCollections[j].CateId = idCollections[j].CateId, idCollections[i].CateId
}

var _ sort.Interface = IDCollections{}

// Sort is a helper function to sort the set of strings in place
func (idCollections IDCollections) Sort() IDCollections {
	sort.Sort(idCollections)
	return idCollections
}
