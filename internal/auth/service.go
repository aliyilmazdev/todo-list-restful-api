package auth

import (
	"errors"

	"github.com/aliyilmazdev/todo-list-restful-api/internal/user"
	"github.com/aliyilmazdev/todo-list-restful-api/pkg/jwt"
)

type Service interface {
	Login(loginRequest LoginRequest) (*string, error)
	Register(u user.User) (*string, error)
}

type service struct {
	userService user.Service
	jwtClient jwt.Client
}

func NewService(userService user.Service, jwtClient jwt.Client) Service {
	return &service{userService: userService, jwtClient: jwtClient}
}

func (s *service) Login(loginRequest LoginRequest) (*string, error) {
	u, err := s.userService.GetUserByEmail(loginRequest.Email)
	if err != nil {
		return nil, err
	}

	if u.Password != loginRequest.Password {
		return nil, errors.New("invalid password")
	}

	token, err := s.jwtClient.GenerateToken(u.ID)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *service) Register(u user.User) (*string, error) {

	createdUser, err := s.userService.Create(u)

	if err != nil {
		return nil, err
	}

	token, err := s.jwtClient.GenerateToken(createdUser.ID)

	if err != nil {
		return nil, err
	}

	return &token, nil
}