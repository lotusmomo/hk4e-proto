// Code generated by hk4e-proto-gen. DO NOT EDIT.

package pb

func init() {
	ProtoCommandNewFuncMap.Add(
		func() ProtoMessage { return new(AranaraCollectionDataNotify) },
		func() ProtoMessage { return new(AddAranaraCollectionNotify) },
		func() ProtoMessage { return new(CataLogFinishedGlobalWatcherAllDataNotify) },
		func() ProtoMessage { return new(CataLogNewFinishedGlobalWatcherNotify) },
	)
}

const (
	ProtoCommandAranaraCollectionDataNotify               ProtoCommand = 6376
	ProtoCommandAddAranaraCollectionNotify                ProtoCommand = 6368
	ProtoCommandCataLogFinishedGlobalWatcherAllDataNotify ProtoCommand = 6370
	ProtoCommandCataLogNewFinishedGlobalWatcherNotify     ProtoCommand = 6395
)

func (*AranaraCollectionDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAranaraCollectionDataNotify
}
func (*AranaraCollectionDataNotify) ProtoMessageType() ProtoMessageType {
	return "AranaraCollectionDataNotify"
}

func (*AddAranaraCollectionNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandAddAranaraCollectionNotify
}
func (*AddAranaraCollectionNotify) ProtoMessageType() ProtoMessageType {
	return "AddAranaraCollectionNotify"
}

func (*CataLogFinishedGlobalWatcherAllDataNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandCataLogFinishedGlobalWatcherAllDataNotify
}
func (*CataLogFinishedGlobalWatcherAllDataNotify) ProtoMessageType() ProtoMessageType {
	return "CataLogFinishedGlobalWatcherAllDataNotify"
}

func (*CataLogNewFinishedGlobalWatcherNotify) ProtoCommand() ProtoCommand {
	return ProtoCommandCataLogNewFinishedGlobalWatcherNotify
}
func (*CataLogNewFinishedGlobalWatcherNotify) ProtoMessageType() ProtoMessageType {
	return "CataLogNewFinishedGlobalWatcherNotify"
}