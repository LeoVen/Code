package repo_sql

import (
	"cellphone/internal/entity"
	"database/sql"
	"errors"
)

type ProviderRepository struct {
	Db *sql.DB
}

func (self *ProviderRepository) GetById(id int) (interface{}, error) {
	query := "SELECT * FROM PROVIDER WHERE ID = ?"

	row := self.Db.QueryRow(query, id)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var provider entity.Provider
	if err := row.Scan(&provider.Id, &provider.Name, &provider.Total); err != nil {
		return nil, err
	}

	return &provider, nil
}

func (self *ProviderRepository) GetByName(name string) (*entity.Provider, error) {
	query := "SELECT * FROM PROVIDER WHERE NAME = ?;"

	row := self.Db.QueryRow(query, name)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var entity entity.Provider
	err := row.Scan(&entity.Id, &entity.Name, &entity.Total)

	return &entity, err
}

func (self *ProviderRepository) GetCount(id int) (*int, error) {
	query := "SELECT Total from PROVIDER WHERE id = ?;"

	row := self.Db.QueryRow(query, id)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var result int
	err := row.Scan(&result)

	return &result, err
}

func (self *ProviderRepository) Insert(provider *entity.Provider) error {
	query := "INSERT INTO PROVIDER (NAME, TOTAL) VALUES (?, ?);"

	tx, err := self.Db.Begin()

	if err != nil {
		return err
	}

	res, err := tx.Exec(query, provider.Name, provider.Total)

	if err != nil {
		return err
	}

	lines, err := res.RowsAffected()

	if lines == 0 {
		// TODO
	}

	return tx.Commit()
}

func (self *ProviderRepository) Delete(id int) error {
	// TODO
	return errors.New("Unimplemented")
}

func (self *ProviderRepository) Update(provider *entity.Provider) error {
	// TODO
	return errors.New("Unimplemented")
}
