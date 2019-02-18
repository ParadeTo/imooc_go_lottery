package models

type LtResult struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	GiftId     int    `xorm:"not null default 0 comment('奖品 id，关联 lt_gift 表') INT(11)"`
	GiftName   string `xorm:"not null default '' comment('奖品名称') VARCHAR(255)"`
	GiftType   int    `xorm:"not null default 0 comment('奖品类型，同 lt_gift.type') INT(11)"`
	Uid        int    `xorm:"not null default 0 comment('用户 id') INT(11)"`
	Username   string `xorm:"not null default '' comment('用户名') VARCHAR(50)"`
	PrizeCode  int    `xorm:"not null default 0 comment('抽奖号码（4位的随机数）') INT(11)"`
	GiftData   string `xorm:"not null default '' comment('获奖信息') VARCHAR(255)"`
	SysCreated int    `xorm:"not null default 0 comment('创建时间') INT(11)"`
	SysIp      string `xorm:"not null default '' comment('用户抽奖的 ip') VARCHAR(50)"`
	SysStatus  int    `xorm:"not null default 0 comment('状态，0-正常，1-删除') SMALLINT(11)"`
}
