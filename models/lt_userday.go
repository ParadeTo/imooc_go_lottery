package models

type LtUserday struct {
	Id         int `xorm:"not null pk autoincr INT(11)"`
	Uid        int `xorm:"not null default 0 comment('用户id') INT(11)"`
	Day        int `xorm:"not null default 0 comment('日期') INT(11)"`
	Num        int `xorm:"not null default 0 comment('次数') INT(11)"`
	SysCreated int `xorm:"not null default 0 comment('创建时间') INT(11)"`
	SysUpdated int `xorm:"not null default 0 comment('更新时间') INT(11)"`
}
