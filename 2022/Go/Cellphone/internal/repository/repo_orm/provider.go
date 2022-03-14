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

func (self *ProviderRepository) GetCount(id int) (int, error) {
	return 0, nil
}

func (self *ProviderRepository) Insert(provider *entity.Provider) error {
	return nil
}

func (self *ProviderRepository) Delete(id int) error {
	// TODO
	return nil
}

func (self *ProviderRepository) Update(provider *entity.Provider) error {
	// TODO
	return nil
}
