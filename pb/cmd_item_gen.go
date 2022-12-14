// Code generated by hk4e-proto-gen. DO NOT EDIT.

package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(PlayerStoreNotify) },
		func() ProtoMessage { return new(StoreWeightLimitNotify) },
		func() ProtoMessage { return new(StoreItemChangeNotify) },
		func() ProtoMessage { return new(StoreItemDelNotify) },
		func() ProtoMessage { return new(ItemAddHintNotify) },
		func() ProtoMessage { return new(UseItemReq) },
		func() ProtoMessage { return new(UseItemRsp) },
		func() ProtoMessage { return new(DropItemReq) },
		func() ProtoMessage { return new(DropItemRsp) },
		func() ProtoMessage { return new(WearEquipReq) },
		func() ProtoMessage { return new(WearEquipRsp) },
		func() ProtoMessage { return new(TakeoffEquipReq) },
		func() ProtoMessage { return new(TakeoffEquipRsp) },
		func() ProtoMessage { return new(AvatarEquipChangeNotify) },
		func() ProtoMessage { return new(WeaponUpgradeReq) },
		func() ProtoMessage { return new(WeaponUpgradeRsp) },
		func() ProtoMessage { return new(WeaponPromoteReq) },
		func() ProtoMessage { return new(WeaponPromoteRsp) },
		func() ProtoMessage { return new(ReliquaryUpgradeReq) },
		func() ProtoMessage { return new(ReliquaryUpgradeRsp) },
		func() ProtoMessage { return new(ReliquaryPromoteReq) },
		func() ProtoMessage { return new(ReliquaryPromoteRsp) },
		func() ProtoMessage { return new(AvatarCardChangeReq) },
		func() ProtoMessage { return new(AvatarCardChangeRsp) },
		func() ProtoMessage { return new(GrantRewardNotify) },
		func() ProtoMessage { return new(WeaponAwakenReq) },
		func() ProtoMessage { return new(WeaponAwakenRsp) },
		func() ProtoMessage { return new(ItemCdGroupTimeNotify) },
		func() ProtoMessage { return new(DropHintNotify) },
		func() ProtoMessage { return new(CombineReq) },
		func() ProtoMessage { return new(CombineRsp) },
		func() ProtoMessage { return new(ForgeQueueDataNotify) },
		func() ProtoMessage { return new(ForgeGetQueueDataReq) },
		func() ProtoMessage { return new(ForgeGetQueueDataRsp) },
		func() ProtoMessage { return new(ForgeStartReq) },
		func() ProtoMessage { return new(ForgeStartRsp) },
		func() ProtoMessage { return new(ForgeQueueManipulateReq) },
		func() ProtoMessage { return new(ForgeQueueManipulateRsp) },
		func() ProtoMessage { return new(ResinChangeNotify) },
		func() ProtoMessage { return new(BuyResinReq) },
		func() ProtoMessage { return new(BuyResinRsp) },
		func() ProtoMessage { return new(MaterialDeleteReturnNotify) },
		func() ProtoMessage { return new(TakeMaterialDeleteReturnReq) },
		func() ProtoMessage { return new(TakeMaterialDeleteReturnRsp) },
		func() ProtoMessage { return new(MaterialDeleteUpdateNotify) },
		func() ProtoMessage { return new(McoinExchangeHcoinReq) },
		func() ProtoMessage { return new(McoinExchangeHcoinRsp) },
		func() ProtoMessage { return new(DestroyMaterialReq) },
		func() ProtoMessage { return new(DestroyMaterialRsp) },
		func() ProtoMessage { return new(SetEquipLockStateReq) },
		func() ProtoMessage { return new(SetEquipLockStateRsp) },
		func() ProtoMessage { return new(CalcWeaponUpgradeReturnItemsReq) },
		func() ProtoMessage { return new(CalcWeaponUpgradeReturnItemsRsp) },
		func() ProtoMessage { return new(ForgeDataNotify) },
		func() ProtoMessage { return new(ForgeFormulaDataNotify) },
		func() ProtoMessage { return new(CombineDataNotify) },
		func() ProtoMessage { return new(CombineFormulaDataNotify) },
		func() ProtoMessage { return new(ClosedItemNotify) },
		func() ProtoMessage { return new(CheckAddItemExceedLimitNotify) },
		func() ProtoMessage { return new(SetIsAutoUnlockSpecificEquipReq) },
		func() ProtoMessage { return new(SetIsAutoUnlockSpecificEquipRsp) },
		func() ProtoMessage { return new(ReliquaryDecomposeReq) },
		func() ProtoMessage { return new(ReliquaryDecomposeRsp) },
		func() ProtoMessage { return new(ReliquaryFilterStateSaveNotify) },
		func() ProtoMessage { return new(ReliquaryFilterStateNotify) },
	)
}

