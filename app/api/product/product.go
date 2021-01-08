package product

import (
	"github.com/gogf/gf/net/ghttp"
	"wms/app/model"
	"wms/app/service"
	"wms/library/response"
)

func Transfer(r *ghttp.Request) {
	var (
		data *model.ProductTransferReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.User.Transfer(data);err != nil {
		response.JsonExit(r, 1, err.Error())
	}

}