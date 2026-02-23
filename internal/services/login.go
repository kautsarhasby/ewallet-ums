package services

import (
	"context"
	"kautsarhasby/ewallet-ums/helpers"
	"kautsarhasby/ewallet-ums/internal/interfaces"
	"kautsarhasby/ewallet-ums/internal/models"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	UserRepository interfaces.IUserRepository
}

func (s *LoginService) Login(ctx context.Context, request models.LoginRequest) (models.LoginResponse, error) {
	var (
		response models.LoginResponse
		now      = time.Now()
	)

	user, err := s.UserRepository.GetUserByUsername(ctx, request.Username)
	if err != nil {
		return response, errors.Wrap(err, "failed to get username")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return response, errors.Wrap(err, "failed to compare password")
	}
	token, err := helpers.GenerateToken(ctx, user.ID, user.Username, user.Fullname, "token")
	if err != nil {
		return response, errors.Wrap(err, "failed to generate token")
	}
	refreshToken, err := helpers.GenerateToken(ctx, user.ID, user.Username, user.Fullname, "refreshToken")
	if err != nil {
		return response, errors.Wrap(err, "failed to generate refresh token")
	}

	session := &models.UserSession{
		UserID:              user.ID,
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        now.Add(helpers.TokenType["token"]),
		RefreshTokenExpired: now.Add(helpers.TokenType["refreshToken"]),
	}

	err = s.UserRepository.InsertUserSession(ctx, session)
	if err != nil {
		return response, errors.Wrap(err, "failed to create new session")
	}

	response.UserID = user.ID
	response.Username = user.Username
	response.Email = user.Email
	response.Fullname = user.Fullname
	response.Token = token
	response.RefreshToken = refreshToken

	return response, nil
}
