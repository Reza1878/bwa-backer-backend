package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
}

type ServiceImpl struct {
	repository Repository
}

func NewService(repository Repository) *ServiceImpl {
	return &ServiceImpl{
		repository: repository,
	}
}

func (service *ServiceImpl) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	user.Password = string(passwordHash)
	user.Role = "user"

	user, err = service.repository.Save(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *ServiceImpl) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.Name == "" {
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *ServiceImpl) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	user, err := s.repository.FindByEmail(input.Email)

	if user.Id != 0 {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
