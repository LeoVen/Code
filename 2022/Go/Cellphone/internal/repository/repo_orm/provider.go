package repo_orm

import (
	pb "cellphone/protos/go"
	"errors"

	"gorm.io/gorm"
)

type ProviderRepository struct {
	Db *gorm.DB
}

func (self *ProviderRepository) GetById(id int) (*pb.Provider, error) {
	var result pb.Provider
	tx := self.Db.Model(&pb.Provider{}).First(&result, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &result, nil
}

func (self *ProviderRepository) GetByName(name string) (*pb.Provider, error) {
	var result pb.Provider
	self.Db.Model(&pb.Provider{}).First(&result, "NAME = ?", name)

	if self.Db.Error != nil {
		return nil, self.Db.Error
	}

	return &result, nil
}

func (self *ProviderRepository) GetCount(id int) (*int, error) {
	var result pb.Provider
	tx := self.Db.Model(&pb.Provider{}).First(&result, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	v := int(result.Total)
	return &v, nil
}

func (self *ProviderRepository) Insert(provider *pb.Provider) error {
	// TODO
	return errors.New("Unimplemented")
}

func (self *ProviderRepository) Delete(id int) error {
	// TODO
	return errors.New("Unimplemented")
}

func (self *ProviderRepository) Update(provider *pb.Provider) error {
	// TODO
	return errors.New("Unimplemented")
}