const (
	ProtoCommandPlayerStoreNotify               ProtoCommand = 672
	ProtoCommandStoreWeightLimitNotify          ProtoCommand = 698
	ProtoCommandStoreItemChangeNotify           ProtoCommand = 612
	ProtoCommandStoreItemDelNotify              ProtoCommand = 635
	ProtoCommandItemAddHintNotify               ProtoCommand = 607
	ProtoCommandUseItemReq                      ProtoCommand = 690
	ProtoCommandUseItemRsp                      ProtoCommand = 673
	ProtoCommandDropItemReq                     ProtoCommand = 699
	ProtoCommandDropItemRsp                     ProtoCommand = 631
	ProtoCommandWearEquipReq                    ProtoCommand = 697
	ProtoCommandWearEquipRsp                    ProtoCommand = 681
	ProtoCommandTakeoffEquipReq                 ProtoCommand = 605
	ProtoCommandTakeoffEquipRsp                 ProtoCommand = 682
	ProtoCommandAvatarEquipChangeNotify         ProtoCommand = 647
	ProtoCommandWeaponUpgradeReq                ProtoCommand = 639
	ProtoCommandWeaponUpgradeRsp                ProtoCommand = 653
	ProtoCommandWeaponPromoteReq                ProtoCommand = 622
	ProtoCommandWeaponPromoteRsp                ProtoCommand = 665
	ProtoCommandReliquaryUpgradeReq             ProtoCommand = 604
	ProtoCommandReliquaryUpgradeRsp             ProtoCommand = 693
	ProtoCommandReliquaryPromoteReq             ProtoCommand = 627
	ProtoCommandReliquaryPromoteRsp             ProtoCommand = 694
	ProtoCommandAvatarCardChangeReq             ProtoCommand = 688
	ProtoCommandAvatarCardChangeRsp             ProtoCommand = 626
	ProtoCommandGrantRewardNotify               ProtoCommand = 663
	ProtoCommandWeaponAwakenReq                 ProtoCommand = 695
	ProtoCommandWeaponAwakenRsp                 ProtoCommand = 606
	ProtoCommandItemCdGroupTimeNotify           ProtoCommand = 634
	ProtoCommandDropHintNotify                  ProtoCommand = 650
	ProtoCommandCombineReq                      ProtoCommand = 643
	ProtoCommandCombineRsp                      ProtoCommand = 674
	ProtoCommandForgeQueueDataNotify            ProtoCommand = 676
	ProtoCommandForgeGetQueueDataReq            ProtoCommand = 646
	ProtoCommandForgeGetQueueDataRsp            ProtoCommand = 641
	ProtoCommandForgeStartReq                   ProtoCommand = 649
	ProtoCommandForgeStartRsp                   ProtoCommand = 691
	ProtoCommandForgeQueueManipulateReq         ProtoCommand = 624
	ProtoCommandForgeQueueManipulateRsp         ProtoCommand = 656
	ProtoCommandResinChangeNotify               ProtoCommand = 642
	ProtoCommandBuyResinReq                     ProtoCommand = 602
	ProtoCommandBuyResinRsp                     ProtoCommand = 619
	ProtoCommandMaterialDeleteReturnNotify      ProtoCommand = 661
	ProtoCommandTakeMaterialDeleteReturnReq     ProtoCommand = 629
	ProtoCommandTakeMaterialDeleteReturnRsp     ProtoCommand = 657
	ProtoCommandMaterialDeleteUpdateNotify      ProtoCommand = 700
	ProtoCommandMcoinExchangeHcoinReq           ProtoCommand = 616
	ProtoCommandMcoinExchangeHcoinRsp           ProtoCommand = 687
	ProtoCommandDestroyMaterialReq              ProtoCommand = 640
	ProtoCommandDestroyMaterialRsp              ProtoCommand = 618
	ProtoCommandSetEquipLockStateReq            ProtoCommand = 666
	ProtoCommandSetEquipLockStateRsp            ProtoCommand = 668
	ProtoCommandCalcWeaponUpgradeReturnItemsReq ProtoCommand = 633
	ProtoCommandCalcWeaponUpgradeReturnItemsRsp ProtoCommand = 684
	ProtoCommandForgeDataNotify                 ProtoCommand = 680
	ProtoCommandForgeFormulaDataNotify          ProtoCommand = 689
	ProtoCommandCombineDataNotify               ProtoCommand = 659
	ProtoCommandCombineFormulaDataNotify        ProtoCommand = 632
	ProtoCommandClosedItemNotify                ProtoCommand = 614
	ProtoCommandCheckAddItemExceedLimitNotify   ProtoCommand = 692
	ProtoCommandSetIsAutoUnlockSpecificEquipReq ProtoCommand = 620
	ProtoCommandSetIsAutoUnlockSpecificEquipRsp ProtoCommand = 664
	ProtoCommandReliquaryDecomposeReq           ProtoCommand = 638
	ProtoCommandReliquaryDecomposeRsp           ProtoCommand = 611
	ProtoCommandReliquaryFilterStateSaveNotify  ProtoCommand = 644
	ProtoCommandReliquaryFilterStateNotify      ProtoCommand = 686
)

