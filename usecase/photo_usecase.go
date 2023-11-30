package usecase

import (
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/helper/common"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/model"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/repository"
)

type PhotoUseCase interface {
	Create(photo model.Photo) error
	FindById(id string) (model.PhotoResponse, error)
	Get() ([]model.PhotoResponse, error)
	Update(photo model.Photo) error
	Delete(photo model.Photo) error
}

type photoUseCase struct {
	photoRepository repository.PhotoRepository
}

func (u *photoUseCase) Create(photo model.Photo) error {
	photo.Id = common.GenerateUUID()

	return u.photoRepository.Create(photo)
}

func (u *photoUseCase) FindById(id string) (model.PhotoResponse, error) {
	return u.photoRepository.FindById(id)
}

func (u *photoUseCase) Get() ([]model.PhotoResponse, error) {
	return u.photoRepository.Get()
}

func (u *photoUseCase) Update(photo model.Photo) error {
	return u.photoRepository.Update(photo)
}

func (u *photoUseCase) Delete(photo model.Photo) error {
	return u.photoRepository.Delete(photo)
}

func NewPhotoUseCase(photoRepository repository.PhotoRepository) PhotoUseCase {
	return &photoUseCase{photoRepository}
}
