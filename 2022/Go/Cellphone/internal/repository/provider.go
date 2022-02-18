package repository

import (
	"cellphone/internal/entity"
	"database/sql"

	mysqlGorm "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	getByIdProvider   = "SELECT * FROM PROVIDER WHERE ID = ?;"
	getByNameProvider = "SELECT * FROM PROVIDER WHERE NAME = ?;"

	insertSingleProvider = "INSERT INTO PROVIDER (NAME) VALUES (?);"
)

type ProviderRepositoryMock struct {
	db *sql.DB
}

type ProviderRepositorySql struct {
	db *sql.DB
}

type ProviderRepositoryGorm struct {
	db *gorm.DB
}

func NewProviderRepository(repoType int, conn *sql.DB) (ProviderRepository, error) {
	switch repoType {
	case REPO_SQL:
		return newProviderRepositorySql(conn)
	case REPO_MOCK:
		return newProviderRepositoryMock(conn)
	case REPO_GORM:
		return newProviderRepositoryGorm(conn)
	}
	return newProviderRepositorySql(conn)
}

func newProviderRepositoryMock(conn *sql.DB) (ProviderRepository, error) {
	return &ProviderRepositoryMock{conn}, nil
}

func newProviderRepositorySql(conn *sql.DB) (ProviderRepository, error) {
	return &ProviderRepositorySql{conn}, nil
}

func newProviderRepositoryGorm(conn *sql.DB) (ProviderRepository, error) {
	gormDB, err := gorm.Open(mysqlGorm.New(mysqlGorm.Config{
		Conn: conn,
	}))

	if err != nil {
		return nil, err
	}

	return &ProviderRepositoryGorm{gormDB}, nil
}

func (self *ProviderRepositorySql) GetById(id int) (interface{}, error) {
	row := self.db.QueryRow(getByIdProvider)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var provider entity.Provider
	if err := row.Scan(&provider.Id, &provider.Name); err != nil {
		return nil, err
	}

	return provider, nil
}

func (self *ProviderRepositorySql) GetByName(name string) (*entity.Provider, error) {
	row := self.db.QueryRow(getByNameProvider, name)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var entity entity.Provider
	err := row.Scan(&entity.Id, &entity.Name)

	return &entity, err
}

func (self *ProviderRepositorySql) InsertSingle(provider *entity.Provider) error {
	tx, err := self.db.Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec(insertSingleProvider, provider.Name)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (self *ProviderRepositoryMock) GetById(id int) (interface{}, error) {
	row := self.db.QueryRow(getByIdProvider)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var provider entity.Provider
	if err := row.Scan(&provider.Id, &provider.Name); err != nil {
		return nil, err
	}

	return provider, nil
}

func (self *ProviderRepositoryMock) GetByName(name string) (*entity.Provider, error) {
	row := self.db.QueryRow(getByNameProvider, name)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var entity entity.Provider
	err := row.Scan(&entity.Id, &entity.Name)

	return &entity, err
}

func (self *ProviderRepositoryMock) InsertSingle(provider *entity.Provider) error {
	tx, err := self.db.Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec(insertSingleProvider, provider.Name)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (self *ProviderRepositoryGorm) GetById(id int) (interface{}, error) {
	return nil, nil
}

func (self *ProviderRepositoryGorm) GetByName(name string) (*entity.Provider, error) {
	return nil, nil
}

func (self *ProviderRepositoryGorm) InsertSingle(provider *entity.Provider) error {
	return nil
}
