package main

import (
	"imooc_go_lottery/dao"
	"imooc_go_lottery/datasource"
)

func main() {
	giftDao := dao.NewGiftDao(datasource.InstanceDbMaster())
	giftDao.DecrLeftNum(1, 1)
}
