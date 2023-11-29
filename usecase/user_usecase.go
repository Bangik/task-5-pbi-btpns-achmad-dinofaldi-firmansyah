package usecase

import (
	"errors"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/helper/common"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/model"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	Register(user model.User) (model.User, error)
	Login(userCredential model.UserCredential) (model.User, error)
	Update(user model.User) (model.User, error)
	Delete(user model.User) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func (u *userUseCase) Register(user model.User) (model.User, error) {
	bytesPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}

	user.Id = common.GenerateUUID()
	user.Password = string(bytesPassword)

	return u.userRepository.Register(user)
}

func (u *userUseCase) Login(userCredential model.UserCredential) (model.User, error) {
	user, err := u.userRepository.FindByEmail(userCredential.Email)
	if err != nil {
		return model.User{}, errors.New("email or password is wrong")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userCredential.Password)); err != nil {
		return model.User{}, errors.New("email or password is wrong")
	}

	return user, nil
}

func (u *userUseCase) Update(user model.User) (model.User, error) {
	return u.userRepository.Update(user)
}

func (u *userUseCase) Delete(user model.User) error {
	return u.userRepository.Delete(user)
}

func NewUserUseCase(userRepository repository.UserRepository) *userUseCase {
	return &userUseCase{userRepository}
}
