package main

import (
	_ "wms/boot"
	_ "wms/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
