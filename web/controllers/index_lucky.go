package controllers

import (
	"fmt"
	"imooc_go_lottery/comm"
	"imooc_go_lottery/conf"
	"imooc_go_lottery/models"
	"imooc_go_lottery/web/utils"
	"log"
)

func (c *IndexController) GetLucky() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""

	// 验证登录
	loginuser := comm.GetLoginUser(c.Ctx.Request())
	if loginuser == nil || loginuser.Uid < 1 {
		rs["code"] = 101
		rs["msg"] = "请先登录"
		return rs
	}

	// 分布式锁
	ok := utils.LockLucky(loginuser.Uid)
	if ok {
		defer utils.UnlockLucky(loginuser.Uid)
	} else {
		rs["code"] = 102
		rs["msg"] = "正在抽奖，请稍后重试"
		return rs
	}

	// 验证用户今日参与次数
	ok = c.checkUserday(loginuser.Uid)
	if !ok {
		rs["code"] = 103
		rs["msg"] = "今日抽奖次数已用完"
		return rs
	}

	// 验证ip今日参与次数
	ip := comm.ClientIp(c.Ctx.Request())
	ipDayNum := utils.IncrIpLuckyNum(ip)
	if ipDayNum > conf.IpLimitMax {
		rs["code"] = 104
		rs["msg"] = "抽奖次数过多"
		return rs
	}

	limitBlack := false
	if ipDayNum > conf.IpPrizeMax {
		limitBlack = true
	}

	// 验证ip黑名单
	var blackIpInfo *models.LtBlackip
	if !limitBlack {
		ok, blackIpInfo = c.checkBlackIp(ip)
		if !ok {
			fmt.Println("黑名单中的ip", ip, limitBlack)
			limitBlack = true
		}
	}

	// 验证用户黑名单
	var userInfo *models.LtUser
	if !limitBlack {
		ok, userInfo = c.checkBlackUser(loginuser.Uid)
		if !ok {
			fmt.Println("黑名单中的用户", loginuser.Uid, limitBlack)
			limitBlack = true
		}
	}

	// 获得抽奖编码
	prizeCode := comm.Random(10000)

	// 匹配奖品是否中奖
	prizeGift := c.prize(prizeCode, limitBlack)
	if prizeGift == nil ||
		prizeGift.PrizeNum < 0 ||
		(prizeGift.PrizeNum > 0 && prizeGift.LeftNum <= 0) {
		rs["code"] = 205
		rs["msg"] = "sorry, no prize, try again"
		return rs
	}

	// 有限制奖品发放
	if prizeGift.PrizeNum > 0 {
		ok = utils.PrizeGift(prizeGift.Id, prizeGift.LeftNum)
		if !ok {
			rs["code"] = "207"
			rs["msg"] = "sorry, no prize, try again"
			return rs
		}
	}

	// 不同编码的优惠券的发放
	if prizeGift.Gtype == conf.GtypeCodeDiff {
		code := utils.PrizeCodeDiff(prizeGift.Id, c.CodeService)
		if code == "" {
			rs["code"] = 208
			rs["msg"] = "sorry, no prize, try again"
			return rs
		}
		prizeGift.Gdata = code
	}

	// 记录中奖
	result := models.LtResult{
		GiftId:     prizeGift.Id,
		GiftName:   prizeGift.Title,
		GiftType:   prizeGift.Gtype,
		Uid:        loginuser.Uid,
		Username:   loginuser.Username,
		PrizeCode:  prizeCode,
		GiftData:   prizeGift.Gdata,
		SysCreated: comm.NowUnix(),
		SysIp:      ip,
		SysStatus:  0,
	}
	err := c.ResultService.Create(&result)
	if err != nil {
		log.Println("index_lucky.GetLucky ResultService.Create ", result, ", error=", err)
		rs["code"] = 209
		rs["msg"] = "sorry, no prize, try again"
		return rs
	}

	// 获得了实物大奖，黑名单一段时间
	if prizeGift.Gtype == conf.GtypeGiftLarge {
		c.prizeLarge(ip, loginuser, userInfo, blackIpInfo)
	}

	rs["gift"] = prizeGift
	return rs
}
