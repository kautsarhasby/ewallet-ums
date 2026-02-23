package interfaces

import (
	"context"
	"kautsarhasby/ewallet-ums/helpers"
	"kautsarhasby/ewallet-ums/internal/models"
)

type IRefreshTokenService interface {
	RefreshToken(ctx context.Context, refreshToken string, claims helpers.TokenClaims) (models.RefreshTokenResponse, error)
}
