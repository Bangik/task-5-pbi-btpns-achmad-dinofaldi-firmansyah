package repository

import (
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/model"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(photo model.Photo) error
	FindById(id string) (model.PhotoResponse, error)
	Get() ([]model.PhotoResponse, error)
	Update(photo model.Photo) error
	Delete(photo model.Photo) error
}

type photoRepository struct {
	db *gorm.DB
}

func (r *photoRepository) Create(photo model.Photo) error {
	if err := r.db.Create(&photo).Error; err != nil {
		return err
	}

	return nil
}

func (r *photoRepository) FindById(id string) (model.PhotoResponse, error) {
	var photo model.PhotoResponse

	if err := r.db.Table("photos").Select("id, title, caption, url, user_id, created_at, updated_at").Where("id = ?", id).First(&photo).Error; err != nil {
		return model.PhotoResponse{}, err
	}

	return photo, nil
}

func (r *photoRepository) Get() ([]model.PhotoResponse, error) {
	var photo []model.PhotoResponse

	// get all data and join with users table
	err := r.db.Table("photos").
		Select("photos.id, photos.title, photos.caption, photos.url, photos.user_id, photos.created_at, photos.updated_at, users.username, users.email").
		Joins("JOIN users ON photos.user_id = users.id").
		Scan(&photo).
		Error
	if err != nil {
		return []model.PhotoResponse{}, err
	}

	return photo, nil
}

func (r *photoRepository) Update(photo model.Photo) error {
	if err := r.db.Save(photo).Error; err != nil {
		return err
	}

	return nil
}

func (r *photoRepository) Delete(photo model.Photo) error {
	if err := r.db.Delete(photo).Error; err != nil {
		return err
	}

	return nil
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoRepository{
		db: db,
	}
}
