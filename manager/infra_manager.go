package manager

import (
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type InfraManager interface {
	Connection() *gorm.DB
	GetConfig() *config.Config
}

type infraManager struct {
	db  *gorm.DB
	cfg *config.Config
}

func (i *infraManager) initDb() error {
	var dbConf = i.cfg.DBConfig

	db, err := gorm.Open(postgres.Open(dbConf.Url), &gorm.Config{})
	if err != nil {
		return err
	}

	i.db = db
	return nil
}

func (i *infraManager) Connection() *gorm.DB {
	return i.db
}

func (i *infraManager) GetConfig() *config.Config {
	return i.cfg
}

func NewInfraManager(configParam *config.Config) (InfraManager, error) {
	infra := &infraManager{
		cfg: configParam,
	}

	err := infra.initDb()
	if err != nil {
		return nil, err
	}

	return infra, nil
}
