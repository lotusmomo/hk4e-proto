// Code generated by hk4e-proto-gen. DO NOT EDIT.

package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(RegionalPlayInfoNotify) },
		func() ProtoMessage { return new(DeathZoneInfoNotify) },
		func() ProtoMessage { return new(PlayerDeathZoneNotify) },
	)
}

const (
	ProtoCommandRegionalPlayInfoNotify ProtoCommand = 6276
	ProtoCommandDeathZoneInfoNotify    ProtoCommand = 6268
	ProtoCommandPlayerDeathZoneNotify  ProtoCommand = 6275
)

func (*RegionalPlayInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandRegionalPlayInfoNotify }
func (*RegionalPlayInfoNotify) ProtoMessageType() ProtoMessageType { return "RegionalPlayInfoNotify" }

func (*DeathZoneInfoNotify) ProtoCommand() ProtoCommand         { return ProtoCommandDeathZoneInfoNotify }
func (*DeathZoneInfoNotify) ProtoMessageType() ProtoMessageType { return "DeathZoneInfoNotify" }

func (*PlayerDeathZoneNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerDeathZoneNotify }
func (*PlayerDeathZoneNotify) ProtoMessageType() ProtoMessageType { return "PlayerDeathZoneNotify" }
