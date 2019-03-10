package services

import (
	"imooc_go_lottery/dao"
	"imooc_go_lottery/datasource"
	"imooc_go_lottery/models"
)

type UserService interface {
	GetAll() []models.LtUser
	CountAll() int64
	Get(id int) *models.LtUser
	Delete(id int) error
	Update(data *models.LtUser, columns []string) error
	Create(data *models.LtUser) error
}

type userService struct {
	dao *dao.UserDao
}

func (s *userService) GetAll() []models.LtUser {
	return s.dao.GetAll()
}

func (s *userService) CountAll() int64 {
	panic("implement me")
}

func (s *userService) Get(id int) *models.LtUser {
	return s.dao.Get(id)
}

func (s *userService) Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *userService) Update(data *models.LtUser, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *userService) Create(data *models.LtUser) error {
	return s.dao.Create(data)
}

func NewUserService() UserService {
	return &userService{
		dao: dao.NewUserDao(datasource.InstanceDbMaster()),
	}
}
