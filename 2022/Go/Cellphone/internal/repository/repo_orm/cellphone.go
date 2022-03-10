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

func (self *CellphoneRepository) ServeSingleFromProvider(providerId int) (*entity.Cellphone, error) {
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

func (self *CellphoneRepository) InsertSingle(cellphone *entity.Cellphone) error {
	tx := self.Db.Model(&entity.Cellphone{}).Begin()
	defer func() {
		if tx.Error == nil {
			tx.Commit()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	tx.Create(cellphone)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (self *CellphoneRepository) GetAllByProviderName(providerName string) ([]*entity.Cellphone, error) {

	var provider entity.Provider

	res := self.Db.Model(&entity.Provider{}).Where(&entity.Provider{Name: providerName}).First(&provider)

	if res.Error != nil {
		return nil, res.Error
	}

	rows, err := self.Db.Model(&entity.Cellphone{}).Where(&entity.Cellphone{ProviderId: provider.Id}).Rows()

	if res.Error != nil {
		return nil, res.Error
	}

	var result []*entity.Cellphone

	for rows.Next() {
		var cellphone entity.Cellphone

		err = self.Db.ScanRows(rows, &cellphone)

		if err != nil {
			return nil, err
		}

		result = append(result, &cellphone)
	}

	return result, nil
}

func (self *CellphoneRepository) DeleteAllFromProvider(providerId int) (int, error) {
	// TODO
	return 0, nil
}
