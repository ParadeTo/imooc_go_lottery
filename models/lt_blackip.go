package models

type LtBlackip struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	Ip         string `xorm:"not null default '' comment('ip地址') VARCHAR(50)"`
	Blacktime  int    `xorm:"not null default 0 comment('限制到期时间') INT(11)"`
	SysCreated int    `xorm:"not null default 0 comment('创建时间') INT(11)"`
	SysUpdated int    `xorm:"not null default 0 comment('更新时间') INT(11)"`
}
