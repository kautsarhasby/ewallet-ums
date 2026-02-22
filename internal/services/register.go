package services

import (
	"context"
	"kautsarhasby/ewallet-ums/internal/interfaces"
	"kautsarhasby/ewallet-ums/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	UserRepository interfaces.IUserRepository
}

func (s *RegisterService) Register(ctx context.Context, request models.User) (interface{}, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	request.Password = string(hashPassword)

	err = s.UserRepository.InsertUser(ctx, &request)
	if err != nil {
		return nil, err
	}

	response := request
	response.Password = ""
	return response, nil

}
