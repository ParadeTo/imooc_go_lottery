package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"imooc_go_lottery/services"
)

type AdminController struct {
	Ctx iris.Context
	GiftService services.GiftService
}

func (c *AdminController) Get() mvc.Result {
	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title": "管理后台",
		},
	}
}
