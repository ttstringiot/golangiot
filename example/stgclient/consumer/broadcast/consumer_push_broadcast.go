package main

import (
	"fmt"
	"github.com/ttstringiot/golangiot/stgclient/consumer"
	"github.com/ttstringiot/golangiot/stgclient/consumer/listener"
	"github.com/ttstringiot/golangiot/stgclient/process"
	"github.com/ttstringiot/golangiot/stgcommon/message"
	"github.com/ttstringiot/golangiot/stgcommon/protocol/heartbeat"
	"time"
)

type MessageListenerImpl struct {
}

func (listenerImpl *MessageListenerImpl) ConsumeMessage(msgs []*message.MessageExt, context *consumer.ConsumeConcurrentlyContext) listener.ConsumeConcurrentlyStatus {
	for _, msg := range msgs {
		fmt.Println(msg.ToString())
	}
	return listener.CONSUME_SUCCESS
}

func taskC() {
	t := time.NewTicker(time.Second * 1000)
	for {
		select {
		case <-t.C:
		}

	}
}

func main() {
	defaultMQPushConsumer := process.NewDefaultMQPushConsumer("conasuemer1")
	defaultMQPushConsumer.SetConsumeFromWhere(heartbeat.CONSUME_FROM_LAST_OFFSET)
	defaultMQPushConsumer.SetMessageModel(heartbeat.BROADCASTING)
	defaultMQPushConsumer.SetNamesrvAddr("10.112.68.189:9876")
	defaultMQPushConsumer.Subscribe("cloudzone4", "tagA")
	defaultMQPushConsumer.RegisterMessageListener(&MessageListenerImpl{})
	defaultMQPushConsumer.Start()
	time.Sleep(time.Second * 6000)
	defaultMQPushConsumer.Shutdown()
	go taskC()
	select {}
}
