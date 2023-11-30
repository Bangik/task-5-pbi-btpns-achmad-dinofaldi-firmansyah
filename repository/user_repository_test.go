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

type UserRepositoryTestSuite struct {
	suite.Suite
	db         *gorm.DB
	repository repository.UserRepository
}

func (s *UserRepositoryTestSuite) SetupTest() {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=rakamin_btpns_test port=5432 sslmode=disable TimeZone=Asia/Jakarta"), &gorm.Config{})
	assert.NoError(s.T(), err)
	s.db = db
	s.repository = repository.NewUserRepository(db)
	s.db.Migrator().DropTable(&model.User{})
	s.db.AutoMigrate(&model.User{})
}

func (s *UserRepositoryTestSuite) TearDownTest() {
	s.db.Migrator().DropTable(&model.User{})
}

var payload = []model.User{
	{
		Id:       "1",
		Username: "user1",
		Email:    "user1@gmail.com",
		Password: "password",
	},
	{
		Id:       "2",
		Username: "user2",
		Email:    "user2@gmail.com",
		Password: "password",
	},
}

func (s *UserRepositoryTestSuite) TestRegisterSuccess() {
	err := s.repository.Register(payload[0])
	assert.NoError(s.T(), err)
}

func (s *UserRepositoryTestSuite) TestRegisterFailed() {
	err := s.repository.Register(payload[0])
	assert.NoError(s.T(), err)

	err = s.repository.Register(payload[0])
	assert.Error(s.T(), err)
}

func (s *UserRepositoryTestSuite) TestFindByEmailSuccess() {
	err := s.repository.Register(payload[0])
	assert.NoError(s.T(), err)

	user, err := s.repository.FindByEmail(payload[0].Email)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), payload[0].Email, user.Email)
}

func (s *UserRepositoryTestSuite) TestFindByEmailFailed() {
	user, err := s.repository.FindByEmail(payload[0].Email)
	assert.Error(s.T(), err)
	assert.Empty(s.T(), user)
}

func (s *UserRepositoryTestSuite) TestFindByIdSuccess() {
	err := s.repository.Register(payload[0])
	assert.NoError(s.T(), err)

	user, err := s.repository.FindById(payload[0].Id)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), payload[0].Id, user.Id)
}

func (s *UserRepositoryTestSuite) TestFindByIdFailed() {
	user, err := s.repository.FindById(payload[0].Id)
	assert.Error(s.T(), err)
	assert.Empty(s.T(), user)
}

func (s *UserRepositoryTestSuite) TestUpdateSuccess() {
	err := s.repository.Register(payload[0])
	assert.NoError(s.T(), err)

	err = s.repository.Update(payload[0])
	assert.NoError(s.T(), err)
}

func (s *UserRepositoryTestSuite) TestUpdateFailed() {
	err := s.repository.Register(payload[0])
	assert.NoError(s.T(), err)

	err = s.repository.Register(payload[1])
	assert.NoError(s.T(), err)

	payload[0].Email = payload[1].Email
	err = s.repository.Update(payload[0])
	assert.Error(s.T(), err)
}

func (s *UserRepositoryTestSuite) TestDeleteSuccess() {
	err := s.repository.Register(payload[0])
	assert.NoError(s.T(), err)

	err = s.repository.Delete(payload[0])
	assert.NoError(s.T(), err)
}

func (s *UserRepositoryTestSuite) TestDeleteFailed() {
	s.db.Migrator().DropTable(&model.User{})

	err := s.repository.Delete(payload[0])
	assert.Error(s.T(), err)
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
