package security

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	configuration "safehouse-main-back/src/internal/config"
)

type JWTClaims struct {
	jwt.RegisteredClaims
	ClientType string `json:"client_type"`
}

type JWTManager struct {
	signingKey []byte
	config     configuration.Config
}

func NewJWTManager(config configuration.Config) *JWTManager {
	return &JWTManager{
		signingKey: []byte(config.JWTSigningKey),
		config:     config,
	}
}

func (j *JWTManager) GenerateToken() (string, error) {
	claims := &JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.config.JWTExpirationMinutes) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   "frontend-client",
		},
		ClientType: "frontend",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(j.signingKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}

func (j *JWTManager) ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.signingKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}
