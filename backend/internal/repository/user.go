package repository

import (
	"github.com/aternity/zense/internal/entity/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) (*domain.User, error)
	FindAll() ([]domain.User, error)
	FindByID(id uint) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(user *domain.User) (*domain.User, error) {
	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) FindAll() ([]domain.User, error) {
	var users []domain.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return users, nil
}

func (u *userRepository) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) Update(user *domain.User) (*domain.User, error) {
	if err := u.db.Updates(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) Delete(id uint) error {
	return u.db.Delete(&domain.User{}, id).Error
}
