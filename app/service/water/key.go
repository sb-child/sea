package service

import (
	"context"
	"sea/app/dao"
	"sea/app/model"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/grand"
)

// manage all keys and sessions

func getCtx() *context.Context {
	ctx := context.Background()
	return &ctx
}

var WaterKey = waterKeyService{
	ctx: getCtx(),
}

type waterKeyService struct {
	ctx *context.Context
}

const (
	WATER_KEY_STATUS_OK              = 0
	WATER_KEY_STATUS_WAIT_FOR_RESULT = 1
	WATER_KEY_STATUS_BANNED          = 2
	WATER_KEY_STATUS_NOT_FOUND       = 3
	WATER_KEY_CHECK_OK               = 0
	WATER_KEY_CHECK_TEST_FAILED      = 1
	WATER_KEY_CHECK_USELESS          = 2
	WATER_KEY_CHECK_TYPE_ERROR       = 3
	WATER_KEY_CHECK_EXPIRED          = 4
)

// GetSelfKeyID returns the self key ID (same as water ID)
func (s *waterKeyService) GetSelfKeyID(ctx context.Context) (string, error) {
	return "", nil
}

// GetSelfKeyID returns all the keys ID stored
func (s *waterKeyService) GetKeyIDList(ctx context.Context) []string {
	return make([]string, 0)
}

func (s *waterKeyService) AddKey(ctx context.Context, key string) (string, error) {
	return "", nil
}

func (s *waterKeyService) GetKey(ctx context.Context, id string) (string, error) {
	return "", nil
}

func (s *waterKeyService) GetKeySession(ctx context.Context, id string) (string, error) {
	return "", nil
}

func (s *waterKeyService) SetKey(ctx context.Context, id, key string, self bool) error {
	r := s.CheckKey(ctx, key, self)
	if r != WATER_KEY_CHECK_OK {
		return gerror.Newf("check failed: %d", r)
	}
	return nil
}

func (s *waterKeyService) SetKeySession(ctx context.Context, id, sessionId string) error {
	_, err := dao.Water.Ctx(*s.ctx).Where(model.Water{
		WaterId: id,
	}).Update(model.Water{
		Session: sessionId,
	})
	return err
}

func (s *waterKeyService) SetKeySessionRandom(ctx context.Context, id string) (string, error) {
	sessionId := grand.S(64, true)
	s.SetKeySession(ctx, id, sessionId)
	return sessionId, nil
}

func (s *waterKeyService) CheckKey(ctx context.Context, key string, self bool) int {
	k, err := crypto.NewKeyFromArmored(key)
	if err != nil {
		return WATER_KEY_CHECK_TEST_FAILED
	}
	if (!k.CanVerify()) || (!k.CanEncrypt()) {
		return WATER_KEY_CHECK_USELESS
	}
	if k.IsPrivate() != self {
		return WATER_KEY_CHECK_TYPE_ERROR
	}
	if k.IsExpired() {
		return WATER_KEY_CHECK_EXPIRED
	}
	return WATER_KEY_CHECK_OK
}

func (s *waterKeyService) GetKeyStatus(ctx context.Context, id string) int {
	return WATER_KEY_STATUS_OK
}

func (s *waterKeyService) SetKeyStatus(ctx context.Context, id string, status int) error {
	return nil
}

func (s *waterKeyService) DeleteKey(ctx context.Context, id string) error {
	return nil
}
