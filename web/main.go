package main

import (
	"fmt"
	"imooc_go_lottery/bootstrap"
	"imooc_go_lottery/web/routes"
)

var port = 8080

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("抽奖系统", "ayou")
	app.Bootstrap()
	app.Configure(routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(fmt.Sprintf(":%d", port))
}
