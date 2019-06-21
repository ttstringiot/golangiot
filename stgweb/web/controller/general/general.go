package general

import (
	"git.oschina.net/cloudzone/cloudcommon-go/web/resp"
	"github.com/ttstringiot/golangiot/stgcommon/logger"
	"github.com/ttstringiot/golangiot/stgweb/modules/generalService"
	"github.com/kataras/iris/context"
)

// General 查询云平台的概况数据
// Author: tianyuliang
// Since: 2017/11/7
func GeneralStats(ctx context.Context) {
	data, err := generalService.Default().GeneralStats()
	if err != nil {
		logger.Errorf("%s %s %s", err.Error(), ctx.Method(), ctx.Path())
		ctx.JSON(resp.NewFailedResponse(resp.ResponseCodes.ServerError, err.Error()))
		return
	}
	ctx.JSON(resp.NewSuccessResponse(data))
}
