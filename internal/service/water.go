package service

import (
	model "sea/internal/model/entity"
	"sea/internal/service/internal/dao"

	"github.com/gogf/gf/v2/util/grand"
)

var Water = waterService{}

type waterService struct{}

func (s *waterService) GetSelfWater() (model.Water, error) {
	result, _ := dao.Water.DB().Model("water").Where("self", true).One()
	resultStruct := model.Water{}
	err := result.Struct(&resultStruct)
	return resultStruct, err
}

func (s *waterService) ReGenWaterID() error {
	dao.Water.DB().Model("water").Where("self", true).Delete()
	dao.Water.DB().Model("water").Data(model.Water{
		WaterId: grand.S(256, false),
		IsSelf:  true,
	}).Insert()
	return nil
}
