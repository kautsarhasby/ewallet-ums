package interfaces

import (
	"context"
	"kautsarhasby/ewallet-ums/internal/models"
)

type IRegisterService interface {
	Register(ctx context.Context, request models.User) (interface{}, error)
}
