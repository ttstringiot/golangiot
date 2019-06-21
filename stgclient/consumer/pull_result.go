package consumer

import "github.com/ttstringiot/golangiot/stgcommon/message"

type PullResult struct {
	PullStatus      PullStatus
	NextBeginOffset int64
	MinOffset       int64
	MaxOffset       int64
	MsgFoundList    []*message.MessageExt
}
