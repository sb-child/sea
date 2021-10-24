package service

import (
	"github.com/gogf/gf/os/gtime"
)

var WaterInvite = waterInviteService{}

type waterInviteService struct {
}

func (s *waterInviteService) CreateSession() (string, error) {
	// result, _ := dao.Water.DB().Model("water").Where("self", true).One()
	// resultStruct := model.Water{}
	// err := result.Struct(&resultStruct)
	return "", nil
}

func (s *waterInviteService) SetSessionSender(publicKey string) error {
	return nil
}

func (s *waterInviteService) GetSessionSender(sessionId string) (string, error) {
	return "", nil
}

func (s *waterInviteService) GetSessionCreateTime(sessionId string) (*gtime.Time, error) {
	return gtime.New(), nil
}

func (s *waterInviteService) DeleteSession(sessionId string) error {
	return nil
}
