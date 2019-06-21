package test

import (
	"github.com/ttstringiot/golangiot/stgbroker"
	"github.com/ttstringiot/golangiot/stgcommon/logger"
	"github.com/ttstringiot/golangiot/stgcommon/protocol/body"
	"testing"
)

func TestTopicConfigManagerDecode(t *testing.T) {
	brokerController := stgbroker.CreateBrokerController()
	brokerController.Initialize()
	brokerController.TopicConfigManager.Load()
	data := brokerController.TopicConfigManager.Encode(false)
	buf := []byte(data)

	topicWrapper := body.NewTopicConfigSerializeWrapper()
	topicConfigManager := stgbroker.TopicConfigManager{}
	topicConfigManager.TopicConfigSerializeWrapper = topicWrapper
	topicConfigManager.Decode(buf)
	logger.Infof("TopicConfigManager.Decode() success.\n\t\t %s", topicConfigManager.TopicConfigSerializeWrapper.ToString())

}
