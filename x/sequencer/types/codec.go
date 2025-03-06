package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

var ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateSequencer{}, "sequencer/CreateSequencer", nil)
	cdc.RegisterConcrete(&MsgUnbond{}, "sequencer/Unbond", nil)
	cdc.RegisterConcrete(&MsgIncreaseBond{}, "sequencer/IncreaseBond", nil)
	cdc.RegisterConcrete(&MsgDecreaseBond{}, "sequencer/DecreaseBond", nil)
	cdc.RegisterConcrete(&MsgUpdateRewardAddress{}, "sequencer/UpdateRewardAddress", nil)
	cdc.RegisterConcrete(&MsgUpdateWhitelistedRelayers{}, "sequencer/UpdateWhitelistedRelayers", nil)
	cdc.RegisterConcrete(&MsgKickProposer{}, "sequencer/KickProposer", nil)
	cdc.RegisterConcrete(&MsgUpdateOptInStatus{}, "sequencer/UpdateOtpInStatus", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "sequencer/UpdateParams", nil)
	cdc.RegisterConcrete(&PunishSequencerProposal{}, "sequencer/PunishSequencerProposal", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSequencer{},
		&MsgUnbond{},
		&MsgIncreaseBond{},
		&MsgDecreaseBond{},
		&MsgUpdateRewardAddress{},
		&MsgUpdateWhitelistedRelayers{},
		&MsgKickProposer{},
		&MsgUpdateOptInStatus{},
		&MsgUpdateParams{},
	)

	registry.RegisterImplementations((*govtypes.Content)(nil), &PunishSequencerProposal{})

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
