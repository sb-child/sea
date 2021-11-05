package service

import (
	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/gogf/gf/v2/errors/gerror"
)

var WaterKey = waterKeyService{}

type waterKeyService struct{}

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

// GetSelfKeyID returns the self key ID
func (s *waterKeyService) GetSelfKeyID() (string, error) {
	return "", nil
}

// GetSelfKeyID returns all the keys ID stored
func (s *waterKeyService) GetKeyIDList() []string {
	return make([]string, 0)
}

func (s *waterKeyService) GetKey(id string) (string, error) {
	return "", nil
}

func (s *waterKeyService) GetKeySession(id string) (string, error) {
	return "", nil
}

func (s *waterKeyService) SetKey(id, key string, self bool) error {
	r := s.CheckKey(key, self)
	if r != WATER_KEY_CHECK_OK {
		return gerror.Newf("check failed: %d", r)
	}
	return nil
}

func (s *waterKeyService) SetKeySession(id, sessionId string) error {
	return nil
}

func (s *waterKeyService) SetKeySessionRandom(id string) (string, error) {
	return "", nil
}

func (s *waterKeyService) CheckKey(key string, self bool) int {
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

func (s *waterKeyService) GetKeyStatus(id string) int {
	return WATER_KEY_STATUS_OK
}

func (s *waterKeyService) SetKeyStatus(id string) error {
	return nil
}

func (s *waterKeyService) DeleteKey(id string) error {
	return nil
}
