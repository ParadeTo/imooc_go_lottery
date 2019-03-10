package services

import (
	"imooc_go_lottery/dao"
	"imooc_go_lottery/datasource"
	"imooc_go_lottery/models"
)

type BlackipService interface {
	GetAll() []models.LtBlackip
	CountAll() int64
	Get(id int) *models.LtBlackip
	//Delete(id int) error
	Update(data *models.LtBlackip, columns []string) error
	Create(data *models.LtBlackip) error
	GetByIp(string) *models.LtBlackip
	//GetUserToday(int) *models.LtBlackip
}

type blackipService struct {
	dao *dao.BlackipDao
}

func (s *blackipService) GetAll() []models.LtBlackip {
	return s.dao.GetAll()
}

func (s *blackipService) CountAll() int64 {
	panic("implement me")
}

func (s *blackipService) Get(id int) *models.LtBlackip {
	return s.dao.Get(id)
}

//func (s *userdayService) Delete(id int) error {
//	return s.dao.Delete(id)
//}

func (s *blackipService) Update(data *models.LtBlackip, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *blackipService) Create(data *models.LtBlackip) error {
	return s.dao.Create(data)
}

func (s *blackipService) GetByIp(ip string) *models.LtBlackip {
	return s.dao.GetByIp(ip)
}

func NewBlackipService() BlackipService {
	return &blackipService{
		dao: dao.NewBlackipDao(datasource.InstanceDbMaster()),
	}
}
