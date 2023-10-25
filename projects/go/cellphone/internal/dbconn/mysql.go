package dbconn

import (
	"cellphone/internal/app_config"
)

type MySQLConfig struct {
	DBAddr string
	DBName string
	DBPass string
	DBPort string
	DBProt string
	DBUser string
}

func MySQLFromConfig(conf app_config.Main) MySQLConfig {
	return MySQLConfig{
		DBAddr: conf.Flags["CELL_DBADDR"],
		DBName: conf.Flags["CELL_DBNAME"],
		DBPass: conf.Flags["CELL_DBPASS"],
		DBProt: conf.Flags["CELL_DBPROT"],
		DBUser: conf.Flags["CELL_DBUSER"],
	}
}

func (c *MySQLConfig) CheckMissingVars() []string {
	result := make([]string, 0)

	if c.DBAddr == "" {
		result = append(result, "CELL_DBADDR")
	}
	if c.DBName == "" {
		result = append(result, "CELL_DBNAME")
	}
	if c.DBPass == "" {
		result = append(result, "CELL_DBPASS")
	}
	if c.DBProt == "" {
		result = append(result, "CELL_DBPROT")
	}
	if c.DBUser == "" {
		result = append(result, "CELL_DBUSER")
	}

	return result
}
