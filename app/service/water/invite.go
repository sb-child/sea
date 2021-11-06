package service

import (
	"context"
	"sea/app/dao"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/gogf/gf/encoding/gparser"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

var WaterInvite = waterInviteService{}

type waterInviteService struct{}

const (
	INVITE_RETURN_CODE_SUCCESS            = 0 // success
	INVITE_RETURN_CODE_DECRYPTION_FAILED  = 1 // failed to decrypt
	INVITE_RETURN_CODE_SESSION_NOT_FOUND  = 2 // session not found
	INVITE_RETURN_CODE_SESSION_ERROR      = 3 // can't create session, needs a retry
	INVITE_RETURN_CODE_BAD_KEY            = 4 // invalid key, expired, a private key, banned key or empty string
	INVITE_RETURN_CODE_BAD_RANDOM_STRING  = 5 // random string is not 32 characters long
	INVITE_RETURN_CODE_KEY_ALREADY_EXISTS = 6 // this key already exists
	INVITE_RETURN_CODE_SERVER_ERROR       = 7 // server isn't ready
)

type WaterInviteStep1Pack struct {
	Session           string `json:"session"` // a 64 character random string
	ReceiverPublicKey string `json:"receiver"`
}

type WaterInviteStep2Pack struct {
	Session      string `json:"session"`
	RandomString string `json:"random"` // a 32 character random string
}

func (*waterInviteService) MakeStep1Pack(session, key string) string {
	r := gparser.New(WaterInviteStep1Pack{Session: session, ReceiverPublicKey: key})
	return r.MustToJsonString()
}
func (s *waterInviteService) InviteStep1(c context.Context, senderPublicKey string) (EncryptedReceiverPublicKey string, ReturnCode int) {
	wrap := func(ctx context.Context, tx *gdb.TX) error {
		EncryptedReceiverPublicKey, ReturnCode = s.inviteStep1(ctx, tx, senderPublicKey)
		if ReturnCode != INVITE_RETURN_CODE_SUCCESS {
			return gerror.Newf("InviteStep1: %d", ReturnCode)
		}
		return nil
	}
	dao.Water.Transaction(c, func(ctx context.Context, tx *gdb.TX) error {
		return wrap(ctx, tx)
	})
	return
}
func (s *waterInviteService) inviteStep1(ctx context.Context, tx *gdb.TX, senderPublicKey string) (string, int) {
	k, err := crypto.NewKeyFromArmored(senderPublicKey)
	ks, _ := k.Armor()
	if kstat := WaterKey.GetKeyStatus(ctx, ks); (err != nil) ||
		(WaterKey.CheckKey(ctx, ks, false) != WATER_KEY_CHECK_OK) ||
		(kstat != WATER_KEY_STATUS_NOT_FOUND) {
		return "", INVITE_RETURN_CODE_BAD_KEY
	}
	senderPublicKeyID, err := WaterKey.AddKey(ctx, ks)
	if err != nil {
		return "", INVITE_RETURN_CODE_KEY_ALREADY_EXISTS
	}
	WaterKey.SetKeyStatus(ctx, senderPublicKeyID, WATER_KEY_STATUS_WAIT_FOR_RESULT)
	session, err := WaterKey.SetKeySessionRandom(ctx, senderPublicKeyID)
	if err != nil {
		return "", INVITE_RETURN_CODE_SESSION_ERROR
	}
	selfKeyID, err := WaterKey.GetSelfKeyID(ctx)
	if err != nil {
		return "", INVITE_RETURN_CODE_SERVER_ERROR
	}
	selfKey, _ := WaterKey.GetKey(ctx, selfKeyID)
	es, err := helper.EncryptMessageArmored(selfKey, s.MakeStep1Pack(session, selfKey))
	if err != nil {
		return "", INVITE_RETURN_CODE_SERVER_ERROR
	}
	return es, INVITE_RETURN_CODE_SUCCESS
}
