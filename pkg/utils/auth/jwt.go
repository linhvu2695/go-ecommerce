package auth

import (
	"errors"
	"go-ecommerce/global"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func SignTokenJWT(payload jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(global.Config.JWT.ApiSecret))
}

func CreateToken(uuidString string) (string, error) {
	timeExpiration := global.Config.JWT.JwtExpiration
	if timeExpiration == "" {
		timeExpiration = "1h"
	}

	expireDuration, err := time.ParseDuration(timeExpiration)
	if err != nil {
		return "", err
	}

	now := time.Now()
	expiresAt := now.Add(expireDuration)

	return SignTokenJWT(jwt.StandardClaims{
		Id:        uuid.NewString(),
		ExpiresAt: expiresAt.Unix(),
		IssuedAt:  now.Unix(),
		Issuer:    "go-ecommerce",
		Subject:   uuidString,
	})
}

func VerifyTokenSubject(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Make sure the signing method is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(global.Config.JWT.ApiSecret), nil
	})
	if err != nil {
		return nil, err
	}

	// Extract and validate claims
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		if claims.ExpiresAt < time.Now().Unix() {
			return nil, errors.New("token expired")
		}
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
