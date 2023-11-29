package manager

import "task-5-pbi-btpns-achmad-dinofaldi-firmansyah/usecase"

type UseCaseManager interface {
	UserUseCase() usecase.UserUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repoManager.UserRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{repoManager}
}
