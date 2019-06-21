package test

import (
	"github.com/ttstringiot/golangiot/stgbroker"
	"github.com/ttstringiot/golangiot/stgbroker/out"
	"github.com/ttstringiot/golangiot/stgnet/remoting"
	"testing"
)

func TestRegisterBroker(t *testing.T) {
	brokerController := stgbroker.CreateBrokerController()
	brokerController.Initialize()
	brokerController.TopicConfigManager.Load()
	topicConfigWrapper := brokerController.TopicConfigManager.TopicConfigSerializeWrapper
	api := brokerController.BrokerOuterAPI
	api.Start()
	api.RegisterBroker("0.0.0.0:9876", "out", "10.122.2.28:10911", "broker-1", "10.122.1.20:10912", 1, topicConfigWrapper, false, nil)
}

func TestUnRegisterBroker(t *testing.T) {
	brokerController := stgbroker.CreateBrokerController()
	brokerController.Initialize()
	brokerController.TopicConfigManager.Load()
	api := brokerController.BrokerOuterAPI
	api.Start()
	api.UnRegisterBroker("0.0.0.0:9876", "out", "10.122.2.28:10911", "broker-1", 1)
}

func TestFetchNameServerAddr(t *testing.T) {
	api := out.NewBrokerOuterAPI(remoting.NewDefalutRemotingClient())
	api.Start()
	api.UpdateNameServerAddressList("0.0.0.0:9999")
	api.FetchNameServerAddr()
}
