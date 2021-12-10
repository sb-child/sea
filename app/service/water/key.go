package service

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/binary"
	"encoding/hex"
	"sea/app/dao"
	"sea/app/model"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/grand"
	"github.com/gogf/gf/v2/errors/gerror"
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
	WATER_KEY_CHECK_WRONG_SIZE       = 2
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
func (s *waterKeyService) AddKey(ctx context.Context, key *rsa.PublicKey, self bool) (waterKey, error) {
	e := CheckPublicKey(key)
	if e != WATER_KEY_CHECK_OK {
		return waterKey{}, gerror.New("key check failed")
	}
	kid, _ := GetKeyID(key)
	kpack, _ := PackPublicKey(key)
	m := &model.Water{
		WaterId: kid,
		Key:     kpack,
		IsSelf:  self,
	}
	_, err := dao.Water.Ctx(ctx).Insert(m)
	if err != nil {
		return waterKey{}, err
	}
	return waterKey{id: m.WaterId, ctx: &ctx}, nil
}

func (s *waterKeyService) AddSelfKey(ctx context.Context, key *rsa.PrivateKey) (waterKey, error) {
	if _, err := s.GetSelfKey(ctx); err == nil {
		return waterKey{}, gerror.New("key already exists")
	}
	e := CheckPrivateKey(key)
	if e != WATER_KEY_CHECK_OK {
		return waterKey{}, gerror.New("key check failed")
	}
	kid, _ := GetKeyID(&key.PublicKey)
	kpack, _ := PackPrivateKey(key)
	m := &model.Water{
		WaterId: kid,
		Key:     kpack,
		IsSelf:  true,
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

func (s *waterKey) getKey() (string, error) {
	m := new(model.Water)
	err := dao.Water.Ctx(*s.ctx).Where(dao.Water.WaterDao.Columns.WaterId, s.id).Scan(m)
	if err != nil {
		return "", err
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
	ks, _ := k.ArmorWithCustomHeaders("", "")
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
	kps, _ := kp.ArmorWithCustomHeaders("", "")
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

// --- utils ---

func GenerateKey() (*rsa.PrivateKey, error) {
	// generate a new rsa 4096 bits key
	k, err := rsa.GenerateKey(rand.Reader, 4096)
	return k, err
}

func PackPublicKey(key *rsa.PublicKey) (string, error) {
	// use x509 pkcs1 to pack public key
	k := x509.MarshalPKCS1PublicKey(key)
	return string(k), nil
}

func UnpackPublicKey(key string) (*rsa.PublicKey, error) {
	// use x509 pkcs1 to unpack public key
	k, err := x509.ParsePKCS1PublicKey([]byte(key))
	return k, err
}

func PackPrivateKey(key *rsa.PrivateKey) (string, error) {
	// use x509 pkcs1 to pack private key
	k := x509.MarshalPKCS1PrivateKey(key)
	return string(k), nil
}

func UnpackPrivateKey(key string) (*rsa.PrivateKey, error) {
	// use x509 pkcs1 to unpack private key
	k, err := x509.ParsePKCS1PrivateKey([]byte(key))
	return k, err
}

func CheckPublicKey(key *rsa.PublicKey) int {
	if key.Size() != 4096 {
		return WATER_KEY_CHECK_WRONG_SIZE
	}
	return WATER_KEY_CHECK_OK
}

func CheckPrivateKey(key *rsa.PrivateKey) int {
	if err := key.Validate(); err != nil {
		g.Log().Error(err)
		return WATER_KEY_CHECK_TEST_FAILED
	}
	if key.Size() != 4096 {
		return WATER_KEY_CHECK_WRONG_SIZE
	}
	return WATER_KEY_CHECK_OK
}

// todo
func GetKeyID(key *rsa.PublicKey) (string, error) {
	// use sha512 to get fingerprint
	h := sha512.New()
	h.Write(key.N.Bytes())
	// convert key.E to byte[]
	e := make([]byte, 4)
	binary.BigEndian.PutUint32(e, uint32(key.E))
	h.Write(e)
	// get hash
	hashed := h.Sum(nil)
	// convert to hex
	sum := hex.EncodeToString(hashed)
	return sum, nil
}
