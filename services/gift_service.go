package services

import (
	"imooc_go_lottery/dao"
	"imooc_go_lottery/datasource"
	"imooc_go_lottery/models"
	"strconv"
	"strings"
)

type GiftService interface {
	GetAll() []models.LtGift
	CountAll() int64
	Get(id int) *models.LtGift
	Delete(id int) error
	Update(data *models.LtGift, columns []string) error
	Create(data *models.LtGift) error
	GetAllUse() []models.ObjGiftPrize
	DecrLeftNum(id, num int) (int64, error)
	IncrLeftNum(id, num int) (int64, error)
}

type giftService struct {
	dao *dao.GiftDao
}

func (s *giftService) GetAll() []models.LtGift {
	return s.dao.GetAll()
}

func (s *giftService) CountAll() int64 {
	panic("implement me")
}

func (s *giftService) Get(id int) *models.LtGift {
	return s.dao.Get(id)
}

func (s *giftService) Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *giftService) Update(data *models.LtGift, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *giftService) Create(data *models.LtGift) error {
	return s.dao.Create(data)
}

func (s *giftService) GetAllUse() []models.ObjGiftPrize {
	list := make([]models.LtGift, 0)
	list = s.dao.GetAllUse()
	if list != nil {
		gifts := make([]models.ObjGiftPrize, 0)
		for _, gift := range list {
			codes := strings.Split(gift.PrizeCode, "-")
			if len(codes) == 2 {
				codeA := codes[0]
				codeB := codes[1]
				a, e1 := strconv.Atoi(codeA)
				b, e2 := strconv.Atoi(codeB)
				if e1 == nil && e2 == nil && b >= a && a >= 0 && b <= 10000 {
					data := models.ObjGiftPrize{
						Id:           gift.Id,
						Title:        gift.Title,
						PrizeNum:     gift.PrizeNum,
						LeftNum:      gift.LeftNum,
						PrizeCodeA:   a,
						PrizeCodeB:   b,
						Img:          gift.Img,
						Displayorder: gift.Displayorder,
						Gtype:        gift.Gtype,
						Gdata:        gift.Gdata,
					}
					gifts = append(gifts, data)
				}
			}
		}
		return gifts
	}
	return []models.ObjGiftPrize{}
}

func (s *giftService) DecrLeftNum(id, num int) (int64, error) {
	return s.dao.DecrLeftNum(id, num)
}

//IncrLeftNum
func (s *giftService) IncrLeftNum(id, num int) (int64, error) {
	return s.dao.IncrLeftNum(id, num)
}

func NewGiftService() GiftService {
	return &giftService{
		dao: dao.NewGiftDao(datasource.InstanceDbMaster()),
	}
}
