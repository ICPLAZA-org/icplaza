package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gauss/gauss/v6/x/blindbox/types"
	"strconv"
	"strings"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateBox(goCtx context.Context, msg *types.MsgCreateBox) (*types.MsgCreateBoxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := k.CreateGroup(ctx, *msg)
	if err != nil {
		return nil, err
	}
	if len(msg.FixedNfts) > 0 {
		err := k.LockNFTs(ctx, msg.FixedNfts, msg.Creator)
		if err != nil {
			return nil, err
		}
	}
	if len(msg.RandomNfts) > 0 {
		err := k.LockNFTs(ctx, msg.RandomNfts, msg.Creator)
		if err != nil {
			return nil, err
		}
	}

	boxes := make([]*types.Box, 0)
	for i := 0; i < int(msg.BoxSize); i++ {
		box := &types.Box{
			BoxId:   strconv.Itoa(i + 1),
			GroupId: msg.GroupId,
			Opened:  false,
			Nfts:    nil,
		}
		boxes = append(boxes, box)
	}
	// save all box
	k.CreateBoxes(ctx, boxes)
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateBox,
			sdk.NewAttribute("group_id", msg.GroupId),
			sdk.NewAttribute("box_size", strconv.Itoa(int(msg.BoxSize))),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})
	return &types.MsgCreateBoxResponse{}, nil
}

// lock nft list
func (k Keeper) LockNFTs(ctx sdk.Context, tokens []string, creator string) error {
	for _, token := range tokens {
		nft, err := k.nftKeeper.GetNFT(ctx, "", token)
		if err != nil {
			return err
		}
		if nft.Owner != creator {
			return sdkerrors.Wrapf(types.ErrInvalidNFTOwner, "%s", creator)
		}
		// locked
		err = k.nftKeeper.LockNFT(ctx, nft.CateId, nft.TokenId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.Keeper.CreatePool(ctx, *msg)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreatePool,
			sdk.NewAttribute("pool_id", msg.PoolId),
			sdk.NewAttribute("fee_rate", msg.FeeRate.String()),
			sdk.NewAttribute("fee_address", msg.FeeAddress),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})
	return &types.MsgCreatePoolResponse{}, nil
}
func (k msgServer) OpenBox(goCtx context.Context, msg *types.MsgOpenBox) (*types.MsgOpenBoxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.Keeper.OpenBox(ctx, msg.GroupId, msg.BoxIds, msg.Creator)

	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeOpenBox,
			sdk.NewAttribute("group_id", msg.GroupId),
			sdk.NewAttribute("box_ids", strings.Join(msg.BoxIds, ",")),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgOpenBoxResponse{}, nil
}
func (k msgServer) RemovePool(goCtx context.Context, msg *types.MsgRemovePool) (*types.MsgRemovePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.Keeper.RemovePool(ctx, *msg)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRemovePool,
			sdk.NewAttribute("pool_id", msg.PoolId),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgRemovePoolResponse{}, nil
}
func (k msgServer) RevokeBoxGroup(goCtx context.Context, msg *types.MsgRevokeBoxGroup) (*types.MsgRevokeBoxGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	groupInfo, err := k.Keeper.GetGroupInfo(ctx, msg.GroupId)
	if err != nil {
		return nil, err
	}
	if groupInfo.Creator != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCreator, "")
	}
	if groupInfo.OpenSize == groupInfo.BoxSize {
		return nil, sdkerrors.Wrapf(types.ErrRevoke, "group(%s) already revoked", groupInfo.GroupId)
	}

	err = k.RevokeGroup(ctx, groupInfo)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRevokeBox,
			sdk.NewAttribute("group_id", msg.GroupId),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgRevokeBoxGroupResponse{}, nil
}
func (k msgServer) UpdatePool(goCtx context.Context, msg *types.MsgUpdatePool) (*types.MsgUpdatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.Keeper.UpdatePool(ctx, *msg)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdatePool,
			sdk.NewAttribute("pool_id", msg.PoolId),
			sdk.NewAttribute("fee_rate", msg.FeeRate.String()),
			sdk.NewAttribute("fee_address", msg.FeeAddress),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})
	return &types.MsgUpdatePoolResponse{}, nil
}
