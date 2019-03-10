package dao

import (
	"github.com/go-xorm/xorm"
	"imooc_go_lottery/models"
	"log"
)

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao{engine}
}

func (d *UserDao) Get(id int) *models.LtUser {
	data := &models.LtUser{Id: id}
	ok, err := d.engine.Get(data) // it will modify the data
	if ok && err == nil {
		return data
	} else {
		data.Id = 0 // to keep the same return
		return data
	}
}

func (d *UserDao) GetAll() []models.LtUser {
	datalist := make([]models.LtUser, 0)
	err := d.engine.Asc("sys_status").
		Asc("displayorder").Find(&datalist)
	if err != nil {
		log.Println("gift_dao.GetAll error=", err)
		return datalist
	}
	return datalist
}

func (d *UserDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtUser{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *UserDao) Delete(id int) error {
	data := &models.LtUser{Id: id}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *UserDao) Update(data *models.LtUser, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *UserDao) Create(data *models.LtUser) error {
	_, err := d.engine.Insert(data)
	return err
}