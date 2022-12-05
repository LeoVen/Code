package repository

import (
	"cellphone/internal/repository/repo_orm"
	"cellphone/internal/repository/repo_sql"
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewCellphoneRepository(repoType int, conn *sql.DB) (CellphoneRepository, error) {
	switch repoType {
	case REPO_MOCK:
		return newCellphoneRepositorySql(conn)
	case REPO_SQL:
		return newCellphoneRepositorySql(conn)
	case REPO_GORM:
		return newCellphoneRepositoryGorm(conn)
	}
	return newCellphoneRepositorySql(conn)
}

func newCellphoneRepositorySql(conn *sql.DB) (*repo_sql.CellphoneRepository, error) {
	return &repo_sql.CellphoneRepository{Db: conn}, nil
}

func newCellphoneRepositoryGorm(conn *sql.DB) (*repo_orm.CellphoneRepository, error) {
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: conn,
	}))

	if err != nil {
		return nil, err
	}

	return &repo_orm.CellphoneRepository{Db: gormDB}, nil
}
