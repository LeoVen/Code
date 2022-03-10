package repo_sql

import (
	"cellphone/internal/entity"
	"database/sql"
)

type ProviderRepository struct {
	Db *sql.DB
}

func (self *ProviderRepository) GetById(id int) (interface{}, error) {
	query := "SELECT * FROM PROVIDER WHERE ID = ?;"

	row := self.Db.QueryRow(query)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var provider entity.Provider
	if err := row.Scan(&provider.Id, &provider.Name); err != nil {
		return nil, err
	}

	return provider, nil
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

func (self *ProviderRepository) InsertSingle(provider *entity.Provider) error {
	query := "INSERT INTO PROVIDER (NAME) VALUES (?);"

	tx, err := self.Db.Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec(query, provider.Name)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}
