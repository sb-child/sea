package service

import (
	"context"
	"sea/internal/service/internal/dao"

	"github.com/gogf/gf/v2/container/gvar"
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
	INVITE_RETURN_CODE_BANNED             = 8 // key is banned
)

type WaterInviteStep1Pack struct {
	Session                    string `json:"session"` // a 64 character random string
	SenderPublicKeyFingerprint string `json:"sender"`
	ReceiverPublicKey          string `json:"receiver"` // public key of the receiver(server)
}

type WaterInviteStep2Pack struct {
	Session      string `json:"session"`
	RandomString string `json:"random"` // a 32 character random string
}

func (*waterInviteService) MakeStep1Pack(session string, recvKey, sendKey *waterKey) *gvar.Var {
	k, err := recvKey.GetPublicKey()
	if err != nil {
		return nil
	}
	ks, err := PackPublicKey(k)
	if err != nil {
		return nil
	}
	r := WaterInviteStep1Pack{
		Session:                    session,
		ReceiverPublicKey:          ks,
		SenderPublicKeyFingerprint: sendKey.GetKeyID(),
	}
	return gvar.New(r)
}

func (*waterInviteService) MakeStep2Pack(session string, random string) *gvar.Var {
	r := WaterInviteStep2Pack{
		Session:      session,
		RandomString: random,
	}
	return gvar.New(r)
}

func (s *waterInviteService) InviteStep1(c context.Context, senderPublicKey string) (encryptedReceiverPublicKey string, returnCode int) {
	wrap := func(ctx context.Context, tx *gdb.TX) error {
		encryptedReceiverPublicKey, returnCode = s.inviteStep1(ctx, tx, senderPublicKey)
		if returnCode != INVITE_RETURN_CODE_SUCCESS {
			return gerror.Newf("InviteStep1: %d", returnCode)
		}
		return nil
	}
	dao.Water.Transaction(c, func(ctx context.Context, tx *gdb.TX) error {
		return wrap(ctx, tx)
	})
	return
}
func (s *waterInviteService) inviteStep1(ctx context.Context, tx *gdb.TX, senderPublicKey string) (string, int) {
	// ensure this key is valid
	k, err := UnpackPublicKey(senderPublicKey, false)
	if err != nil {
		return "", INVITE_RETURN_CODE_BAD_KEY
	}
	if CheckPublicKey(k) != WATER_KEY_CHECK_OK {
		return "", INVITE_RETURN_CODE_BAD_KEY
	}
	// the key is valid, now check if it's banned
	wk, err := WaterKey.GetKey(ctx, k)
	// if the key is not found, it's not banned
	// if the key is found, but the status is not banned, it's not banned
	if (err == nil) && (wk.IsBanned()) {
		return "", INVITE_RETURN_CODE_BANNED
	}
	// add it to database
	senderWaterKey, err := WaterKey.AddKey(ctx, k)
	if err != nil {
		return "", INVITE_RETURN_CODE_KEY_ALREADY_EXISTS
	}
	// bind the key to a new session
	session, err := senderWaterKey.SetKeySessionRandom()
	if err != nil {
		return "", INVITE_RETURN_CODE_SESSION_ERROR
	}
	// get self key from database
	selfWaterKey, err := WaterKey.GetSelfKey(ctx)
	if err != nil {
		return "", INVITE_RETURN_CODE_SERVER_ERROR
	}
	// encrypt and response
	es, err := senderWaterKey.EncryptJsonBase64(
		s.MakeStep1Pack(
			session,
			&selfWaterKey,
			&senderWaterKey,
		),
	)
	if err != nil {
		return "", INVITE_RETURN_CODE_SERVER_ERROR
	}
	return es, INVITE_RETURN_CODE_SUCCESS
}
func (s *waterInviteService) InviteStep2(c context.Context, encryptedRandomString string) (returnCode int) {
	wrap := func(ctx context.Context, tx *gdb.TX) error {
		returnCode = s.inviteStep2(ctx, tx, encryptedRandomString)
		if returnCode != INVITE_RETURN_CODE_SUCCESS {
			return gerror.Newf("InviteStep2: %d", returnCode)
		}
		return nil
	}
	dao.Water.Transaction(c, func(ctx context.Context, tx *gdb.TX) error {
		return wrap(ctx, tx)
	})
	return
}
func (s *waterInviteService) inviteStep2(ctx context.Context, tx *gdb.TX, encryptedRandomString string) int {
	// get self key from database
	selfWaterKey, err := WaterKey.GetSelfKey(ctx)
	if err != nil {
		return INVITE_RETURN_CODE_SERVER_ERROR
	}
	// decrypt
	r, err := selfWaterKey.DecryptJsonBase64(encryptedRandomString)
	if err != nil {
		return INVITE_RETURN_CODE_DECRYPTION_FAILED
	}
	_ = r // todo
	return 0
}
