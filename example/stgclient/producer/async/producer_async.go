package main

import (
	"fmt"
	"github.com/ttstringiot/golangiot/stgclient/process"
	"github.com/ttstringiot/golangiot/stgcommon/message"
	"time"
)

func TaskCallBack() {
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
	for i := 0; i < 10; i++ {
		err := defaultMQProducer.SendCallBack(message.NewMessage("TestTopic", "tagA", []byte("send callback msg")),
			func(sendResult *process.SendResult, err error) {
				if err != nil {
					fmt.Println(err)
				}
				if sendResult != nil {
					fmt.Println(sendResult.ToString())
				}
			})
		if err != nil {
			fmt.Println(err)
		}

	}
	go TaskCallBack()
	time.Sleep(time.Second * 600)
	defaultMQProducer.Shutdown()
	select {}
}
