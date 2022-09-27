package user

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type serviceImpl struct {
	repository Repository
}

func NewService(repository Repository) *serviceImpl {
	return &serviceImpl{
		repository: repository,
	}
}

func (service *serviceImpl) RegisterUser(input RegisterUserInput) (User, error) {
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

func (s *serviceImpl) UpdateUser(input UpdateUserInput) (User, error) {
	user, err := s.repository.FindById(input.User.Id)
	if err != nil {
		return input.User, err
	}

	userByEmail, err := s.repository.FindByEmail(input.Email)
	if err == nil {
		if userByEmail.Id != input.User.Id {
			return input.User, errors.New("email has been taken")
		}
	}

	user.Name = input.Name
	user.Occupation = input.Occupation
	user.Email = input.Email

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return input.User, err
	}

	return updatedUser, nil
}

func (s *serviceImpl) UpdatePassword(input UpdateUserPasswordInput) error {
	fmt.Println(input)
	err := bcrypt.CompareHashAndPassword([]byte(input.User.Password), []byte(input.OldPassword))

	if err != nil {
		return errors.New("old password is invalid")
	}

	if input.NewPassword != input.ConfirmationPassword {
		return errors.New("new password didnt match")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.MinCost)
	if err != nil {
		return err
	}

	currentUser, _ := s.repository.FindById(input.User.Id)
	currentUser.Password = string(passwordHash)
	_, err = s.repository.Update(currentUser)

	return err
}

func (s *serviceImpl) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.Name == "" {
		return user, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("credential is invalid")
	}

	return user, nil
}

func (s *serviceImpl) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	user, err := s.repository.FindByEmail(input.Email)

	if user.Id != 0 {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *serviceImpl) SaveAvatar(ID int, fileLocation string) (User, error) {
	user, err := s.repository.FindById(ID)
	if err != nil {
		return user, err
	}

	user.Avatar = fileLocation

	user, err = s.repository.Update(user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *serviceImpl) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindById(ID)
	if err != nil {
		return user, err
	}

	if user.Id == 0 {
		return user, errors.New("user not found")
	}

	return user, nil
}
