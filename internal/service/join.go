package service

import (
	"context"
	"sea/internal/consts"
	"sea/internal/service/internal/dao"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var WaterJoin = waterJoinService{}

type waterJoinService struct{}

type WaterJoinStep1Pack struct {
	Session                    string `json:"session"` // a 64 character random string
	SenderPublicKeyFingerprint string `json:"sender"`
	ReceiverPublicKey          string `json:"receiver"` // public key of the receiver(server)
}

type WaterJoinStep2Pack struct {
	Session      string `json:"session"`
	RandomString string `json:"random"` // a 32 character random string
}

func (*waterJoinService) MakeStep1Pack(session string, recvKey, sendKey *waterKey) *gvar.Var {
	k, err := recvKey.GetPublicKey()
	if err != nil {
		return nil
	}
	ks, err := PackPublicKey(k)
	if err != nil {
		return nil
	}
	r := WaterJoinStep1Pack{
		Session:                    session,
		ReceiverPublicKey:          ks,
		SenderPublicKeyFingerprint: sendKey.GetKeyID(),
	}
	return gvar.New(r)
}

func (*waterJoinService) MakeStep2Pack(session string, random string) *gvar.Var {
	r := WaterJoinStep2Pack{
		Session:      session,
		RandomString: random,
	}
	return gvar.New(r)
}

func (s *waterJoinService) JoinStep1(c context.Context, senderPublicKey string) (encryptedReceiverPublicKey string, returnCode int, err error) {
	wrap := func(ctx context.Context, tx *gdb.TX) error {
		encryptedReceiverPublicKey, returnCode = s.joinStep1(ctx, tx, senderPublicKey)
		if returnCode != consts.JOIN_RETURN_CODE_SUCCESS {
			return gerror.Newf("error: %d", returnCode)
		}
		return nil
	}
	err = dao.Water.Transaction(c, func(ctx context.Context, tx *gdb.TX) error {
		return wrap(ctx, tx)
	})
	return
}
func (s *waterJoinService) joinStep1(ctx context.Context, tx *gdb.TX, senderPublicKey string) (string, int) {
	// ensure this key is valid
	k, err := UnpackPublicKey(senderPublicKey, false)
	if err != nil {
		return "", consts.JOIN_RETURN_CODE_BAD_KEY
	}
	if v := CheckPublicKey(k); v != consts.WATER_KEY_CHECK_OK {
		g.Log().Debugf(ctx, "key check failed: %d", v)
		return "", consts.JOIN_RETURN_CODE_BAD_KEY
	}
	// the key is valid, now check if it's banned
	wk, err := WaterKey.GetKey(ctx, k)
	// if the key is not found, it's not banned
	// if the key is found, but the status is not banned, it's not banned
	if (err == nil) && (wk.IsBanned()) {
		return "", consts.JOIN_RETURN_CODE_BANNED
	}
	// add it to database
	senderWaterKey, err := WaterKey.AddKey(ctx, k)
	if err != nil {
		return "", consts.JOIN_RETURN_CODE_KEY_ALREADY_EXISTS
	}
	// bind the key to a new session
	session, err := senderWaterKey.SetKeySessionRandom()
	if err != nil {
		return "", consts.JOIN_RETURN_CODE_SESSION_ERROR
	}
	// get self key from database
	selfWaterKey, err := WaterKey.GetSelfKey(ctx)
	if err != nil {
		g.Log().Debugf(ctx, "unable to get self key: %v", err)
		return "", consts.JOIN_RETURN_CODE_SERVER_ERROR
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
		g.Log().Debugf(ctx, "failed to encrypt: %s", err.Error())
		return "", consts.JOIN_RETURN_CODE_SERVER_ERROR
	}
	return es, consts.JOIN_RETURN_CODE_SUCCESS
}
func (s *waterJoinService) JoinStep2(c context.Context, encryptedRandomString string) (returnCode int, err error) {
	wrap := func(ctx context.Context, tx *gdb.TX) error {
		returnCode = s.joinStep2(ctx, tx, encryptedRandomString)
		if returnCode != consts.JOIN_RETURN_CODE_SUCCESS {
			return gerror.Newf("error: %d", returnCode)
		}
		return nil
	}
	err = dao.Water.Transaction(c, func(ctx context.Context, tx *gdb.TX) error {
		return wrap(ctx, tx)
	})
	return
}
func (s *waterJoinService) joinStep2(ctx context.Context, tx *gdb.TX, encryptedRandomString string) int {
	// get self key from database
	selfWaterKey, err := WaterKey.GetSelfKey(ctx)
	if err != nil {
		return consts.JOIN_RETURN_CODE_SERVER_ERROR
	}
	// decrypt
	r, _, err := selfWaterKey.DecryptJsonBase64(encryptedRandomString) // todo
	if err != nil {
		return consts.JOIN_RETURN_CODE_DECRYPTION_FAILED
	}
	_ = r // todo
	return 0
}
