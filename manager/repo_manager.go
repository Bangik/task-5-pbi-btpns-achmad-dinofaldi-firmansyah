package manager

import "task-5-pbi-btpns-achmad-dinofaldi-firmansyah/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
	PhotoRepo() repository.PhotoRepository
}

type repoManager struct {
	infraManager InfraManager
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infraManager.Connection())
}

func (r *repoManager) PhotoRepo() repository.PhotoRepository {
	return repository.NewPhotoRepository(r.infraManager.Connection())
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{infraManager}
}
