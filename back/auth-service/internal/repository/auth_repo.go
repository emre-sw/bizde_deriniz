package repository

import (
	"auth/internal/domain"
	"context"
	"errors"

	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{
		db: db,
	}
}

func (d *AuthRepositoryImpl) GetAllUsers() ([]domain.Auth, error){
	var auths []domain.Auth
	if err := d.db.Find(&auths); err.Error != nil {
		return nil, err.Error
	}

	return auths, nil
}


func (d *AuthRepositoryImpl) CreateAuth(ctx context.Context, auth *domain.Auth) error {
	return d.db.WithContext(ctx).Create(auth).Error
}

func (d *AuthRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*domain.Auth, error) {
	var auth domain.Auth
	err := d.db.WithContext(ctx).Where("email = ?", email).First(&auth).Error
	 if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to db")
	 }
	return &auth, nil
	
}

func (r *AuthRepositoryImpl) UpdatePassword(ctx context.Context, email, password string) error {
	var user domain.Auth
	
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return errors.New("failed to db")
	}

	result := r.db.WithContext(ctx).Model(&user).Update("password", password)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
