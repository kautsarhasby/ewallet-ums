package services

import (
	"context"
	"kautsarhasby/ewallet-ums/internal/interfaces"
	"kautsarhasby/ewallet-ums/internal/models"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	UserRepository interfaces.IUserRepository
}

func (s *RegisterService) Register(ctx context.Context, request models.User) (interface{}, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate password")
	}
	request.Password = string(hashPassword)

	err = s.UserRepository.InsertUser(ctx, &request)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate password")
	}

	response := request
	response.Password = ""
	return response, nil

}
