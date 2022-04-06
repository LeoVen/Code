package repo_orm

import (
	pb "cellphone/protos/go"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type CellphoneRepository struct {
	Db *gorm.DB
}

func (self *CellphoneRepository) GetById(id int) (*pb.Cellphone, error) {
	var result pb.Cellphone
	tx := self.Db.Model(&pb.Cellphone{}).First(&result, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &result, nil
}

func (self *CellphoneRepository) FetchSingle(providerId int) (*pb.Cellphone, error) {
	var result pb.Cellphone

	self.Db.Model(&pb.Cellphone{}).Where(&pb.Cellphone{
		ProviderId: int64(providerId),
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

	self.Db.Model(&pb.Cellphone{}).Delete(&pb.Cellphone{}, 0)

	if self.Db.Error != nil {
		return nil, self.Db.Error
	}

	return &result, nil
}

func (self *CellphoneRepository) BulkInsert(providerId int, entities []*pb.Cellphone) (int, error) {
	if len(entities) == 0 {
		return 0, errors.New("bulkInsert received empty array")
	}

	query := "INSERT INTO CELLPHONE (PROVIDER_ID, NUMBER) VALUES"
	sb := strings.Builder{}

	sb.WriteString(query)

	values := []interface{}{}

	for _, entity := range entities {
		sb.WriteString(" (?, ?),")
		if entity.ProviderId == 0 {
			entity.ProviderId = int64(providerId)
		}
		values = append(values, entity.ProviderId, entity.Number)
	}

	fullQuery := sb.String()
	fullQuery = strings.TrimSuffix(fullQuery, ",") // Remove last ','

	tx := self.Db.Exec(fullQuery, values...)

	if tx.Error != nil {
		return 0, tx.Error
	}

	if tx.RowsAffected <= 0 {
		return 0, errors.New("bulkInsert failed with no rows affected")
	}

	return int(tx.RowsAffected), nil
}
