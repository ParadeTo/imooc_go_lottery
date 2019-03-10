package services

import (
	"imooc_go_lottery/dao"
	"imooc_go_lottery/datasource"
	"imooc_go_lottery/models"
)

type CodeService interface {
	GetAll() []models.LtCode
	CountAll() int64
	Get(id int) *models.LtCode
	Delete(id int) error
	Update(data *models.LtCode, columns []string) error
	Create(data *models.LtCode) error
	NextUsingCode(giftId, codeId int) *models.LtCode
}

type codeService struct {
	dao *dao.CodeDao
}

func (s *codeService) GetAll() []models.LtCode {
	return s.dao.GetAll()
}

func (s *codeService) CountAll() int64 {
	panic("implement me")
}

func (s *codeService) Get(id int) *models.LtCode {
	return s.dao.Get(id)
}

func (s *codeService) Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *codeService) Update(data *models.LtCode, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *codeService) Create(data *models.LtCode) error {
	return s.dao.Create(data)
}

func (s *codeService) NextUsingCode(giftId, codeId int) *models.LtCode {
	return s.dao.NextUsingCode(giftId, codeId)
}

func NewCodeService() CodeService {
	return &codeService{
		dao: dao.NewCodeDao(datasource.InstanceDbMaster()),
	}
}
