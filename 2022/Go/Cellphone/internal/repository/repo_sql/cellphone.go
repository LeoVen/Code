package repo_sql

import (
	pb "cellphone/protos/go"
	"database/sql"
	"errors"
)

type CellphoneRepository struct {
	Db *sql.DB
}

func (self *CellphoneRepository) GetById(id int) (*pb.Cellphone, error) {
	query := "SELECT * FROM CELLPHONE WHERE ID = ?;"

	row := self.Db.QueryRow(query, id)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var phone pb.Cellphone
	if err := row.Scan(&phone.Id, &phone.ProviderId, &phone.Number); err != nil {
		return nil, err
	}

	return &phone, nil
}

// Retrieves a semi-random phone number and deletes it from the database
func (self *CellphoneRepository) FetchSingle(providerId int) (*pb.Cellphone, error) {
	query := "SELECT * FROM CELLPHONE WHERE PROVIDER_ID = ? LIMIT 1"
	queryDelete := "DELETE FROM CELLPHONE WHERE ID = ?;"

	row := self.Db.QueryRow(query, providerId)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var entity pb.Cellphone

	err := row.Scan(&entity.Id, &entity.ProviderId, &entity.Number)

	if err != nil {
		return nil, err
	}

	tx, err := self.Db.Begin()

	_, err = tx.Exec(queryDelete, entity.Id)

	if err != nil {
		return nil, err
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (self *CellphoneRepository) BulkInsert(providerId int, entities []*pb.Cellphone) (int, error) {
	return 0, errors.New("Unimplemented")
}
