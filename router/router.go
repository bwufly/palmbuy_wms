package router

import (
    "wms/app/api/hello"
    "github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"wms/app/api/order"
	"wms/app/api/product"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", hello.Hello)
	})

	s.Group("/product", func(group *ghttp.RouterGroup) {
		group.ALL("/transfer", product.Transfer)
	})

	s.Group("/order", func(group *ghttp.RouterGroup) {
		group.GET("/upload", order.UploadShow)
		group.POST("/upload", order.Upload)
	})
}