func (*PlayerStoreNotify) ProtoCommand() ProtoCommand         { return ProtoCommandPlayerStoreNotify }
func (*PlayerStoreNotify) ProtoMessageType() ProtoMessageType { return "PlayerStoreNotify" }

func (*StoreWeightLimitNotify) ProtoCommand() ProtoCommand         { return ProtoCommandStoreWeightLimitNotify }
func (*StoreWeightLimitNotify) ProtoMessageType() ProtoMessageType { return "StoreWeightLimitNotify" }

func (*StoreItemChangeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandStoreItemChangeNotify }
func (*StoreItemChangeNotify) ProtoMessageType() ProtoMessageType { return "StoreItemChangeNotify" }

func (*StoreItemDelNotify) ProtoCommand() ProtoCommand         { return ProtoCommandStoreItemDelNotify }
func (*StoreItemDelNotify) ProtoMessageType() ProtoMessageType { return "StoreItemDelNotify" }

func (*ItemAddHintNotify) ProtoCommand() ProtoCommand         { return ProtoCommandItemAddHintNotify }
func (*ItemAddHintNotify) ProtoMessageType() ProtoMessageType { return "ItemAddHintNotify" }

func (*UseItemReq) ProtoCommand() ProtoCommand         { return ProtoCommandUseItemReq }
func (*UseItemReq) ProtoMessageType() ProtoMessageType { return "UseItemReq" }

func (*UseItemRsp) ProtoCommand() ProtoCommand         { return ProtoCommandUseItemRsp }
func (*UseItemRsp) ProtoMessageType() ProtoMessageType { return "UseItemRsp" }

func (*DropItemReq) ProtoCommand() ProtoCommand         { return ProtoCommandDropItemReq }
func (*DropItemReq) ProtoMessageType() ProtoMessageType { return "DropItemReq" }

func (*DropItemRsp) ProtoCommand() ProtoCommand         { return ProtoCommandDropItemRsp }
func (*DropItemRsp) ProtoMessageType() ProtoMessageType { return "DropItemRsp" }

func (*WearEquipReq) ProtoCommand() ProtoCommand         { return ProtoCommandWearEquipReq }
func (*WearEquipReq) ProtoMessageType() ProtoMessageType { return "WearEquipReq" }

func (*WearEquipRsp) ProtoCommand() ProtoCommand         { return ProtoCommandWearEquipRsp }
func (*WearEquipRsp) ProtoMessageType() ProtoMessageType { return "WearEquipRsp" }

func (*TakeoffEquipReq) ProtoCommand() ProtoCommand         { return ProtoCommandTakeoffEquipReq }
func (*TakeoffEquipReq) ProtoMessageType() ProtoMessageType { return "TakeoffEquipReq" }

func (*TakeoffEquipRsp) ProtoCommand() ProtoCommand         { return ProtoCommandTakeoffEquipRsp }
func (*TakeoffEquipRsp) ProtoMessageType() ProtoMessageType { return "TakeoffEquipRsp" }

