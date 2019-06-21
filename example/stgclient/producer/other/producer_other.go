package main

import (
	"github.com/ttstringiot/golangiot/stgclient/process"
	"github.com/ttstringiot/golangiot/stgcommon"
	"time"
)

func TaskOther() {
	t := time.NewTicker(time.Second * 1000)
	for {
		select {
		case <-t.C:
		}

	}
}
func main() {
	defaultMQProducer := process.NewDefaultMQProducer("producer")
	defaultMQProducer.SetNamesrvAddr("127.0.0.1:10911")
	defaultMQProducer.Start()
	defaultMQProducer.CreateTopic(stgcommon.DEFAULT_TOPIC, "TestTopic", 8)
	go TaskOther()
	time.Sleep(time.Second * 600)
	defaultMQProducer.Shutdown()
	select {}
}
