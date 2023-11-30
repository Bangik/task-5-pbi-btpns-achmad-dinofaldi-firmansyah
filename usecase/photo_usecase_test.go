package usecase_test

import (
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/model"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockPhotoRepository struct {
	mock.Mock
}

func (s *mockPhotoRepository) Create(photo model.Photo) error {
	args := s.Called(photo)
	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (s *mockPhotoRepository) Delete(photo model.Photo) error {
	args := s.Called(photo)
	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (s *mockPhotoRepository) FindById(id string) (model.PhotoResponse, error) {
	args := s.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(model.PhotoResponse), args.Error(1)
	}

	return model.PhotoResponse{}, nil
}

func (s *mockPhotoRepository) Get() ([]model.PhotoResponse, error) {
	args := s.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]model.PhotoResponse), args.Error(1)
	}

	return []model.PhotoResponse{}, nil
}

func (s *mockPhotoRepository) Update(photo model.Photo) error {
	args := s.Called(photo)
	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

type PhotoUseCaseTestSuite struct {
	suite.Suite
	mockRepo *mockPhotoRepository
	useCase  usecase.PhotoUseCase
}

func (s *PhotoUseCaseTestSuite) SetupTest() {
	s.mockRepo = new(mockPhotoRepository)
	s.useCase = usecase.NewPhotoUseCase(s.mockRepo)
}

func (s *PhotoUseCaseTestSuite) TestCreate() {
	s.mockRepo.On("Create", mock.Anything).Return(nil)

	err := s.useCase.Create(model.Photo{})
	assert.NoError(s.T(), err)
}

func (s *PhotoUseCaseTestSuite) TestDelete() {
	s.mockRepo.On("Delete", mock.Anything).Return(nil)

	err := s.useCase.Delete(model.Photo{})
	assert.NoError(s.T(), err)
}

func (s *PhotoUseCaseTestSuite) TestFindById() {
	s.mockRepo.On("FindById", mock.Anything).Return(model.PhotoResponse{}, nil)

	_, err := s.useCase.FindById("")
	assert.NoError(s.T(), err)
}

func (s *PhotoUseCaseTestSuite) TestGet() {
	s.mockRepo.On("Get").Return([]model.PhotoResponse{}, nil)

	_, err := s.useCase.Get()
	assert.NoError(s.T(), err)
}

func (s *PhotoUseCaseTestSuite) TestUpdate() {
	s.mockRepo.On("Update", mock.Anything).Return(nil)

	err := s.useCase.Update(model.Photo{})
	assert.NoError(s.T(), err)
}

func TestPhotoUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(PhotoUseCaseTestSuite))
}
