package repository

import (
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user model.User) error
	FindByEmail(email string) (model.User, error)
	FindById(id string) (model.UserResponse, error)
	Update(user model.User) error
	Delete(user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) Register(user model.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *userRepository) FindById(id string) (model.UserResponse, error) {
	var user model.UserResponse

	if err := r.db.Table("users").Select("id, username, email, created_at, updated_at").Where("id = ?", id).First(&user).Error; err != nil {
		return model.UserResponse{}, err
	}

	return user, nil
}

func (r *userRepository) Update(user model.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Delete(user model.User) error {
	if err := r.db.Delete(user).Error; err != nil {
		return err
	}

	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
