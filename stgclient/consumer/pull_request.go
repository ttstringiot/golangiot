package consumer

import "github.com/ttstringiot/golangiot/stgcommon/message"
// PullRequest: 拉消息request
// Author: yintongqiang
// Since:  2017/8/10

type PullRequest struct {
	ConsumerGroup string
	MessageQueue  *message.MessageQueue
	ProcessQueue  *ProcessQueue
	NextOffset    int64
}