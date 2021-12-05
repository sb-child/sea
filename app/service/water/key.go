package service

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
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
	m := new(model.Water)
	err := dao.Water.Ctx(ctx).Where(dao.Water.WaterDao.Columns.IsSelf, true).Scan(m)
	if err != nil {
		return waterKey{}, err
	}
	return waterKey{id: m.WaterId, ctx: &ctx}, nil
}

// AddKey add a key to the database
func (s *waterKeyService) AddKey(ctx context.Context, key string, self bool) (waterKey, error) {
	key, e := CheckKey(key, self)
	if e != WATER_KEY_CHECK_OK {
		return waterKey{}, gerror.New("key check failed")
	}
	kid, _ := GetKeyID(key)
	m := &model.Water{
		WaterId: kid,
		Key:     key,
		IsSelf:  self,
	}
	_, err := dao.Water.Ctx(ctx).Insert(m)
	if err != nil {
		return waterKey{}, err
	}
	return waterKey{id: m.WaterId, ctx: &ctx}, nil
}

// GetKeyByID returns the key by ID
func (s *waterKeyService) GetKeyByID(ctx context.Context, id string) (waterKey, error) {
	m := new(model.Water)
	err := dao.Water.Ctx(ctx).Where(dao.Water.WaterDao.Columns.WaterId, id).Scan(m)
	if err != nil {
		return waterKey{}, err
	}
	return waterKey{id: m.WaterId, ctx: &ctx}, nil
}

// GetKeyByString returns the key by string
func (s *waterKeyService) GetKeyByString(ctx context.Context, ks string) (waterKey, error) {
	ks, c := CheckKeyWithoutType(ks)
	if c != WATER_KEY_CHECK_OK {
		return waterKey{}, gerror.New("key check failed")
	}
	kid, err := GetKeyID(ks)
	if err != nil {
		return waterKey{}, err
	}
	return s.GetKeyByID(ctx, kid)
}

func (s *waterKey) getKey() (*crypto.Key, error) {
	m := new(model.Water)
	err := dao.Water.Ctx(*s.ctx).Where(dao.Water.WaterDao.Columns.WaterId, s.id).Scan(m)
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

func (s *waterKey) GetKeyID() string {
	return s.id
}

func (s *waterKey) SetKey(k string) error {
	_, err := dao.Water.Ctx(*s.ctx).
		Where(model.Water{WaterId: s.id}).
		Update(model.Water{Key: k})
	return err
}

func (s *waterKey) GetKeySession() (string, error) {
	m := new(model.Water)
	err := dao.Water.Ctx(*s.ctx).Where(dao.Water.WaterDao.Columns.WaterId, s.id).Scan(m)
	if err != nil {
		return "", err
	}
	return m.VerifySession, nil
}

func (s *waterKey) SetKeySession(sessionId string) error {
	_, err := dao.Water.Ctx(*s.ctx).Where(dao.Water.WaterDao.Columns.WaterId, s.id).Update(model.Water{
		VerifySession: sessionId,
	})
	return err
}

func (s *waterKey) SetKeySessionRandom() (string, error) {
	sessionId := grand.S(64, true)
	s.SetKeySession(sessionId)
	return sessionId, nil
}

func (s *waterKey) GetStatus() int {
	m := new(model.Water)
	err := dao.Water.Ctx(*s.ctx).Where(dao.Water.WaterDao.Columns.WaterId, s.id).Scan(m)
	if err != nil {
		return WATER_KEY_STATUS_NOT_FOUND
	}
	if m.IsBanned {
		return WATER_KEY_STATUS_BANNED
	}
	if !m.IsReviewed {
		return WATER_KEY_STATUS_WAIT_FOR_RESULT
	}
	return WATER_KEY_STATUS_OK
}

func (s *waterKey) IsBanned() bool {
	m := new(model.Water)
	err := dao.Water.Ctx(*s.ctx).Where(dao.Water.WaterDao.Columns.WaterId, s.id).Scan(m)
	if err != nil {
		return false
	}
	return m.IsBanned
}

func (s *waterKey) SetBanned(b bool) error {
	_, err := dao.Water.Ctx(*s.ctx).Where(dao.Water.WaterDao.Columns.WaterId, s.id).Update(model.Water{IsBanned: b})
	return err
}

func (s *waterKey) IsReviewed() bool {
	m := new(model.Water)
	err := dao.Water.Ctx(*s.ctx).Where(dao.Water.WaterDao.Columns.WaterId, s.id).Scan(m)
	if err != nil {
		return false
	}
	return m.IsReviewed
}

func (s *waterKey) SetReviewed(b bool) error {
	_, err := dao.Water.Ctx(*s.ctx).Where(dao.Water.WaterDao.Columns.WaterId, s.id).Update(model.Water{IsReviewed: b})
	return err
}

func (s *waterKey) IsVerified() bool {
	m := new(model.Water)
	err := dao.Water.Ctx(*s.ctx).Where(dao.Water.WaterDao.Columns.WaterId, s.id).Scan(m)
	if err != nil {
		return false
	}
	return m.IsVerified
}

func (s *waterKey) SetVerified(b bool) error {
	_, err := dao.Water.Ctx(*s.ctx).Where(dao.Water.WaterDao.Columns.WaterId, s.id).Update(model.Water{IsVerified: b})
	return err
}

func (s *waterKey) DeleteKey() error {
	_, err := dao.Water.Ctx(*s.ctx).Where(dao.Water.WaterDao.Columns.WaterId, s.id).Delete()
	return err
}

func CheckKey(key string, self bool) (string, int) {
	key, err := CheckKeyWithoutType(key)
	if err != WATER_KEY_CHECK_OK {
		return "", err
	}
	k, _ := crypto.NewKeyFromArmored(key)
	if k.IsPrivate() != self {
		return "", WATER_KEY_CHECK_TYPE_ERROR
	}
	return key, WATER_KEY_CHECK_OK
}

func CheckKeyWithoutType(key string) (string, int) {
	k, err := crypto.NewKeyFromArmored(key)
	kstring, _ := k.ArmorWithCustomHeaders("", "")
	if err != nil {
		return "", WATER_KEY_CHECK_TEST_FAILED
	}
	if (!k.CanVerify()) || (!k.CanEncrypt()) {
		return "", WATER_KEY_CHECK_USELESS
	}
	if k.IsExpired() {
		return "", WATER_KEY_CHECK_EXPIRED
	}
	return kstring, WATER_KEY_CHECK_OK
}

func GetKeyID(key string) (string, error) {
	k, err := crypto.NewKeyFromArmored(key)
	if err != nil {
		return "", err
	}
	kp, err := k.ToPublic()
	if err != nil {
		return "", err
	}
	// regenerate a clean key without header
	// to avoid the key id being changed
	ks, _ := kp.ArmorWithCustomHeaders("", "")
	// use sha512 to get fingerprint
	h := sha512.New()
	h.Write([]byte(ks))
	return hex.EncodeToString(h.Sum(nil)), nil
}

func MustCheckKey(key string, self bool) int {
	_, err := CheckKey(key, self)
	return err
}

func GenerateKey() (string, error) {
	k, err := crypto.GenerateKey(
		"water:"+grand.S(64, false),
		"none@none.com",
		"rsa",
		4096,
	)
	if err != nil {
		return "", err
	}
	return k.ArmorWithCustomHeaders("", "")
}
