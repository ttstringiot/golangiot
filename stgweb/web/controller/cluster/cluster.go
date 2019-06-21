package cluster

import (
	"git.oschina.net/cloudzone/cloudcommon-go/web/resp"
	"github.com/ttstringiot/golangiot/stgcommon/logger"
	"github.com/ttstringiot/golangiot/stgweb/modules/brokerService"
	"github.com/ttstringiot/golangiot/stgweb/modules/clusterService"
	"github.com/kataras/iris/context"
)

// ClusterList 查询集群节点
// Author: tianyuliang
// Since: 2017/11/9
func ClusterList(ctx context.Context) {
	data, err := clusterService.Default().GetCluserNames()
	if err != nil {
		logger.Errorf("%s %s %s", err.Error(), ctx.Method(), ctx.Path())
		ctx.JSON(resp.NewFailedResponse(resp.ResponseCodes.ServerError, err.Error()))
		return
	}

	ctx.JSON(resp.NewSuccessResponse(data))
}

// ClusterGeneral 查询Broker与Cluster集群概览
// Author: tianyuliang
// Since: 2017/11/9
func ClusterGeneral(ctx context.Context) {
	data, err := brokerService.Default().GetBrokerRuntimeInfo()
	if err != nil {
		logger.Errorf("%s %s %s", err.Error(), ctx.Method(), ctx.Path())
		ctx.JSON(resp.NewFailedResponse(resp.ResponseCodes.ServerError, err.Error()))
		return
	}

	ctx.JSON(resp.NewSuccessResponse(data))
}
