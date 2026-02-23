package interfaces

import (
	"context"
	"kautsarhasby/ewallet-ums/internal/models"
)

type IUserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
	InsertUserSession(ctx context.Context, session *models.UserSession) error
	GetUserSessionByToken(ctx context.Context, token string) (models.UserSession, error)
	DeleteUserSession(ctx context.Context, token string) error
	UpdateTokenByRefreshToken(ctx context.Context, token, refreshToken string) error
	GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (models.UserSession, error)
}
