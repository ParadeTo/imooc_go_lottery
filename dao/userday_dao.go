package dao

import (
	"github.com/go-xorm/xorm"
	"imooc_go_lottery/models"
	"log"
)

type UserdayDao struct {
	engine *xorm.Engine
}

func NewUserdayDao(engine *xorm.Engine) *UserdayDao {
	return &UserdayDao{engine}
}

func (d *UserdayDao) Get(id int) *models.LtUserday {
	data := &models.LtUserday{Id: id}
	ok, err := d.engine.Get(data) // it will modify the data
	if ok && err == nil {
		return data
	} else {
		data.Id = 0 // to keep the same return
		return data
	}
}

func (d *UserdayDao) Search(uid, day int) []models.LtUserday {
	datalist := make([]models.LtUserday, 0)
	err := d.engine.Where("uid = ? AND day = ?", uid, day).Find(&datalist)
	if err != nil {
		log.Println("UserdayDao.Search error=", err)
		return datalist
	}
	return datalist
}

func (d *UserdayDao) GetAll() []models.LtUserday {
	datalist := make([]models.LtUserday, 0)
	err := d.engine.Asc("sys_status").
		Asc("displayorder").Find(&datalist)
	if err != nil {
		log.Println("gift_dao.GetAll error=", err)
		return datalist
	}
	return datalist
}

func (d *UserdayDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtUserday{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *UserdayDao) Update(data *models.LtUserday, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *UserdayDao) Create(data *models.LtUserday) error {
	_, err := d.engine.Insert(data)
	return err
}