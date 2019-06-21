package registry

import (
	"github.com/ttstringiot/golangiot/stgcommon/logger"
	"git.oschina.net/cloudzone/cloudcommon-go/web/resp"
	"github.com/ttstringiot/golangiot/stgweb/modules/clusterService"
	"github.com/kataras/iris/context"
)

// QueryNamesrvAddrs 查询namesrv节点
// Author: tianyuliang
// Since: 2017/11/9
func QueryNamesrvAddrs(ctx context.Context) {
	data, err := clusterService.Default().GetNamesrvNodes()
	if err != nil {
		logger.Errorf("%s %s %s", err.Error(), ctx.Method(), ctx.Path())
		ctx.JSON(resp.NewFailedResponse(resp.ResponseCodes.ServerError, err.Error()))
		return
	}

	ctx.JSON(resp.NewSuccessResponse(data))
}
