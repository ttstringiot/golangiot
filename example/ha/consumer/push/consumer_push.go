package main

import (
	"fmt"
	"github.com/ttstringiot/golangiot/stgclient/consumer"
	"github.com/ttstringiot/golangiot/stgclient/consumer/listener"
	"github.com/ttstringiot/golangiot/stgclient/process"
	"github.com/ttstringiot/golangiot/stgcommon/message"
	"github.com/ttstringiot/golangiot/stgcommon/protocol/heartbeat"
	"github.com/ttstringiot/golangiot/stgcommon/sync"
	"sync/atomic"
	"time"
)

type MessageListenerImpl struct {
	MsgCount   int64
	StartTime  int64
	MapContent *sync.Map
}

func (listenerImpl *MessageListenerImpl) ConsumeMessage(msgs []*message.MessageExt, context *consumer.ConsumeConcurrentlyContext) listener.ConsumeConcurrentlyStatus {
	for _, msg := range msgs {
		count := atomic.AddInt64(&listenerImpl.MsgCount, 1)
		listenerImpl.MapContent.Put(msg.ToString(), 0)
		fmt.Println(count, msg.ToString(), listenerImpl.MapContent.Size())
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
	var (
		consumerGroup = "myConsumerGroup"
		nameServer    = "127.0.0.1:9876"
		topic         = "test"
		tag           = "tagA"
	)

	defaultMQPushConsumer := process.NewDefaultMQPushConsumer(consumerGroup)
	defaultMQPushConsumer.SetConsumeFromWhere(heartbeat.CONSUME_FROM_LAST_OFFSET)
	defaultMQPushConsumer.SetMessageModel(heartbeat.CLUSTERING)
	defaultMQPushConsumer.SetNamesrvAddr(nameServer)
	defaultMQPushConsumer.Subscribe(topic, tag)
	defaultMQPushConsumer.RegisterMessageListener(&MessageListenerImpl{StartTime: time.Now().Unix(), MapContent: sync.NewMap()})
	defaultMQPushConsumer.Start()
	go taskC()
	select {}
}
