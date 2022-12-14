// Code generated by hk4e-proto-gen. DO NOT EDIT.

package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(GetCityReputationInfoReq) },
		func() ProtoMessage { return new(GetCityReputationInfoRsp) },
		func() ProtoMessage { return new(TakeCityReputationLevelRewardReq) },
		func() ProtoMessage { return new(TakeCityReputationLevelRewardRsp) },
		func() ProtoMessage { return new(CityReputationLevelupNotify) },
		func() ProtoMessage { return new(TakeCityReputationParentQuestReq) },
		func() ProtoMessage { return new(TakeCityReputationParentQuestRsp) },
		func() ProtoMessage { return new(AcceptCityReputationRequestReq) },
		func() ProtoMessage { return new(AcceptCityReputationRequestRsp) },
		func() ProtoMessage { return new(CancelCityReputationRequestReq) },
		func() ProtoMessage { return new(CancelCityReputationRequestRsp) },
		func() ProtoMessage { return new(GetCityReputationMapInfoReq) },
		func() ProtoMessage { return new(GetCityReputationMapInfoRsp) },
		func() ProtoMessage { return new(TakeCityReputationExploreRewardReq) },
		func() ProtoMessage { return new(TakeCityReputationExploreRewardRsp) },
		func() ProtoMessage { return new(CityReputationDataNotify) },
	)
}

const (
	ProtoCommandGetCityReputationInfoReq           ProtoCommand = 2872
	ProtoCommandGetCityReputationInfoRsp           ProtoCommand = 2898
	ProtoCommandTakeCityReputationLevelRewardReq   ProtoCommand = 2812
	ProtoCommandTakeCityReputationLevelRewardRsp   ProtoCommand = 2835
	ProtoCommandCityReputationLevelupNotify        ProtoCommand = 2807
	ProtoCommandTakeCityReputationParentQuestReq   ProtoCommand = 2821
	ProtoCommandTakeCityReputationParentQuestRsp   ProtoCommand = 2803
	ProtoCommandAcceptCityReputationRequestReq     ProtoCommand = 2890
	ProtoCommandAcceptCityReputationRequestRsp     ProtoCommand = 2873
	ProtoCommandCancelCityReputationRequestReq     ProtoCommand = 2899
	ProtoCommandCancelCityReputationRequestRsp     ProtoCommand = 2831
	ProtoCommandGetCityReputationMapInfoReq        ProtoCommand = 2875
	ProtoCommandGetCityReputationMapInfoRsp        ProtoCommand = 2848
	ProtoCommandTakeCityReputationExploreRewardReq ProtoCommand = 2897
	ProtoCommandTakeCityReputationExploreRewardRsp ProtoCommand = 2881
	ProtoCommandCityReputationDataNotify           ProtoCommand = 2805
)

func (*GetCityReputationInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetCityReputationInfoReq
}
func (*GetCityReputationInfoReq) ProtoMessageType() ProtoMessageType {
	return "GetCityReputationInfoReq"
}

func (*GetCityReputationInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetCityReputationInfoRsp
}
func (*GetCityReputationInfoRsp) ProtoMessageType() ProtoMessageType {
	return "GetCityReputationInfoRsp"
}

func (*TakeCityReputationLevelRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeCityReputationLevelRewardReq
}
func (*TakeCityReputationLevelRewardReq) ProtoMessageType() ProtoMessageType {
	return "TakeCityReputationLevelRewardReq"
}

func (*TakeCityReputationLevelRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeCityReputationLevelRewardRsp
}
func (*TakeCityReputationLevelRewardRsp) ProtoMessageType() ProtoMessageType {
	return "TakeCityReputationLevelRewardRsp"
}

func (*CityReputationLevelupNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandCityReputationLevelupNotify
}
func (*CityReputationLevelupNotify) ProtoMessageType() ProtoMessageType {
	return "CityReputationLevelupNotify"
}

func (*TakeCityReputationParentQuestReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeCityReputationParentQuestReq
}
func (*TakeCityReputationParentQuestReq) ProtoMessageType() ProtoMessageType {
	return "TakeCityReputationParentQuestReq"
}

func (*TakeCityReputationParentQuestRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeCityReputationParentQuestRsp
}
func (*TakeCityReputationParentQuestRsp) ProtoMessageType() ProtoMessageType {
	return "TakeCityReputationParentQuestRsp"
}

func (*AcceptCityReputationRequestReq) ProtoCommand() ProtoCommand {
	return ProtoCommandAcceptCityReputationRequestReq
}
func (*AcceptCityReputationRequestReq) ProtoMessageType() ProtoMessageType {
	return "AcceptCityReputationRequestReq"
}

func (*AcceptCityReputationRequestRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandAcceptCityReputationRequestRsp
}
func (*AcceptCityReputationRequestRsp) ProtoMessageType() ProtoMessageType {
	return "AcceptCityReputationRequestRsp"
}

func (*CancelCityReputationRequestReq) ProtoCommand() ProtoCommand {
	return ProtoCommandCancelCityReputationRequestReq
}
func (*CancelCityReputationRequestReq) ProtoMessageType() ProtoMessageType {
	return "CancelCityReputationRequestReq"
}

func (*CancelCityReputationRequestRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandCancelCityReputationRequestRsp
}
func (*CancelCityReputationRequestRsp) ProtoMessageType() ProtoMessageType {
	return "CancelCityReputationRequestRsp"
}

func (*GetCityReputationMapInfoReq) ProtoCommand() ProtoCommand {
	return ProtoCommandGetCityReputationMapInfoReq
}
func (*GetCityReputationMapInfoReq) ProtoMessageType() ProtoMessageType {
	return "GetCityReputationMapInfoReq"
}

func (*GetCityReputationMapInfoRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandGetCityReputationMapInfoRsp
}
func (*GetCityReputationMapInfoRsp) ProtoMessageType() ProtoMessageType {
	return "GetCityReputationMapInfoRsp"
}

func (*TakeCityReputationExploreRewardReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeCityReputationExploreRewardReq
}
func (*TakeCityReputationExploreRewardReq) ProtoMessageType() ProtoMessageType {
	return "TakeCityReputationExploreRewardReq"
}

func (*TakeCityReputationExploreRewardRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeCityReputationExploreRewardRsp
}
func (*TakeCityReputationExploreRewardRsp) ProtoMessageType() ProtoMessageType {
	return "TakeCityReputationExploreRewardRsp"
}

func (*CityReputationDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandCityReputationDataNotify
}
func (*CityReputationDataNotify) ProtoMessageType() ProtoMessageType {
	return "CityReputationDataNotify"
}
