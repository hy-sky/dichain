package types

import (
	"testing"

	"dichain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgMintCoins_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgMintCoins
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgMintCoins{
				Admin: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgMintCoins{
				Admin: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
