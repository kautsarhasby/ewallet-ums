package services

import (
	"context"
	"kautsarhasby/ewallet-ums/helpers"
	"kautsarhasby/ewallet-ums/internal/interfaces"
	"kautsarhasby/ewallet-ums/internal/models"

	"github.com/pkg/errors"
)

type RefreshTokenService struct {
	UserRepository interfaces.IUserRepository
}

func (s *RefreshTokenService) RefreshToken(ctx context.Context, refreshToken string, claims helpers.TokenClaims) (models.RefreshTokenResponse, error) {
	token, err := helpers.GenerateToken(ctx, claims.UserID, claims.Username, claims.Fullname, "refreshToken")
	if err != nil {
		return models.RefreshTokenResponse{}, errors.Wrap(err, "failed to generate token")
	}

	err = s.UserRepository.UpdateTokenByRefreshToken(ctx, token, refreshToken)
	if err != nil {
		return models.RefreshTokenResponse{}, errors.Wrap(err, "failed to update token")
	}

	return models.RefreshTokenResponse{Token: token}, nil
}
