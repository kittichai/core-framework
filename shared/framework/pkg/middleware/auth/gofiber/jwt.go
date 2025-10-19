// shared/pkg/auth/jwt.go
package gofiber

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kittichai/core-framework/shared/framework/pkg/errors"
)

type JWTAuth struct {
	secret string
}

func NewJWTAuth(secret string) *JWTAuth {
	return &JWTAuth{secret: secret}
}

func (a *JWTAuth) GenerateToken(userID string, roles []string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"roles":   roles,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.secret))
}

func (a *JWTAuth) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.NewAppError(401, "Invalid signing method", nil)
		}
		return []byte(a.secret), nil
	})
	if err != nil {
		return nil, errors.NewAppError(401, "Invalid token", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.NewAppError(401, "Invalid token", nil)
}
