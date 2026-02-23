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

func (r *UserRepository) GetUserSessionByToken(ctx context.Context, token string) (models.UserSession, error) {
	var (
		err  error
		user models.UserSession
	)
	err = r.DB.WithContext(ctx).Where("token = ?", token).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) DeleteUserSession(ctx context.Context, token string) error {
	return r.DB.Exec("DELETE FROM user_sessions WHERE token = ?", token).Error
}

func (r *UserRepository) GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (models.UserSession, error) {

	var (
		err     error
		session models.UserSession
	)
	err = r.DB.WithContext(ctx).Where("refresh_token = ?", refreshToken).First(&session).Error
	if err != nil {
		return session, err
	}
	return session, nil
}

func (r *UserRepository) UpdateTokenByRefreshToken(ctx context.Context, token, refreshToken string) error {
	return r.DB.Exec("UPDATE user_sessions SET token = ? WHERE refresh_token = ?", token, refreshToken).Error
}
