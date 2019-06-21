package consumer

import (
	"github.com/ttstringiot/golangiot/stgcommon/message"
	"github.com/ttstringiot/golangiot/stgclient/consumer/listener"
)
// MessageListenerConcurrently: 普通消息消费接口
// Author: yintongqiang
// Since:  2017/8/10

type MessageListenerConcurrently interface {
	ConsumeMessage(msgs []*message.MessageExt, context *ConsumeConcurrentlyContext) listener.ConsumeConcurrentlyStatus
}