func (*AvatarEquipChangeNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAvatarEquipChangeNotify
}
func (*AvatarEquipChangeNotify) ProtoMessageType() ProtoMessageType { return "AvatarEquipChangeNotify" }

func (*WeaponUpgradeReq) ProtoCommand() ProtoCommand         { return ProtoCommandWeaponUpgradeReq }
func (*WeaponUpgradeReq) ProtoMessageType() ProtoMessageType { return "WeaponUpgradeReq" }

func (*WeaponUpgradeRsp) ProtoCommand() ProtoCommand         { return ProtoCommandWeaponUpgradeRsp }
func (*WeaponUpgradeRsp) ProtoMessageType() ProtoMessageType { return "WeaponUpgradeRsp" }

func (*WeaponPromoteReq) ProtoCommand() ProtoCommand         { return ProtoCommandWeaponPromoteReq }
func (*WeaponPromoteReq) ProtoMessageType() ProtoMessageType { return "WeaponPromoteReq" }

func (*WeaponPromoteRsp) ProtoCommand() ProtoCommand         { return ProtoCommandWeaponPromoteRsp }
func (*WeaponPromoteRsp) ProtoMessageType() ProtoMessageType { return "WeaponPromoteRsp" }

func (*ReliquaryUpgradeReq) ProtoCommand() ProtoCommand         { return ProtoCommandReliquaryUpgradeReq }
func (*ReliquaryUpgradeReq) ProtoMessageType() ProtoMessageType { return "ReliquaryUpgradeReq" }

func (*ReliquaryUpgradeRsp) ProtoCommand() ProtoCommand         { return ProtoCommandReliquaryUpgradeRsp }
func (*ReliquaryUpgradeRsp) ProtoMessageType() ProtoMessageType { return "ReliquaryUpgradeRsp" }

func (*ReliquaryPromoteReq) ProtoCommand() ProtoCommand         { return ProtoCommandReliquaryPromoteReq }
func (*ReliquaryPromoteReq) ProtoMessageType() ProtoMessageType { return "ReliquaryPromoteReq" }

func (*ReliquaryPromoteRsp) ProtoCommand() ProtoCommand         { return ProtoCommandReliquaryPromoteRsp }
func (*ReliquaryPromoteRsp) ProtoMessageType() ProtoMessageType { return "ReliquaryPromoteRsp" }

func (*AvatarCardChangeReq) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarCardChangeReq }
func (*AvatarCardChangeReq) ProtoMessageType() ProtoMessageType { return "AvatarCardChangeReq" }

func (*AvatarCardChangeRsp) ProtoCommand() ProtoCommand         { return ProtoCommandAvatarCardChangeRsp }
func (*AvatarCardChangeRsp) ProtoMessageType() ProtoMessageType { return "AvatarCardChangeRsp" }

func (*GrantRewardNotify) ProtoCommand() ProtoCommand         { return ProtoCommandGrantRewardNotify }
func (*GrantRewardNotify) ProtoMessageType() ProtoMessageType { return "GrantRewardNotify" }

func (*WeaponAwakenReq) ProtoCommand() ProtoCommand         { return ProtoCommandWeaponAwakenReq }
func (*WeaponAwakenReq) ProtoMessageType() ProtoMessageType { return "WeaponAwakenReq" }

func (*WeaponAwakenRsp) ProtoCommand() ProtoCommand         { return ProtoCommandWeaponAwakenRsp }
func (*WeaponAwakenRsp) ProtoMessageType() ProtoMessageType { return "WeaponAwakenRsp" }

func (*ItemCdGroupTimeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandItemCdGroupTimeNotify }
func (*ItemCdGroupTimeNotify) ProtoMessageType() ProtoMessageType { return "ItemCdGroupTimeNotify" }

func (*DropHintNotify) ProtoCommand() ProtoCommand         { return ProtoCommandDropHintNotify }
func (*DropHintNotify) ProtoMessageType() ProtoMessageType { return "DropHintNotify" }

func (*CombineReq) ProtoCommand() ProtoCommand         { return ProtoCommandCombineReq }
func (*CombineReq) ProtoMessageType() ProtoMessageType { return "CombineReq" }

func (*CombineRsp) ProtoCommand() ProtoCommand         { return ProtoCommandCombineRsp }
func (*CombineRsp) ProtoMessageType() ProtoMessageType { return "CombineRsp" }

