package controllers

import (
	"imooc_go_lottery/models"
	"time"
)

// 返回 true 说明是正常用户
func (c *IndexController) checkBlackIp(ip string) (bool, *models.LtBlackip) {
	info := c.BlackipService.GetByIp(ip)
	if info == nil || info.Ip == "" {
		return true, nil
	}

	// ip 黑名单存在，并且还在黑名单限制时间内
	if info.Blacktime > int(time.Now().Unix()) {
		return false, info
	}
	return true, info
}