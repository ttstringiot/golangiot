package body

import (
	"github.com/ttstringiot/golangiot/stgcommon"
	"github.com/ttstringiot/golangiot/stgcommon/sync"
	"github.com/ttstringiot/golangiot/stgnet/protocol"
)

// SubscriptionGroupWrapper 订阅组配置，序列化包装
// Author gaoyanlei
// Since 2017/8/22
type SubscriptionGroupWrapper struct {
	SubscriptionGroupTable *sync.Map             `json:"subscriptionGroupTable"`
	DataVersion            stgcommon.DataVersion `json:"dataVersion"`
	*protocol.RemotingSerializable
}

func NewSubscriptionGroupWrapper() *SubscriptionGroupWrapper {
	wrapper := new(SubscriptionGroupWrapper)
	wrapper.SubscriptionGroupTable = sync.NewMap()
	wrapper.DataVersion = *stgcommon.NewDataVersion()
	wrapper.RemotingSerializable = new(protocol.RemotingSerializable)
	return wrapper
}
