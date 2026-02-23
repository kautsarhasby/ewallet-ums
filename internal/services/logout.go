package services

import (
	"context"
	"kautsarhasby/ewallet-ums/internal/interfaces"

	"github.com/pkg/errors"
)

type LogoutService struct {
	UserRepository interfaces.IUserRepository
}

func (s *LogoutService) Logout(ctx context.Context, token string) error {

	err := s.UserRepository.DeleteUserSession(ctx, token)
	if err != nil {
		return errors.Wrap(err, "failed to delte user session")
	}

	return nil
}
