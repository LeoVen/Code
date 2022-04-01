package repo_orm

import (
	"cellphone/internal/entity"
	"errors"

	"gorm.io/gorm"
)

type ProviderRepository struct {
	Db *gorm.DB
}

func (self *ProviderRepository) GetById(id int) (interface{}, error) {
	var result entity.Provider
	tx := self.Db.Model(&entity.Provider{}).First(&result, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &result, nil
}

func (self *ProviderRepository) GetByName(name string) (*entity.Provider, error) {
	var result entity.Provider
	self.Db.Model(&entity.Provider{}).First(&result, "NAME = ?", name)

	if self.Db.Error != nil {
		return nil, self.Db.Error
	}

	return &result, nil
}

func (self *ProviderRepository) GetCount(id int) (*int, error) {
	var result entity.Provider
	tx := self.Db.Model(&entity.Provider{}).First(&result, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	v := int(result.Total)
	return &v, nil
}

func (self *ProviderRepository) Insert(provider *entity.Provider) error {
	// TODO
	return errors.New("Unimplemented")
}

func (self *ProviderRepository) Delete(id int) error {
	// TODO
	return errors.New("Unimplemented")
}

func (self *ProviderRepository) Update(provider *entity.Provider) error {
	// TODO
	return errors.New("Unimplemented")
}
