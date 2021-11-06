package service

import (
	"context"
	"sea/app/dao"
	"sea/app/model"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

var WaterInvite = waterInviteService{}

type waterInviteService struct {}

const (
	INVITE_RETURN_CODE_SUCCESS            = 0 // success
	INVITE_RETURN_CODE_DECRYPTION_FAILED  = 1 // failed to decrypt
	INVITE_RETURN_CODE_SESSION_NOT_FOUND  = 2 // session not found
	INVITE_RETURN_CODE_SESSION_ERROR      = 3 // can't create session, needs a retry
	INVITE_RETURN_CODE_BAD_KEY            = 4 // invalid key, expired, a private key, banned key or empty string
	INVITE_RETURN_CODE_BAD_RANDOM_STRING  = 5 // random string is not 32 characters long
	INVITE_RETURN_CODE_KEY_ALREADY_EXISTS = 6 // this key already exists, return it after a successful authentication
	INVITE_RETURN_CODE_SERVER_ERROR       = 7 // server isn't ready
)

func (s *waterInviteService) InviteStep1(senderPublicKey string) (EncryptedReceiverPublicKey string, ReturnCode int) {
	wrap := func(ctx context.Context, tx *gdb.TX) error {
		EncryptedReceiverPublicKey, ReturnCode = s.inviteStep1(ctx, tx, senderPublicKey)
		return nil
	}
	dao.Water.Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		return wrap(ctx, tx)
	})
	return
}
func (s *waterInviteService) inviteStep1(ctx context.Context, tx *gdb.TX, senderPublicKey string) (string, int) {
	return "", INVITE_RETURN_CODE_SUCCESS
}

// func (s *waterInviteService) CreateSession() (string, error) {
// 	sessionId := grand.S(64, true)
// 	_, err := dao.WaterInvite.Ctx(*s.ctx).Data(model.WaterInvite{
// 		Session:         sessionId,
// 		SenderPublicKey: "",
// 	}).Insert()
// 	return sessionId, err
// }

// func (s *waterInviteService) SetSessionSender(sessionId string, publicKey string) error {
// 	_, err := dao.WaterInvite.Ctx(*s.ctx).Where(model.WaterInvite{Session: sessionId}).Data(model.WaterInvite{SenderPublicKey: publicKey}).Update()
// 	return err
// }

// func (s *waterInviteService) GetSessionSender(sessionId string) (string, error) {
// 	var m *model.WaterInvite
// 	err := dao.WaterInvite.Ctx(*s.ctx).Where(model.WaterInvite{Session: sessionId}).Scan(&m)
// 	if err != nil {
// 		return "", err
// 	}
// 	return m.Session, nil
// }

// func (s *waterInviteService) GetSessionCreateTime(sessionId string) (*gtime.Time, error) {
// 	var m *model.WaterInvite
// 	err := dao.WaterInvite.Ctx(*s.ctx).Where(model.WaterInvite{Session: sessionId}).Scan(&m)
// 	if err != nil {
// 		return gtime.New(), err
// 	}
// 	return m.CreatedAt, nil
// }

// func (s *waterInviteService) DeleteSession(sessionId string) error {
// 	_, err := dao.WaterInvite.Ctx(*s.ctx).Where(model.WaterInvite{Session: sessionId}).Delete()
// 	return err
// }
