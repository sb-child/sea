package service

import (
	"context"
	"sea/app/dao"
	"sea/app/model"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

var ctx = context.Background()

var WaterInvite = waterInviteService{
	ctx: &ctx,
}

type waterInviteService struct {
	ctx *context.Context
}

func (s *waterInviteService) CreateSession() (string, error) {
	sessionId := grand.S(64, true)
	_, err := dao.WaterInvite.Ctx(*s.ctx).Data(model.WaterInvite{
		Session:         sessionId,
		SenderPublicKey: "",
	}).Insert()
	return sessionId, err
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
	_, err := dao.WaterInvite.Ctx(*s.ctx).Where(model.WaterInvite{Session: sessionId}).Delete()
	return err
}
