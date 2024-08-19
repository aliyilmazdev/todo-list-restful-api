package jwt

import (
	"time"

	"github.com/aliyilmazdev/todo-list-restful-api/config"
	"github.com/golang-jwt/jwt/v5"
)

type Client interface {
	GenerateToken(userID uint) (string, error)
}

type client struct {}

func NewClient() Client {
	return &client{}
}

func (c *client) GenerateToken(userID uint) (string, error) {
	day := time.Hour * 24

	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(day * 1)),
		},
		ID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(config.Config("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

