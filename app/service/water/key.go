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

var WaterKey = waterKeyService{}

type waterKeyService struct{}

type waterKey struct {
	id  string
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
func (s *waterKeyService) GetSelfKey(ctx context.Context) (waterKey, error) {
	return waterKey{}, nil
}

func (s *waterKeyService) AddKey(ctx context.Context, key string) (waterKey, error) {
	return waterKey{}, nil
}

func (s *waterKeyService) GetKeyByID(ctx context.Context, id string) (waterKey, error) {
	return waterKey{}, nil
}

func (s *waterKey) getKey() (*crypto.Key, error) {
	var m *model.Water
	err := dao.Water.Ctx(*s.ctx).Where(model.Water{WaterId: s.id}).Scan(&m)
	if err != nil {
		return nil, err
	}
	k, err := crypto.NewKeyFromArmored(m.Key)
	return k, err
}

func (s *waterKey) GetPrivateKey() (string, error) {
	k, err := s.getKey()
	if err != nil {
		return "", err
	}
	if !k.IsPrivate() {
		return "", gerror.New("not private key")
	}
	ks, _ := k.Armor()
	return ks, nil
}

func (s *waterKey) GetPublicKey() (string, error) {
	k, err := s.getKey()
	if err != nil {
		return "", err
	}
	kp, err := k.ToPublic()
	if err != nil {
		return "", err
	}
	kps, _ := kp.Armor()
	return kps, nil
}

func (s *waterKey) SetKey(k string) error {
	_, err := dao.Water.Ctx(*s.ctx).
		Where(model.Water{WaterId: s.id}).
		Update(model.Water{Key: k})
	return err
}

func (s *waterKey) GetKeySession() (string, error) {
	return "", nil
}

func (s *waterKey) SetKeySession(sessionId string) error {
	_, err := dao.Water.Ctx(*s.ctx).Where(model.Water{
		WaterId: s.id,
	}).Update(model.Water{
		Session: sessionId,
	})
	return err
}

func (s *waterKey) SetKeySessionRandom() (string, error) {
	sessionId := grand.S(64, true)
	s.SetKeySession(sessionId)
	return sessionId, nil
}

func (s *waterKey) GetKeyStatus() int {
	return WATER_KEY_STATUS_OK
}

func (s *waterKey) SetKeyStatus(status int) error {
	return nil
}

func (s *waterKey) DeleteKey() error {
	return nil
}

func CheckKey(key string, self bool) (string, int) {
	k, err := crypto.NewKeyFromArmored(key)
	kstring, _ := k.Armor()
	if err != nil {
		return "", WATER_KEY_CHECK_TEST_FAILED
	}
	if (!k.CanVerify()) || (!k.CanEncrypt()) {
		return "", WATER_KEY_CHECK_USELESS
	}
	if k.IsPrivate() != self {
		return "", WATER_KEY_CHECK_TYPE_ERROR
	}
	if k.IsExpired() {
		return "", WATER_KEY_CHECK_EXPIRED
	}
	return kstring, WATER_KEY_CHECK_OK
}
