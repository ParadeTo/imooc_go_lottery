package controllers

import (
	"imooc_go_lottery/comm"
	"imooc_go_lottery/models"
)

func (c *IndexController) prizeLarge(
	ip string,
	loginuser *models.ObjLoginuser,
	userinfo *models.LtUser, // 用户
	blackipInfo *models.LtBlackip) {
	nowTime := comm.NowUnix()
	blackTime := 30 * 86400
	// 用户黑名单
	if userinfo == nil || userinfo.Id <= 0 {
		userinfo = &models.LtUser{
			Id:         loginuser.Uid,
			Username:   loginuser.Username,
			Blacktime:  nowTime+blackTime,
			SysCreated: nowTime,
			SysIp:      ip,
		}
		c.UserService.Create(userinfo)
	} else {
		userinfo = &models.LtUser{
			Id:        loginuser.Uid,
			Blacktime: nowTime + blackTime,
			SysUpdate: nowTime,
		}
		c.UserService.Update(userinfo, nil)
	}
	// ip 黑名单
	if blackipInfo == nil || blackipInfo.Id <= 0 {
		blackipInfo = &models.LtBlackip{
			Id:         loginuser.Uid,
			Blacktime:  nowTime+blackTime,
			SysCreated: nowTime,
		}
		c.BlackipService.Create(blackipInfo)
	} else {
		blackipInfo.Blacktime = nowTime + blackTime
		blackipInfo.SysUpdated = nowTime
		c.BlackipService.Update(blackipInfo, nil)
	}
}