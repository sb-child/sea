package service

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"sea/internal/consts"
	model "sea/internal/model/entity"
	"sea/internal/service/internal/dao"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/encoding/gbinary"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
)

// manage all keys and sessions

var WaterKey = waterKeyService{}

type waterKeyService struct{}

type waterKey struct {
	id  string
	ctx *context.Context
}

// GetSelfKeyID returns the self key ID (same as water ID)
func (s *waterKeyService) GetSelfKey(ctx context.Context) (waterKey, error) {
	m := new(model.Water)
	err := dao.Water.Ctx(ctx).
		Where(dao.Water.Columns().IsSelf, true).
		Fields(dao.Water.Columns().WaterId).
		Scan(m)
	if err != nil {
		return waterKey{}, err
	}
	return waterKey{id: m.WaterId, ctx: &ctx}, nil
}

// AddKey add a key to the database
func (s *waterKeyService) AddKey(ctx context.Context, key *rsa.PublicKey) (waterKey, error) {
	e := CheckPublicKey(key)
	if e != consts.WATER_KEY_CHECK_OK {
		return waterKey{}, gerror.New("key check failed")
	}
	kid, _ := GetKeyID(key)
	kpack, _ := PackPublicKey(key)
	m := &model.Water{
		WaterId: kid,
		Key:     kpack,
		IsSelf:  false,
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
	if e != consts.WATER_KEY_CHECK_OK {
		return waterKey{}, gerror.New("key check failed")
	}
	kid, _ := GetKeyID(&key.PublicKey)
	kpack, _ := PackPrivateKey(key)
	m := &model.Water{
		WaterId:    kid,
		Key:        kpack,
		IsSelf:     true,
		IsVerified: true,
		IsReviewed: true,
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
	err := dao.Water.Ctx(ctx).Where(dao.Water.Columns().WaterId, id).Scan(m)
	if err != nil {
		return waterKey{}, err
	}
	return waterKey{id: m.WaterId, ctx: &ctx}, nil
}

func (s *waterKeyService) GetKey(ctx context.Context, k *rsa.PublicKey) (waterKey, error) {
	e := CheckPublicKey(k)
	if e != consts.WATER_KEY_CHECK_OK {
		return waterKey{}, gerror.New("key check failed")
	}
	kid, err := GetKeyID(k)
	if err != nil {
		return waterKey{}, err
	}
	return s.GetKeyByID(ctx, kid)
}

func (s *waterKey) getKey() (string, error) {
	m := new(model.Water)
	err := dao.Water.Ctx(*s.ctx).Where(dao.Water.Columns().WaterId, s.id).
		Fields(dao.Water.Columns().Key).
		Scan(m)
	if err != nil {
		fmt.Println(s.id)
		return "", err
	}
	return m.Key, nil
}

func (s *waterKey) GetPrivateKey() (*rsa.PrivateKey, error) {
	k, err := s.getKey()
	if err != nil {
		return nil, err
	}
	return UnpackPrivateKey(k)
}

func (s *waterKey) GetPublicKey() (*rsa.PublicKey, error) {
	k, err := s.getKey()
	if err != nil {
		return nil, err
	}
	return UnpackPublicKey(k, true)
}

func (s *waterKey) GetKeyID() string {
	return s.id
}

// todo: unsafe code
func (s *waterKey) SetKey(k string) error {
	_, err := dao.Water.Ctx(*s.ctx).
		Where(model.Water{WaterId: s.id}).
		Update(model.Water{Key: k})
	return err
}

func (s *waterKey) GetKeySession() (string, error) {
	m := new(model.Water)
	err := dao.Water.Ctx(*s.ctx).Where(dao.Water.Columns().WaterId, s.id).
		Fields(dao.Water.Columns().VerifySession).
		Scan(m)
	if err != nil {
		return "", err
	}
	return m.VerifySession, nil
}

func (s *waterKey) SetKeySession(sessionId string) error {
	_, err := dao.Water.Ctx(*s.ctx).Where(dao.Water.Columns().WaterId, s.id).Data(dao.Water.Columns().VerifySession, sessionId).Update()
	return err
}

func (s *waterKey) SetKeySessionRandom() (string, error) {
	sessionId := grand.S(64, false)
	s.SetKeySession(sessionId)
	return sessionId, nil
}

func (s *waterKey) GetStatus() int {
	m := new(model.Water)
	err := dao.Water.Ctx(*s.ctx).Where(dao.Water.Columns().WaterId, s.id).
		Fields(
			dao.Water.Columns().IsBanned,
			dao.Water.Columns().IsReviewed,
		).
		Scan(m)
	if err != nil {
		return consts.WATER_KEY_STATUS_NOT_FOUND
	}
	if m.IsBanned {
		return consts.WATER_KEY_STATUS_BANNED
	}
	if !m.IsReviewed {
		return consts.WATER_KEY_STATUS_WAIT_FOR_RESULT
	}
	return consts.WATER_KEY_STATUS_OK
}

func (s *waterKey) IsBanned() bool {
	m := new(model.Water)
	err := dao.Water.Ctx(*s.ctx).Where(dao.Water.Columns().WaterId, s.id).
		Fields(dao.Water.Columns().IsBanned).
		Scan(m)
	if err != nil {
		return false
	}
	return m.IsBanned
}

func (s *waterKey) SetBanned(b bool) error {
	_, err := dao.Water.Ctx(*s.ctx).Where(dao.Water.Columns().WaterId, s.id).Data(dao.Water.Columns().IsBanned, b).Update()
	return err
}

func (s *waterKey) IsReviewed() bool {
	m := new(model.Water)
	err := dao.Water.Ctx(*s.ctx).Where(dao.Water.Columns().WaterId, s.id).
		Fields(dao.Water.Columns().IsReviewed).
		Scan(m)
	if err != nil {
		return false
	}
	return m.IsReviewed
}

func (s *waterKey) SetReviewed(b bool) error {
	_, err := dao.Water.Ctx(*s.ctx).Where(dao.Water.Columns().WaterId, s.id).Data(dao.Water.Columns().IsReviewed, b).Update()
	return err
}

func (s *waterKey) IsVerified() bool {
	m := new(model.Water)
	err := dao.Water.Ctx(*s.ctx).Where(dao.Water.Columns().WaterId, s.id).
		Fields(dao.Water.Columns().IsVerified).
		Scan(m)
	if err != nil {
		return false
	}
	return m.IsVerified
}

func (s *waterKey) SetVerified(b bool) error {
	_, err := dao.Water.Ctx(*s.ctx).Where(dao.Water.Columns().WaterId, s.id).Data(dao.Water.Columns().IsVerified, b).Update()
	return err
}

func (s *waterKey) DeleteKey() error {
	_, err := dao.Water.Ctx(*s.ctx).Where(dao.Water.Columns().WaterId, s.id).Delete()
	return err
}

func (s *waterKey) EncryptBytes(m []byte) ([]byte, error) {
	k, err := s.GetPublicKey()
	if err != nil {
		return nil, err
	}
	hash := sha512.New()
	sz := k.Size() - 2*hash.Size() - 2
	var buf bytes.Buffer
	for len(m) > 0 {
		if len(m) <= sz {
			b, _ := rsa.EncryptOAEP(hash, rand.Reader, k, m, nil)
			buf.Write(b)
			break
		}
		b, _ := rsa.EncryptOAEP(hash, rand.Reader, k, m[:sz], nil)
		buf.Write(b)
		m = m[sz:]
	}
	return buf.Bytes(), nil
}

func (s *waterKey) DecryptBytes(m []byte) ([]byte, error) {
	k, err := s.GetPrivateKey()
	if err != nil {
		return nil, err
	}
	hash := sha512.New()
	sz := k.Size()
	var buf bytes.Buffer
	for len(m) > 0 {
		if len(m) <= sz {
			b, err := rsa.DecryptOAEP(hash, rand.Reader, k, m, nil)
			if err != nil {
				return nil, err
			}
			buf.Write(b)
			break
		}
		b, err := rsa.DecryptOAEP(hash, rand.Reader, k, m[:sz], nil)
		if err != nil {
			return nil, err
		}
		buf.Write(b)
		m = m[sz:]
	}
	return buf.Bytes(), nil
}

func (s *waterKey) EncryptString(m string) (string, error) {
	mBytes, err := s.EncryptBytes([]byte(m))
	if err != nil {
		return "", err
	}
	return gbase64.EncodeToString(mBytes), nil
}

func (s *waterKey) DecryptString(m string) (string, error) {
	mBytes, err := gbase64.DecodeString(m)
	if err != nil {
		return "", err
	}
	mBytes, err = s.DecryptBytes(mBytes)
	if err != nil {
		return "", err
	}
	return string(mBytes), nil
}

func (s *waterKey) EncryptJsonBase64(m *gvar.Var) (string, error) {
	b, err := m.MarshalJSON()
	if err != nil {
		return "", err
	}
	c, err := s.EncryptBytes(b)
	if err != nil {
		return "", err
	}
	return gbase64.EncodeToString(c), nil
}

func (s *waterKey) DecryptJsonBase64(m string) (*gvar.Var, error) {
	b, err := gbase64.DecodeString(m)
	if err != nil {
		return nil, err
	}
	c, err := s.DecryptBytes(b)
	if err != nil {
		return nil, err
	}
	r := gvar.New("")
	err = r.UnmarshalJSON(c)
	return r, err
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
	block := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: k,
	}
	return string(pem.EncodeToMemory(block)), nil
}

func UnpackPublicKey(key string, check bool) (*rsa.PublicKey, error) {
	// use x509 pkcs1 to unpack public key
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return nil, gerror.New("invalid public key")
	}
	k, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if (err != nil) && check {
		// maybe it's private key
		pk, err := UnpackPrivateKey(key)
		if err != nil {
			return nil, err
		}
		return &pk.PublicKey, nil
	}
	return k, err
}

