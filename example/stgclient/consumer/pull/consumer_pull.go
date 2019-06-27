package main

import (
	"fmt"
	"time"

	"github.com/ttstringiot/golangiot/stgclient/process"
)

func main() {
	defaultMQPullConsumer := process.NewDefaultMQPullConsumer("producerGroupId-200")
	defaultMQPullConsumer.SetNamesrvAddr("127.0.0.1:9876")
	defaultMQPullConsumer.Start()

	mqs := defaultMQPullConsumer.FetchSubscribeMessageQueues("cloudzone123")
	for _, mq := range mqs {
		pullResult, err := defaultMQPullConsumer.Pull(mq, "tagA", 0, 32)
		if pullResult == nil || err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(pullResult)
			for _, msgExt := range pullResult.MsgFoundList {

				fmt.Println(string(msgExt.Body))
			}
		}
		time.Sleep(time.Second * 600)
	}

	defaultMQPullConsumer.Shutdown()
	select {}
}
