package routes

import (
	"github.com/kataras/iris/mvc"
	"imooc_go_lottery/bootstrap"
	"imooc_go_lottery/services"
	"imooc_go_lottery/web/controllers"
)

func Configure(b *bootstrap.Bootstrapper) {
	giftService := services.NewGiftService()
	userdayService := services.NewUserdayService()
	userService := services.NewUserService()
	blackipService := services.NewBlackipService()
	codeService := services.NewCodeService()
	resultService := services.NewResultService()

	index := mvc.New(b.Party("/"))
	index.Register(giftService, userdayService, userService, blackipService, codeService, resultService)
	index.Handle(new(controllers.IndexController))

	admin := mvc.New(b.Party("/admin"))
	admin.Register(giftService)
	admin.Handle(new(controllers.AdminController))

	adminResult := admin.Party("/result")
	adminResult.Register()
}
