package process

import (
	"fmt"
	"github.com/ttstringiot/golangiot/stgcommon/logger"
	"github.com/ttstringiot/golangiot/stgcommon/message"
	code "github.com/ttstringiot/golangiot/stgcommon/protocol"
	"github.com/ttstringiot/golangiot/stgcommon/protocol/header"
	"github.com/ttstringiot/golangiot/stgnet/netm"
	"github.com/ttstringiot/golangiot/stgnet/protocol"
)

// 客户端处理器
// Author: yintongqiang
// Since:  2017/8/8
type ClientRemotingProcessor struct {
	MQClientFactory *MQClientInstance
}

func NewClientRemotingProcessor(mqClientFactory *MQClientInstance) *ClientRemotingProcessor {
	return &ClientRemotingProcessor{
		MQClientFactory: mqClientFactory,
	}
}

// 处理request
func (self *ClientRemotingProcessor) ProcessRequest(ctx netm.Context, request *protocol.RemotingCommand) (*protocol.RemotingCommand, error) {
	switch request.Code {
	case code.NOTIFY_CONSUMER_IDS_CHANGED:
		return self.notifyConsumerIdsChanged(ctx, request)
	case code.CONSUME_MESSAGE_DIRECTLY:
		return self.consumeMessageDirectly(ctx, request)
	default:
		return nil, nil
	}
	return nil, nil
}

func (self *ClientRemotingProcessor) notifyConsumerIdsChanged(ctx netm.Context, request *protocol.RemotingCommand) (*protocol.RemotingCommand, error) {
	response := protocol.CreateDefaultResponseCommand()

	requestHeader := &header.NotifyConsumerIdsChangedRequestHeader{}
	err := request.DecodeCommandCustomHeader(requestHeader)
	if err != nil {
		logger.Errorf("err: %s", err.Error())
		return response, err
	}

	format := "receive broker's notification[%s], the consumer group: %s changed, rebalance immediately"
	logger.Infof(format, ctx.RemoteAddr().String(), requestHeader.ConsumerGroup)
	self.MQClientFactory.rebalanceImmediately()

	response.Code = code.SUCCESS
	response.Remark = ""
	return response, nil
}

func (self *ClientRemotingProcessor) consumeMessageDirectly(ctx netm.Context, request *protocol.RemotingCommand) (*protocol.RemotingCommand, error) {
	response := protocol.CreateDefaultResponseCommand()

	requestHeader := &header.ConsumeMessageDirectlyResultRequestHeader{}
	err := request.DecodeCommandCustomHeader(requestHeader)
	if err != nil {
		logger.Errorf("err: %s", err.Error())
		return response, err
	}

	msg, err := message.DecodeMessageExt(request.Body, true, true)
	if err != nil {
		logger.Errorf("err: %s", err.Error())
		return response, err
	}
	result := self.MQClientFactory.ConsumeMessageDirectly(msg, requestHeader.ConsumerGroup, requestHeader.BrokerName)
	if result != nil {
		response.Code = code.SUCCESS
		response.Body = result.CustomEncode(result)
		return response, nil
	}

	response.Remark = fmt.Sprintf("The Consumer Group <%s> not exist in this consumer", requestHeader.ConsumerGroup)
	return response, nil
}
