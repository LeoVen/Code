package repo_orm

import (
	"cellphone/internal/entity"

	"gorm.io/gorm"
)

type CellphoneRepository struct {
	Db *gorm.DB
}

func (self *CellphoneRepository) GetById(id int) (interface{}, error) {
	var result entity.Cellphone
	tx := self.Db.Model(&entity.Cellphone{}).First(&result, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func (self *CellphoneRepository) FetchSingle(providerId int) (*entity.Cellphone, error) {
	var result entity.Cellphone

	self.Db.Model(&entity.Cellphone{}).Where(&entity.Cellphone{
		ProviderId: providerId,
	}).First(&result)

	if self.Db.Error != nil {
		return nil, self.Db.Error
	}

	tx := self.Db.Begin()
	defer func() {
		if tx.Error == nil {
			tx.Commit()
		}
	}()

	if tx.Error != nil {
		return nil, tx.Error
	}

	self.Db.Model(&entity.Cellphone{}).Delete(&entity.Cellphone{}, 0)

	if self.Db.Error != nil {
		return nil, self.Db.Error
	}

	return &result, nil
}

func (self *CellphoneRepository) BulkInsert(providerId int, entities []entity.Cellphone) error {
	return nil
}