func (*ForgeQueueDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandForgeQueueDataNotify }
func (*ForgeQueueDataNotify) ProtoMessageType() ProtoMessageType { return "ForgeQueueDataNotify" }

func (*ForgeGetQueueDataReq) ProtoCommand() ProtoCommand         { return ProtoCommandForgeGetQueueDataReq }
func (*ForgeGetQueueDataReq) ProtoMessageType() ProtoMessageType { return "ForgeGetQueueDataReq" }

func (*ForgeGetQueueDataRsp) ProtoCommand() ProtoCommand         { return ProtoCommandForgeGetQueueDataRsp }
func (*ForgeGetQueueDataRsp) ProtoMessageType() ProtoMessageType { return "ForgeGetQueueDataRsp" }

func (*ForgeStartReq) ProtoCommand() ProtoCommand         { return ProtoCommandForgeStartReq }
func (*ForgeStartReq) ProtoMessageType() ProtoMessageType { return "ForgeStartReq" }

func (*ForgeStartRsp) ProtoCommand() ProtoCommand         { return ProtoCommandForgeStartRsp }
func (*ForgeStartRsp) ProtoMessageType() ProtoMessageType { return "ForgeStartRsp" }

func (*ForgeQueueManipulateReq) ProtoCommand() ProtoCommand {
	return ProtoCommandForgeQueueManipulateReq
}
func (*ForgeQueueManipulateReq) ProtoMessageType() ProtoMessageType { return "ForgeQueueManipulateReq" }

func (*ForgeQueueManipulateRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandForgeQueueManipulateRsp
}
func (*ForgeQueueManipulateRsp) ProtoMessageType() ProtoMessageType { return "ForgeQueueManipulateRsp" }

func (*ResinChangeNotify) ProtoCommand() ProtoCommand         { return ProtoCommandResinChangeNotify }
func (*ResinChangeNotify) ProtoMessageType() ProtoMessageType { return "ResinChangeNotify" }

func (*BuyResinReq) ProtoCommand() ProtoCommand         { return ProtoCommandBuyResinReq }
func (*BuyResinReq) ProtoMessageType() ProtoMessageType { return "BuyResinReq" }

func (*BuyResinRsp) ProtoCommand() ProtoCommand         { return ProtoCommandBuyResinRsp }
func (*BuyResinRsp) ProtoMessageType() ProtoMessageType { return "BuyResinRsp" }

func (*MaterialDeleteReturnNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMaterialDeleteReturnNotify
}
func (*MaterialDeleteReturnNotify) ProtoMessageType() ProtoMessageType {
	return "MaterialDeleteReturnNotify"
}

func (*TakeMaterialDeleteReturnReq) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeMaterialDeleteReturnReq
}
func (*TakeMaterialDeleteReturnReq) ProtoMessageType() ProtoMessageType {
	return "TakeMaterialDeleteReturnReq"
}

func (*TakeMaterialDeleteReturnRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandTakeMaterialDeleteReturnRsp
}
func (*TakeMaterialDeleteReturnRsp) ProtoMessageType() ProtoMessageType {
	return "TakeMaterialDeleteReturnRsp"
}

func (*MaterialDeleteUpdateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandMaterialDeleteUpdateNotify
}
func (*MaterialDeleteUpdateNotify) ProtoMessageType() ProtoMessageType {
	return "MaterialDeleteUpdateNotify"
}

func (*McoinExchangeHcoinReq) ProtoCommand() ProtoCommand         { return ProtoCommandMcoinExchangeHcoinReq }
func (*McoinExchangeHcoinReq) ProtoMessageType() ProtoMessageType { return "McoinExchangeHcoinReq" }

func (*McoinExchangeHcoinRsp) ProtoCommand() ProtoCommand         { return ProtoCommandMcoinExchangeHcoinRsp }
func (*McoinExchangeHcoinRsp) ProtoMessageType() ProtoMessageType { return "McoinExchangeHcoinRsp" }

func (*DestroyMaterialReq) ProtoCommand() ProtoCommand         { return ProtoCommandDestroyMaterialReq }
func (*DestroyMaterialReq) ProtoMessageType() ProtoMessageType { return "DestroyMaterialReq" }

func (*DestroyMaterialRsp) ProtoCommand() ProtoCommand         { return ProtoCommandDestroyMaterialRsp }
func (*DestroyMaterialRsp) ProtoMessageType() ProtoMessageType { return "DestroyMaterialRsp" }

