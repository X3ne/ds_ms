package repositories

import (
	"context"

	"github.com/X3ne/ds_ms/users_service/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	var user models.User
	if err := ur.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil

}

func (ur *UserRepository) UpdateUser(ctx context.Context, user *models.User) error {
	if err := ur.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) DeleteUser(ctx context.Context, userID string) error {
	if err := ur.db.Delete(&models.User{}, userID).Error; err != nil {
		return err
	}
	return nil
}
