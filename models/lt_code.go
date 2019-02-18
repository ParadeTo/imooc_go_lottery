package models

type LtCode struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	GiftId     int    `xorm:"not null default 0 comment('奖品 id，关联 lt_gift 表') INT(11)"`
	Code       string `xorm:"not null default '' comment('虚拟券编码') VARCHAR(255)"`
	SysCreated int    `xorm:"not null default 0 comment('创建时间') INT(11)"`
	SysUpdated int    `xorm:"not null default 0 comment('更新时间') INT(11)"`
	SysStatus  int    `xorm:"not null default 0 comment('状态，0-正常，1-删除') SMALLINT(11)"`
}
