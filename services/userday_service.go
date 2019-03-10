package services

import (
	"imooc_go_lottery/comm"
	"imooc_go_lottery/dao"
	"imooc_go_lottery/datasource"
	"imooc_go_lottery/models"
)

type UserdayService interface {
	GetAll() []models.LtUserday
	CountAll() int64
	Get(id int) *models.LtUserday
	//Delete(id int) error
	Update(data *models.LtUserday, columns []string) error
	Create(data *models.LtUserday) error
	GetUserToday(int) *models.LtUserday
}

type userdayService struct {
	dao *dao.UserdayDao
}

func (s *userdayService) GetUserToday(uid int) *models.LtUserday {
	intDay := comm.GetTodayIntDay()
	list := s.dao.Search(uid, intDay)
	if list != nil && len(list) > 0 {
		return &list[0]
	} else {
		return nil
	}
}

func (s *userdayService) GetAll() []models.LtUserday {
	return s.dao.GetAll()
}

func (s *userdayService) CountAll() int64 {
	panic("implement me")
}

func (s *userdayService) Get(id int) *models.LtUserday {
	return s.dao.Get(id)
}

//func (s *userdayService) Delete(id int) error {
//	return s.dao.Delete(id)
//}

func (s *userdayService) Update(data *models.LtUserday, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *userdayService) Create(data *models.LtUserday) error {
	return s.dao.Create(data)
}

func NewUserdayService() UserdayService {
	return &userdayService{
		dao: dao.NewUserdayDao(datasource.InstanceDbMaster()),
	}
}
