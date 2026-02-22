package interfaces

import (
	"context"
	"kautsarhasby/ewallet-ums/internal/models"
)

type IUserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
	InsertUserSession(ctx context.Context, session *models.UserSession) error
}
