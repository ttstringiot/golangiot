package route

import (
	"git.oschina.net/cloudzone/cloudcommon-go/web"
	"github.com/ttstringiot/golangiot/stgweb/web/controller/broker"
	"github.com/ttstringiot/golangiot/stgweb/web/controller/cluster"
	"github.com/ttstringiot/golangiot/stgweb/web/controller/connection"
	"github.com/ttstringiot/golangiot/stgweb/web/controller/general"
	"github.com/ttstringiot/golangiot/stgweb/web/controller/group"
	"github.com/ttstringiot/golangiot/stgweb/web/controller/message"
	"github.com/ttstringiot/golangiot/stgweb/web/controller/topic"
	"github.com/kataras/iris/context"
)

func Route(ctx *web.Context) error {
	route := ctx.Route()
	api := route.Party("/api/v1/")
	{
		api.Options("/{root:path}", func(ctx context.Context) {}) // fix options not match bug
	}

	// 整体概览
	{
		api.Get("/general", general.GeneralStats)
	}

	// cluster集群管理
	{
		api.Get("/cluster/general", cluster.ClusterGeneral)
		api.Get("/cluster/list", cluster.ClusterList)
	}

	// topic管理
	{
		api.Post("/topic", topic.CreateTopic)
		api.Put("/topic", topic.UpdateTopic)
		api.Delete("/topic", topic.DeleteTopic)
		api.Get("/topic/list", topic.TopicList)
		api.Get("/topic/stats", topic.TopicStats)
		api.Get("/topic/route", topic.TopicRoute)
	}

	// 消费进度
	{
		api.Get("/group/progress", group.ConsumeProgress)
		api.Get("/group/list", group.GroupList)
	}

	// 消费进程
	{
		api.Get("/connection/online", connection.ConnectionOnline)
		api.Get("/connection/detail", connection.ConnectionDetail)
	}

	// 消息查询、消费轨迹
	{
		api.Get("/msg/body", message.MessageBody)
		api.Get("/msg/track", message.MessageTrack)
		api.Get("/msg/query", message.MessageQuery)
	}

	// 运维
	{
		api.Delete("/consumer/subGroup", broker.DeleteSubGroup)
		api.Put("/consumer/subGroup", broker.UpdateSubGroup)
		api.Post("/broker/syncTopic", broker.SyncTopicToBroker)
		api.Post("/broker/wipePerm", broker.WipeWritePermBroker)
	}

	return nil
}
