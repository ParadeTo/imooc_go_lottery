package controllers

import (
	"imooc_go_lottery/comm"
	"imooc_go_lottery/conf"
	"imooc_go_lottery/models"
	"log"
	"time"
)

func (c *IndexController) checkUserday(uid int) bool {
	userdayInfo := c.UserdayService.GetUserToday(uid)
	if userdayInfo != nil && userdayInfo.Uid == uid {
		if userdayInfo.Num > conf.UserPrizeMax {
			return false
		}
		userdayInfo.Num++
		err103 := c.UserdayService.Update(userdayInfo, nil)
		if err103 != nil {
			log.Println("index_lucky_check_userday UserdayService.Update error=", err103)
		}
	} else {
		todayIntDay := comm.GetTodayIntDay()
		userdayInfo = &models.LtUserday{
			Uid:        uid,
			Day:        todayIntDay,
			Num:        1,
			SysCreated: int(time.Now().Unix()),
			SysUpdated: int(time.Now().Unix()),
		}
		err103 := c.UserdayService.Create(userdayInfo)
		if err103 != nil {
			log.Println("index_lucky_check_userday UserdayService.Create error=", err103)
		}
	}
	return true
}
