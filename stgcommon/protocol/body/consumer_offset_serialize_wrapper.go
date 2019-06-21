package body

import (
	"github.com/ttstringiot/golangiot/stgcommon/sync"
	"github.com/ttstringiot/golangiot/stgnet/protocol"
)

// ConsumerOffsetSerializeWrapper Consumer消费进度，序列化包装
// Author gaoyanlei
// Since 2017/8/22
type ConsumerOffsetSerializeWrapper struct {
	OffsetTable *sync.Map `json:"offsetTable"` // key topic@group value:map[int]int64
	*protocol.RemotingSerializable
}

// NewConsumerOffsetSerializeWrapper 初始化
// Author gaoyanlei
// Since 2017/8/22
func NewConsumerOffsetSerializeWrapper() *ConsumerOffsetSerializeWrapper {
	wrapper := new(ConsumerOffsetSerializeWrapper)
	wrapper.OffsetTable = sync.NewMap()
	wrapper.RemotingSerializable = new(protocol.RemotingSerializable)
	return wrapper
}
