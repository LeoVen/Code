package repo_orm

import (
	"cellphone/internal/entity"

	"gorm.io/gorm"
)

type ProviderRepository struct {
	Db *gorm.DB
}

func (self *ProviderRepository) GetById(id int) (interface{}, error) {
	// TODO
	return nil, nil
}

func (self *ProviderRepository) GetByName(name string) (*entity.Provider, error) {
	// TODO
	return nil, nil
}

func (self *ProviderRepository) InsertSingle(provider *entity.Provider) error {
	// TODO
	return nil
}
