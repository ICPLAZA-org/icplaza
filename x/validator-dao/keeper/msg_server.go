package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/gauss/gauss/v6/x/validator-dao/types"
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

func (k msgServer) ReqAuthorization(goCtx context.Context, msg *types.MsgReqAuthorization) (*types.MsgReqAuthorizationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	granteeAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	authorizerAddr, err := sdk.AccAddressFromBech32(msg.Authorizer)
	if err != nil {
		return nil, err
	}
	if !k.Keeper.IsAuthorizer(ctx, authorizerAddr) {
		return nil, sdkerrors.Wrapf(types.ErrNoValidatorFound, "%s is not authorizer.", msg.Authorizer)
	}
	
	err = k.Keeper.ReqAuthorization(ctx, granteeAddr, authorizerAddr, msg.BizName, msg.Fee)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeReqAuthorization,
			sdk.NewAttribute(types.AttributeKeyAuthorizer, msg.Authorizer),
			sdk.NewAttribute(types.AttributeKeyBiz, msg.BizName),
			sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Fee.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})
	
	return &types.MsgReqAuthorizationResponse{}, nil
}

func (k msgServer) WithdrawAuthorization(goCtx context.Context, msg *types.MsgWithdrawAuthorization) (*types.MsgWithdrawAuthorizationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	granteeAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	authorizerAddr, err := sdk.AccAddressFromBech32(msg.Authorizer)
	if err != nil {
		return nil, err
	}
	if !k.Keeper.IsAuthorizer(ctx, authorizerAddr) {
		return nil, sdkerrors.Wrapf(types.ErrNoValidatorFound, "%s is not authorizer.", msg.Authorizer)
	}
	
	balance, err := k.Keeper.WithdrawAuthorization(ctx, granteeAddr, authorizerAddr, msg.BizName)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeWithdrawAuthorization,
			sdk.NewAttribute(types.AttributeKeyAuthorizer, msg.Authorizer),
			sdk.NewAttribute(types.AttributeKeyBiz, msg.BizName),
			sdk.NewAttribute(sdk.AttributeKeyAmount, balance.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})
	
	return &types.MsgWithdrawAuthorizationResponse{}, nil
}

func (k msgServer) AddAuthBiz(goCtx context.Context, msg *types.MsgAddAuthBiz) (*types.MsgAddAuthBizResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	authorizerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	if !k.Keeper.IsAuthorizer(ctx, authorizerAddr) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidAuthorizer, "%s is not a authorizer", msg.Creator)
	}

	err = k.Keeper.AddAuthBiz(ctx, authorizerAddr, msg.BizName, msg.Fee)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAddAuthBiz,
			sdk.NewAttribute(types.AttributeKeyBiz, msg.BizName),
			sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Fee.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgAddAuthBizResponse{}, nil
}

func (k msgServer) UpdateAuthBiz(goCtx context.Context, msg *types.MsgUpdateAuthBiz) (*types.MsgUpdateAuthBizResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	authorizerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	if !k.Keeper.IsAuthorizer(ctx, authorizerAddr) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidAuthorizer, "%s is not a authorizer", msg.Creator)
	}

	err = k.Keeper.UpdateAuthBiz(ctx, authorizerAddr, msg.BizName, msg.Fee)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateAuthBiz,
			sdk.NewAttribute(types.AttributeKeyBiz, msg.BizName),
			sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Fee.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgUpdateAuthBizResponse{}, nil
}

func (k msgServer) RemoveAuthBiz(goCtx context.Context, msg *types.MsgRemoveAuthBiz) (*types.MsgRemoveAuthBizResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	authorizerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	err = k.Keeper.RemoveAuthBiz(ctx, authorizerAddr, msg.BizName)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRemoveAuthBiz,
			sdk.NewAttribute(types.AttributeKeyBiz, msg.BizName),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgRemoveAuthBizResponse{}, nil
}
