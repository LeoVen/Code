package dbconn

import (
	"cellphone/internal/app_config"
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_MOCK = iota
	DB_MYSQL
)

func NewDbConnection(config app_config.Main) (*sql.DB, error) {
	switch config.DbType {
	case DB_MOCK:
		log.Println("Initializing Mock Database Client")
		return newMockConnection()
	case DB_MYSQL:
		log.Println("Initializing MySQL Database Client")
		return newMySQLDBConnection(config)
	}
	log.Println("Initializing MySQL Database Client")
	return newMySQLDBConnection(config)
}

func newMockConnection() (*sql.DB, error) {
	db, _, err := sqlmock.New()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func newMySQLDBConnection(mainConf app_config.Main) (*sql.DB, error) {
	config := MySQLFromConfig(mainConf)

	mysqlCfg := mysql.Config{
		User:   config.DBUser,
		Passwd: config.DBPass,
		Net:    config.DBProt,
		Addr:   config.DBAddr,
		DBName: config.DBName,
	}

	dsn := mysqlCfg.FormatDSN()

	sqlDB, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	log.Println("MySQL Database connection opened")

	return sqlDB, nil
}