func (*SetEquipLockStateReq) ProtoCommand() ProtoCommand         { return ProtoCommandSetEquipLockStateReq }
func (*SetEquipLockStateReq) ProtoMessageType() ProtoMessageType { return "SetEquipLockStateReq" }

func (*SetEquipLockStateRsp) ProtoCommand() ProtoCommand         { return ProtoCommandSetEquipLockStateRsp }
func (*SetEquipLockStateRsp) ProtoMessageType() ProtoMessageType { return "SetEquipLockStateRsp" }

func (*CalcWeaponUpgradeReturnItemsReq) ProtoCommand() ProtoCommand {
	return ProtoCommandCalcWeaponUpgradeReturnItemsReq
}
func (*CalcWeaponUpgradeReturnItemsReq) ProtoMessageType() ProtoMessageType {
	return "CalcWeaponUpgradeReturnItemsReq"
}

func (*CalcWeaponUpgradeReturnItemsRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandCalcWeaponUpgradeReturnItemsRsp
}
func (*CalcWeaponUpgradeReturnItemsRsp) ProtoMessageType() ProtoMessageType {
	return "CalcWeaponUpgradeReturnItemsRsp"
}

func (*ForgeDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandForgeDataNotify }
func (*ForgeDataNotify) ProtoMessageType() ProtoMessageType { return "ForgeDataNotify" }

func (*ForgeFormulaDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandForgeFormulaDataNotify }
func (*ForgeFormulaDataNotify) ProtoMessageType() ProtoMessageType { return "ForgeFormulaDataNotify" }

func (*CombineDataNotify) ProtoCommand() ProtoCommand         { return ProtoCommandCombineDataNotify }
func (*CombineDataNotify) ProtoMessageType() ProtoMessageType { return "CombineDataNotify" }

func (*CombineFormulaDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandCombineFormulaDataNotify
}
func (*CombineFormulaDataNotify) ProtoMessageType() ProtoMessageType {
	return "CombineFormulaDataNotify"
}

func (*ClosedItemNotify) ProtoCommand() ProtoCommand         { return ProtoCommandClosedItemNotify }
func (*ClosedItemNotify) ProtoMessageType() ProtoMessageType { return "ClosedItemNotify" }

func (*CheckAddItemExceedLimitNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandCheckAddItemExceedLimitNotify
}
func (*CheckAddItemExceedLimitNotify) ProtoMessageType() ProtoMessageType {
	return "CheckAddItemExceedLimitNotify"
}

func (*SetIsAutoUnlockSpecificEquipReq) ProtoCommand() ProtoCommand {
	return ProtoCommandSetIsAutoUnlockSpecificEquipReq
}
func (*SetIsAutoUnlockSpecificEquipReq) ProtoMessageType() ProtoMessageType {
	return "SetIsAutoUnlockSpecificEquipReq"
}

func (*SetIsAutoUnlockSpecificEquipRsp) ProtoCommand() ProtoCommand {
	return ProtoCommandSetIsAutoUnlockSpecificEquipRsp
}
func (*SetIsAutoUnlockSpecificEquipRsp) ProtoMessageType() ProtoMessageType {
	return "SetIsAutoUnlockSpecificEquipRsp"
}

func (*ReliquaryDecomposeReq) ProtoCommand() ProtoCommand         { return ProtoCommandReliquaryDecomposeReq }
func (*ReliquaryDecomposeReq) ProtoMessageType() ProtoMessageType { return "ReliquaryDecomposeReq" }

func (*ReliquaryDecomposeRsp) ProtoCommand() ProtoCommand         { return ProtoCommandReliquaryDecomposeRsp }
func (*ReliquaryDecomposeRsp) ProtoMessageType() ProtoMessageType { return "ReliquaryDecomposeRsp" }

func (*ReliquaryFilterStateSaveNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandReliquaryFilterStateSaveNotify
}
func (*ReliquaryFilterStateSaveNotify) ProtoMessageType() ProtoMessageType {
	return "ReliquaryFilterStateSaveNotify"
}

func (*ReliquaryFilterStateNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandReliquaryFilterStateNotify
}
func (*ReliquaryFilterStateNotify) ProtoMessageType() ProtoMessageType {
	return "ReliquaryFilterStateNotify"
}
