package tokenManager

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"time"
)

type tokenClaims struct {
	jwt.StandardClaims
	Ip string `json:"ip"`
}

type Manager struct {
	signingKey string
}

func NewManager(signingKey string) *Manager {
	return &Manager{signingKey: signingKey}
}

func (m *Manager) NewAccessToken(ipClient string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&tokenClaims{
			Ip: ipClient,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(ttl).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
		},
	)
	return token.SignedString([]byte(m.signingKey))
}

func (m *Manager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	if _, err := r.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func (m *Manager) HashRefreshToken(token string) (string, error) {
	hash := sha256.New()
	hash.Write([]byte(token))
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func (m *Manager) GenerateEmailConfirmationToken(id string) (string, error) {
	return "", nil
}
