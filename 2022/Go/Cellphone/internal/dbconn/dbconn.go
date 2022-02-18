package dbconn

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_MOCK = iota
	DB_MYSQL
)

var (
	EnvMissing = errors.New("Missing environment variables:")
)

type MongoDBConfig struct{}

func NewDbConnection(dbType int) (*sql.DB, error) {
	switch dbType {
	case DB_MOCK:
		return newMockConnection()
	case DB_MYSQL:
		return newMySQLDBConnection()
	}
	return newMySQLDBConnection()
}

func newMockConnection() (*sql.DB, error) {
	db, _, err := sqlmock.New()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func newMySQLDBConnection() (*sql.DB, error) {
	config := MySQLConfig{}

	config.LoadEnvs()

	if missing := config.CheckMissingVars(); len(missing) > 0 {
		var buff bytes.Buffer
		buff.WriteString(EnvMissing.Error())
		for _, miss := range missing {
			buff.WriteString(" ")
			buff.WriteString(miss)
		}
		return nil, errors.New(buff.String())
	}

	mysqlCfg := mysql.Config{
		User:   config.DbUser,
		Passwd: config.DbPass,
		Net:    config.DbProt,
		Addr:   fmt.Sprintf("%s:%s", config.DbHost, config.DbPort),
		DBName: config.DbName,
	}

	dsn := mysqlCfg.FormatDSN()

	sqlDB, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	log.Println("MySQL Database connection opened")

	return sqlDB, nil
}
