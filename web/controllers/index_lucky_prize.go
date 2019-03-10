package controllers

import (
	"imooc_go_lottery/conf"
	"imooc_go_lottery/models"
)

func (c *IndexController) prize(prizeCode int, limitBlack bool) *models.ObjGiftPrize {
	var prizeGift *models.ObjGiftPrize
	giftlist := c.GiftService.GetAllUse()
	for _, gift := range giftlist {
		if gift.PrizeCodeA <= prizeCode &&
				gift.PrizeCodeB >= prizeCode {
				if !limitBlack || gift.Gtype < conf.GtypeGiftSmall {
					prizeGift = &gift
					break
				}
		}
	}
	return prizeGift
}
