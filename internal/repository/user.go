package repository

import (
	"context"
	"kautsarhasby/ewallet-ums/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) InsertUser(ctx context.Context, user *models.User) error {
	return r.DB.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	var (
		user models.User
		err  error
	)
	err = r.DB.WithContext(ctx).Where(`username = ?`, username).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) InsertUserSession(ctx context.Context, session *models.UserSession) error {
	return r.DB.WithContext(ctx).Create(&session).Error
}
