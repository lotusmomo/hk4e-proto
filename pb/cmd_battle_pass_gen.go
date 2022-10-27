// Code generated by hk4e-proto-gen. DO NOT EDIT.

package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(BattlePassAllDataNotify) },
		func() ProtoMessage { return new(BattlePassMissionUpdateNotify) },
		func() ProtoMessage { return new(BattlePassMissionDelNotify) },
		func() ProtoMessage { return new(BattlePassCurScheduleUpdateNotify) },
		func() ProtoMessage { return new(TakeBattlePassRewardReq) },
		func() ProtoMessage { return new(TakeBattlePassRewardRsp) },
		func() ProtoMessage { return new(TakeBattlePassMissionPointReq) },
		func() ProtoMessage { return new(TakeBattlePassMissionPointRsp) },
		func() ProtoMessage { return new(GetBattlePassProductReq) },
		func() ProtoMessage { return new(GetBattlePassProductRsp) },
		func() ProtoMessage { return new(SetBattlePassViewedReq) },
		func() ProtoMessage { return new(SetBattlePassViewedRsp) },
		func() ProtoMessage { return new(BattlePassBuySuccNotify) },
		func() ProtoMessage { return new(BuyBattlePassLevelReq) },
		func() ProtoMessage { return new(BuyBattlePassLevelRsp) },
	)
}

const (
	ProtoCommandBattlePassAllDataNotify           ProtoCommand = 2626
	ProtoCommandBattlePassMissionUpdateNotify     ProtoCommand = 2618
	ProtoCommandBattlePassMissionDelNotify        ProtoCommand = 2625
	ProtoCommandBattlePassCurScheduleUpdateNotify ProtoCommand = 2607
	ProtoCommandTakeBattlePassRewardReq           ProtoCommand = 2602
	ProtoCommandTakeBattlePassRewardRsp           ProtoCommand = 2631
	ProtoCommandTakeBattlePassMissionPointReq     ProtoCommand = 2629
	ProtoCommandTakeBattlePassMissionPointRsp     ProtoCommand = 2622
	ProtoCommandGetBattlePassProductReq           ProtoCommand = 2644
	ProtoCommandGetBattlePassProductRsp           ProtoCommand = 2649
	ProtoCommandSetBattlePassViewedReq            ProtoCommand = 2641
	ProtoCommandSetBattlePassViewedRsp            ProtoCommand = 2642
	ProtoCommandBattlePassBuySuccNotify           ProtoCommand = 2614
	ProtoCommandBuyBattlePassLevelReq             ProtoCommand = 2647
	ProtoCommandBuyBattlePassLevelRsp             ProtoCommand = 2637
)

func (*BattlePassAllDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandBattlePassAllDataNotify
}
func (*BattlePassAllDataNotify) ProtoMessageType() ProtoMessageType { return "BattlePassAllDataNotify" }

func (*BattlePassMissionUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandBattlePassMissionUpdateNotify
}
func (*BattlePassMissionUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "BattlePassMissionUpdateNotify"
}

func (*BattlePassMissionDelNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandBattlePassMissionDelNotify
}
func (*BattlePassMissionDelNotify) ProtoMessageType() ProtoMessageType {
	return "BattlePassMissionDelNotify"
}

func (*BattlePassCurScheduleUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandBattlePassCurScheduleUpdateNotify
}
func (*BattlePassCurScheduleUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "BattlePassCurScheduleUpdateNotify"
}

func (*TakeBattlePassRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeBattlePassRewardReq
}
func (*TakeBattlePassRewardReq) ProtoMessageType() ProtoMessageType { return "TakeBattlePassRewardReq" }

func (*TakeBattlePassRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeBattlePassRewardRsp
}
func (*TakeBattlePassRewardRsp) ProtoMessageType() ProtoMessageType { return "TakeBattlePassRewardRsp" }

func (*TakeBattlePassMissionPointReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeBattlePassMissionPointReq
}
func (*TakeBattlePassMissionPointReq) ProtoMessageType() ProtoMessageType {
	return "TakeBattlePassMissionPointReq"
}

func (*TakeBattlePassMissionPointRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeBattlePassMissionPointRsp
}
func (*TakeBattlePassMissionPointRsp) ProtoMessageType() ProtoMessageType {
	return "TakeBattlePassMissionPointRsp"
}

func (*GetBattlePassProductReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetBattlePassProductReq
}
func (*GetBattlePassProductReq) ProtoMessageType() ProtoMessageType { return "GetBattlePassProductReq" }

func (*GetBattlePassProductRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetBattlePassProductRsp
}
func (*GetBattlePassProductRsp) ProtoMessageType() ProtoMessageType { return "GetBattlePassProductRsp" }

func (*SetBattlePassViewedReq) ProtoCommand() ProtoCommand         { return ProtoCommandSetBattlePassViewedReq }
func (*SetBattlePassViewedReq) ProtoMessageType() ProtoMessageType { return "SetBattlePassViewedReq" }

func (*SetBattlePassViewedRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSetBattlePassViewedRsp }
func (*SetBattlePassViewedRsp) ProtoMessageType() ProtoMessageType { return "SetBattlePassViewedRsp" }

func (*BattlePassBuySuccNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandBattlePassBuySuccNotify
}
func (*BattlePassBuySuccNotify) ProtoMessageType() ProtoMessageType { return "BattlePassBuySuccNotify" }

func (*BuyBattlePassLevelReq) ProtoCommand() ProtoCommand         { return ProtoCommandBuyBattlePassLevelReq }
func (*BuyBattlePassLevelReq) ProtoMessageType() ProtoMessageType { return "BuyBattlePassLevelReq" }

func (*BuyBattlePassLevelRsp) ProtoCommand() ProtoCommand         { return ProtoCommandBuyBattlePassLevelRsp }
func (*BuyBattlePassLevelRsp) ProtoMessageType() ProtoMessageType { return "BuyBattlePassLevelRsp" }