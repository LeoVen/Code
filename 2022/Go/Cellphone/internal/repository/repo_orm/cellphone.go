package repo_orm

import (
	"cellphone/internal/entity"
	"errors"
	"strings"

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

	return &result, nil
}

func (self *CellphoneRepository) FetchSingle(providerId int) (*entity.Cellphone, error) {
	var result entity.Cellphone

	self.Db.Model(&entity.Cellphone{}).Where(&entity.Cellphone{
		ProviderId: int32(providerId),
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

func (self *CellphoneRepository) BulkInsert(providerId int, entities []*entity.Cellphone) error {
	if len(entities) == 0 {
		return errors.New("bulkInsert received empty array")
	}

	query := "INSERT INTO CELLPHONE (PROVIDER_ID, NUMBER) VALUES"
	sb := strings.Builder{}

	sb.WriteString(query)

	values := []interface{}{}

	for _, entity := range entities {
		sb.WriteString(" (?, ?),")
		if entity.ProviderId == 0 {
			entity.ProviderId = int32(providerId)
		}
		values = append(values, entity.ProviderId, entity.Number)
	}

	fullQuery := sb.String()
	fullQuery = strings.TrimSuffix(fullQuery, ",") // Remove last ','

	tx := self.Db.Exec(fullQuery, values...)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected <= 0 {
		return errors.New("bulkInsert failed with no rows affected")
	}

	return nil
}
