package service

import (
	model "sea/internal/model/entity"
	"sea/internal/service/internal/dao"
)

var Water = waterService{}

type waterService struct{}

func (s *waterService) GetSelfWater() (model.Water, error) {
	result, _ := dao.Water.DB().Model("water").Where("self", true).One()
	resultStruct := model.Water{}
	err := result.Struct(&resultStruct)
	return resultStruct, err
}
