package manager

import "task-5-pbi-btpns-achmad-dinofaldi-firmansyah/usecase"

type UseCaseManager interface {
	UserUseCase() usecase.UserUseCase
	PhotoUseCase() usecase.PhotoUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repoManager.UserRepo())
}

func (u *useCaseManager) PhotoUseCase() usecase.PhotoUseCase {
	return usecase.NewPhotoUseCase(u.repoManager.PhotoRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{repoManager}
}
