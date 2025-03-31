package service

import (
	"errors"
	"fmt"
	"go-myobokucomerce-app/internal/domain"
	"go-myobokucomerce-app/internal/dto"
	"go-myobokucomerce-app/internal/repository"
	"log"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s UserService) Signup(input dto.UserSignup) (string, error) {

	//Create User Functionality
	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
	})

	//Generate Token
	log.Println(user)

	userInfo := fmt.Sprintf("%v,%v,%v", user.ID, user.Password, user.Phone)

	return userInfo, err
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {

	user, err := s.Repo.FindUser(email)

	return &user, err
}

func (s UserService) Login(email string, password string) (string, error) {

	user, err := s.findUserByEmail(email)

	if err != nil {
		return "", errors.New("user does not exist with the provided email id")

	}

	//compare password and generate token

	return user.Email, nil
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {

	return 0, nil
}

func (s UserService) VerifyCode(id uint, code int) error {

	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {

	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {

	return nil, nil
}

func (s UserService) UpdateProfile(id uint, input any) error {

	return nil
}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {

	return "", nil
}

func (s UserService) FindCart(id uint) ([]interface{}, error) {

	return nil, nil
}

func (s UserService) CreateCart(input any, u domain.User) ([]interface{}, error) {

	return nil, nil
}

func (s UserService) CreateOrder(u domain.User) (int, error) {

	return 0, nil
}

func (s UserService) GetOrders(u domain.User) ([]interface{}, error) {

	return nil, nil
}

func (s UserService) GetOrderById(id uint, uId uint) ([]interface{}, error) {

	return nil, nil
}