func PackPrivateKey(key *rsa.PrivateKey) (string, error) {
	// use x509 pkcs1 to pack private key
	k := x509.MarshalPKCS1PrivateKey(key)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: k,
	}
	return string(pem.EncodeToMemory(block)), nil
}

func UnpackPrivateKey(key string) (*rsa.PrivateKey, error) {
	// use x509 pkcs1 to unpack private key
	block, _ := pem.Decode([]byte(key))
	k, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	return k, err
}

func CheckPublicKey(key *rsa.PublicKey) int {
	if v := key.Size(); v != (4096 / 8) { // 4096 bits = 4096/8 bytes
		g.Log().Debugf(context.Background(), "wrong public key size: %d", v)
		return consts.WATER_KEY_CHECK_WRONG_SIZE
	}
	return consts.WATER_KEY_CHECK_OK
}

func CheckPrivateKey(key *rsa.PrivateKey) int {
	if err := key.Validate(); err != nil {
		return consts.WATER_KEY_CHECK_TEST_FAILED
	}
	if v := key.Size(); v != (4096 / 8) { // 4096 bits = 4096/8 bytes
		g.Log().Debugf(context.Background(), "wrong private key size: %d", v)
		return consts.WATER_KEY_CHECK_WRONG_SIZE
	}
	return consts.WATER_KEY_CHECK_OK
}

func GetKeyID(key *rsa.PublicKey) (string, error) {
	// use sha512 to get fingerprint
	h := sha512.New()
	h.Write(key.N.Bytes())
	// convert key.E to byte[]
	e := gbinary.EncodeInt(key.E)
	h.Write(e)
	// get hash
	hashed := h.Sum(nil)
	// convert to base64
	sum := gbase64.EncodeToString(hashed)
	return sum, nil
}
