package routes

import (
	"github.com/kataras/iris/mvc"
	"imooc_go_lottery/bootstrap"
	"imooc_go_lottery/services"
	"imooc_go_lottery/web/controllers"
)

func Configure(b *bootstrap.Bootstrapper) {
	giftService := services.NewGiftService()

	index := mvc.New(b.Party("/"))
	index.Register(giftService)
	index.Handle(new(controllers.IndexController))
}
