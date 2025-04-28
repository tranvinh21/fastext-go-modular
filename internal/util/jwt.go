package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/vinhtran21/fastext-go-modular/config"
	entity "github.com/vinhtran21/fastext-go-modular/domains/entities"
)

func GenerateJWT(user *entity.User, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    user.ID,
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secret))
}

func VerifyJWT(token string, secret string) (uint, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}

	userId := claims["userId"].(uint)
	return userId, nil
}

func GenerateRefreshToken(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    userId,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	return token.SignedString([]byte(config.Envs.AuthConfig.RefreshTokenSecret))
}

func GenerateAccessToken(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    userId,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(config.Envs.AuthConfig.AccessTokenSecret))
}

func VerifyAccessToken(token string) (uint, error) {
	return VerifyJWT(token, config.Envs.AuthConfig.AccessTokenSecret)
}

func VerifyRefreshToken(token string) (uint, error) {
	return VerifyJWT(token, config.Envs.AuthConfig.RefreshTokenSecret)
}
