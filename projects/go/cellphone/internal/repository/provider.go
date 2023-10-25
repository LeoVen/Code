package repository

import (
	"cellphone/internal/repository/repo_orm"
	"cellphone/internal/repository/repo_sql"
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewProviderRepository(repoType int, conn *sql.DB) (ProviderRepository, error) {
	switch repoType {
	case REPO_MOCK:
		return newProviderRepositorySql(conn)
	case REPO_SQL:
		return newProviderRepositorySql(conn)
	case REPO_GORM:
		return newProviderRepositoryGorm(conn)
	}
	return newProviderRepositorySql(conn)
}

func newProviderRepositorySql(conn *sql.DB) (ProviderRepository, error) {
	return &repo_sql.ProviderRepository{Db: conn}, nil
}

func newProviderRepositoryGorm(conn *sql.DB) (ProviderRepository, error) {
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: conn,
	}))

	if err != nil {
		return nil, err
	}

	return &repo_orm.ProviderRepository{Db: gormDB}, nil
}
