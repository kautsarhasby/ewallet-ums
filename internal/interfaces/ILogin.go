package interfaces

import (
	"context"
	"kautsarhasby/ewallet-ums/internal/models"
)

type ILoginService interface {
	Login(ctx context.Context, request models.LoginRequest) (models.LoginResponse, error)
}
