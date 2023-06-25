package swap

import (
	"math/rand"

	"dichain/testutil/sample"
	swapsimulation "dichain/x/swap/simulation"
	"dichain/x/swap/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = swapsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgMintCoins = "op_weight_msg_mint_coins"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMintCoins int = 100

	opWeightMsgDistributeCoins = "op_weight_msg_distribute_coins"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDistributeCoins int = 100

	opWeightMsgAddLiquidity = "op_weight_msg_add_liquidity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddLiquidity int = 100

	opWeightMsgSwapCoin = "op_weight_msg_swap_coin"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSwapCoin int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	swapGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&swapGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgMintCoins int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMintCoins, &weightMsgMintCoins, nil,
		func(_ *rand.Rand) {
			weightMsgMintCoins = defaultWeightMsgMintCoins
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMintCoins,
		swapsimulation.SimulateMsgMintCoins(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDistributeCoins int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDistributeCoins, &weightMsgDistributeCoins, nil,
		func(_ *rand.Rand) {
			weightMsgDistributeCoins = defaultWeightMsgDistributeCoins
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDistributeCoins,
		swapsimulation.SimulateMsgDistributeCoins(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddLiquidity int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddLiquidity, &weightMsgAddLiquidity, nil,
		func(_ *rand.Rand) {
			weightMsgAddLiquidity = defaultWeightMsgAddLiquidity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddLiquidity,
		swapsimulation.SimulateMsgAddLiquidity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSwapCoin int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSwapCoin, &weightMsgSwapCoin, nil,
		func(_ *rand.Rand) {
			weightMsgSwapCoin = defaultWeightMsgSwapCoin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSwapCoin,
		swapsimulation.SimulateMsgSwapCoin(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgMintCoins,
			defaultWeightMsgMintCoins,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				swapsimulation.SimulateMsgMintCoins(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDistributeCoins,
			defaultWeightMsgDistributeCoins,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				swapsimulation.SimulateMsgDistributeCoins(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgAddLiquidity,
			defaultWeightMsgAddLiquidity,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				swapsimulation.SimulateMsgAddLiquidity(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSwapCoin,
			defaultWeightMsgSwapCoin,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				swapsimulation.SimulateMsgSwapCoin(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
