package repository_test

import (
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/model"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PhotoRepositoryTestSuite struct {
	suite.Suite
	db         *gorm.DB
	repository repository.PhotoRepository
}

func (s *PhotoRepositoryTestSuite) SetupTest() {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=rakamin_btpns_test port=5432 sslmode=disable TimeZone=Asia/Jakarta"), &gorm.Config{})
	assert.NoError(s.T(), err)
	s.db = db
	s.repository = repository.NewPhotoRepository(db)
	s.db.Migrator().DropTable(&model.Photo{})
	s.db.AutoMigrate(&model.Photo{})
}

func (s *PhotoRepositoryTestSuite) TearDownTest() {
	s.db.Migrator().DropTable(&model.Photo{})
}

var payloadPhoto = []model.Photo{
	{
		Id:      "1",
		Title:   "title1",
		Caption: "caption1",
		Url:     "url1",
		UserId:  "1",
	},
	{
		Id:      "2",
		Title:   "title2",
		Caption: "caption2",
		Url:     "url2",
		UserId:  "2",
	},
}

func (s *PhotoRepositoryTestSuite) TestCreateSuccess() {
	err := s.repository.Create(payloadPhoto[0])
	assert.NoError(s.T(), err)
}

func (s *PhotoRepositoryTestSuite) TestCreateFailed() {
	err := s.repository.Create(payloadPhoto[0])
	assert.NoError(s.T(), err)

	err = s.repository.Create(payloadPhoto[0])
	assert.Error(s.T(), err)
}

func (s *PhotoRepositoryTestSuite) TestFindByIdSuccess() {
	err := s.repository.Create(payloadPhoto[0])
	assert.NoError(s.T(), err)

	photo, err := s.repository.FindById(payloadPhoto[0].Id)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), payloadPhoto[0].Id, photo.Id)
}

func (s *PhotoRepositoryTestSuite) TestFindByIdFailed() {
	photo, err := s.repository.FindById(payloadPhoto[0].Id)
	assert.Error(s.T(), err)
	assert.Equal(s.T(), model.PhotoResponse{}, photo)
}

func (s *PhotoRepositoryTestSuite) TestGetSuccess() {
	s.db.AutoMigrate(&model.User{})
	s.db.Create(&model.User{
		Id:       "1",
		Username: "username1",
		Email:    "email1",
	})
	s.db.Create(&model.User{
		Id:       "2",
		Username: "username2",
		Email:    "email2",
	})

	err := s.repository.Create(payloadPhoto[0])
	assert.NoError(s.T(), err)

	err = s.repository.Create(payloadPhoto[1])
	assert.NoError(s.T(), err)

	photos, err := s.repository.Get()
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 2, len(photos))
}

func (s *PhotoRepositoryTestSuite) TestGetFailed() {
	s.db.Migrator().DropTable(&model.Photo{})
	photos, err := s.repository.Get()
	assert.Error(s.T(), err)
	assert.Equal(s.T(), 0, len(photos))
}

func (s *PhotoRepositoryTestSuite) TestUpdateSuccess() {
	s.db.AutoMigrate(&model.User{})
	s.db.Create(&model.User{
		Id:       "1",
		Username: "username1",
		Email:    "email1",
	})

	err := s.repository.Create(payloadPhoto[0])
	assert.NoError(s.T(), err)

	err = s.repository.Update(payloadPhoto[0])
	assert.NoError(s.T(), err)
}

func (s *PhotoRepositoryTestSuite) TestUpdateFailed() {
	s.db.AutoMigrate(&model.User{})
	s.db.Create(&model.User{
		Id:       "1",
		Username: "username1",
		Email:    "email1",
	})

	err := s.repository.Create(payloadPhoto[0])
	assert.NoError(s.T(), err)

	s.db.Migrator().DropTable(&model.Photo{})

	err = s.repository.Update(payloadPhoto[0])
	assert.Error(s.T(), err)
}

func (s *PhotoRepositoryTestSuite) TestDeleteSuccess() {
	s.db.AutoMigrate(&model.User{})
	s.db.Create(&model.User{
		Id:       "1",
		Username: "username1",
		Email:    "email1",
	})

	err := s.repository.Create(payloadPhoto[0])
	assert.NoError(s.T(), err)

	err = s.repository.Delete(payloadPhoto[0])
	assert.NoError(s.T(), err)
}

func (s *PhotoRepositoryTestSuite) TestDeleteFailed() {
	s.db.Migrator().DropTable(&model.Photo{})

	err := s.repository.Delete(payloadPhoto[0])
	assert.Error(s.T(), err)
}

func TestPhotoRepositorySuite(t *testing.T) {
	suite.Run(t, new(PhotoRepositoryTestSuite))
}
