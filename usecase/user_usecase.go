package usecase

import (
	"errors"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/helper/common"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/helper/security"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/model"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	Register(user model.User) error
	Login(userCredential model.UserCredential) (string, error)
	FindById(id string) (model.UserResponse, error)
	Update(user model.User) error
	Delete(user model.User) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func (u *userUseCase) Register(user model.User) error {
	_, err := u.userRepository.FindByEmail(user.Email)
	if err == nil {
		return errors.New("email already exist")
	}

	bytesPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Id = common.GenerateUUID()
	user.Password = string(bytesPassword)

	return u.userRepository.Register(user)
}

func (u *userUseCase) Login(userCredential model.UserCredential) (string, error) {
	user, err := u.userRepository.FindByEmail(userCredential.Email)
	if err != nil {
		return "", errors.New("email or password is wrong")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userCredential.Password)); err != nil {
		return "", errors.New("email or password is wrong")
	}

	token, err := security.CreateAccessToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *userUseCase) FindById(id string) (model.UserResponse, error) {
	return u.userRepository.FindById(id)
}

func (u *userUseCase) Update(user model.User) error {
	bytesPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(bytesPassword)

	return u.userRepository.Update(user)
}

func (u *userUseCase) Delete(user model.User) error {
	return u.userRepository.Delete(user)
}

func NewUserUseCase(userRepository repository.UserRepository) *userUseCase {
	return &userUseCase{userRepository}
}
