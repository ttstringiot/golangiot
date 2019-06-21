package main

import (
	"fmt"
	"time"

	"github.com/ttstringiot/golangiot/stgclient/process"
	"github.com/ttstringiot/golangiot/stgcommon/message"
)

func TaskOneWay() {
	t := time.NewTicker(time.Second * 1000)
	for {
		select {
		case <-t.C:
		}

	}
}
func main() {
	defaultMQProducer := process.NewDefaultMQProducer("producer")
	defaultMQProducer.SetNamesrvAddr("127.0.0.1:9876")
	defaultMQProducer.Start()
	for i := 0; i < 640; i++ {
		err := defaultMQProducer.SendOneWay(message.NewMessage("cloudzoneoneway", "tagA", []byte("send oneway msg 呵呵")))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("sendOneWay>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		}
	}
	go TaskOneWay()
	time.Sleep(time.Second * 600)
	defaultMQProducer.Shutdown()
	select {}
}
