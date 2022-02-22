package dbconn

import (
	"cellphone/internal/app_config"
)

type MySQLConfig struct {
	DbHost string
	DbName string
	DbPass string
	DbPort string
	DbProt string
	DbUser string
}

func MySQLFromConfig(conf app_config.Main) MySQLConfig {
	return MySQLConfig{
		DbHost: conf.Flags["CELL_DBHOST"],
		DbName: conf.Flags["CELL_DBNAME"],
		DbPass: conf.Flags["CELL_DBPASS"],
		DbPort: conf.Flags["CELL_DBPORT"],
		DbProt: conf.Flags["CELL_DBPROT"],
		DbUser: conf.Flags["CELL_DBUSER"],
	}
}

func (c *MySQLConfig) CheckMissingVars() []string {
	result := make([]string, 0)

	if c.DbHost == "" {
		result = append(result, "CELL_DBHOST")
	}
	if c.DbName == "" {
		result = append(result, "CELL_DBNAME")
	}
	if c.DbPass == "" {
		result = append(result, "CELL_DBPASS")
	}
	if c.DbPort == "" {
		result = append(result, "CELL_DBPORT")
	}
	if c.DbProt == "" {
		result = append(result, "CELL_DBPROT")
	}
	if c.DbUser == "" {
		result = append(result, "CELL_DBUSER")
	}

	return result
}
