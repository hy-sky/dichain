package keeper

import (
	"context"
	"errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"dichain/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// x/swap/keeper/msg_server_distribute_coins.go
func (k msgServer) DistributeCoins(goCtx context.Context, msg *types.MsgDistributeCoins) (*types.MsgDistributeCoinsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	admin, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return nil, err
	}

	toAddress, err := sdk.AccAddressFromBech32(msg.ToAddress)
	if err != nil {
		return nil, err
	}

	isAdmin := k.IsAdmin(ctx, admin)
	if !isAdmin {
		return nil, errors.New(msg.Admin + " is not admin")
	}

	err = k.Keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, banktypes.ModuleName, toAddress, msg.Amount)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Admin),
		),
	)

	return &types.MsgDistributeCoinsResponse{}, nil
}
