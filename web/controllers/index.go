package controllers

import (
	"github.com/kataras/iris"
	"imooc_go_lottery/models"
	services "imooc_go_lottery/services"
)

type IndexController struct {
	Ctx iris.Context
	GiftService services.GiftService
}

func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "welcome to Go抽奖系统，<a href='/public/prize.html'>开始抽奖</a>"
}

func (c *IndexController) GetGifts() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	datalist := c.GiftService.GetAll()
	var list []models.LtGift
	for _, data := range datalist {
		if data.SysStatus == 0 {
			list = append(list, data)
		}
	}
	rs["gifts"] = list
	return rs
}

func (c *IndexController) GetNewprize() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	// TODO
	return rs
}


