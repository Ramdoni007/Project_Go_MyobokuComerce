package repository

import (
	"errors"
	"go-myobokucomerce-app/internal/domain"
	"gorm.io/gorm"
	"log"
)

type UserRepository interface {
	CreateUser(usr domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserById(id uint) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}

}

func (r userRepository) CreateUser(usr domain.User) (domain.User, error) {

	err := r.db.Create(&usr).Error

	if err != nil {
		log.Printf("create user error %v", err)
		return domain.User{}, errors.New("failed to create user")
	}

	return usr, nil

}

func (r userRepository) FindUser(email string) (domain.User, error) {

	return domain.User{}, nil

}

func (r userRepository) FindUserById(id uint) (domain.User, error) {

	return domain.User{}, nil

}

func (r userRepository) UpdateUser(id uint, u domain.User) (domain.User, error) {

	return domain.User{}, nil

}
