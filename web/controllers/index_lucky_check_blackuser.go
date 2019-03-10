package controllers

import (
	"imooc_go_lottery/models"
	"time"
)

// 返回 true 说明是正常用户
func (c *IndexController) checkBlackUser(uid int) (bool, *models.LtUser) {
	info := c.UserService.Get(uid)
	if info != nil && info.Blacktime > int(time.Now().Unix()) {
		return false, info
	}
	return true, info
}