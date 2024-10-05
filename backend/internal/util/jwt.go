package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Access  string
	Refresh string
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func NewJWT(jwt JWT) *JWT {
	return &JWT{
		Access:  jwt.Access,
		Refresh: jwt.Refresh,
	}
}

func (j *JWT) GenerateAccessToken(userID uint) (string, error) {
	exp := time.Now().Add(10 * time.Minute)
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.Access))
}

func (j *JWT) GenerateRefreshToken(userID uint) (string, error) {
	exp := time.Now().Add(24 * time.Hour)
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.Refresh))
}

func (j *JWT) ValidateToken(token string, secret string) (*Claims, error) {
	tokenString, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !tokenString.Valid {
		return nil, errors.New("jwt token not valid")
	}

	claims, ok := tokenString.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	if claims.ExpiresAt.Unix() < time.Now().Unix() {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}
