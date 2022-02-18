package dbconn

import "os"

type MySQLConfig struct {
	DbHost string
	DbName string
	DbPass string
	DbPort string
	DbProt string
	DbUser string
}

func (c *MySQLConfig) LoadEnvs() {
	c.DbHost = os.Getenv("CELL_DBHOST")
	c.DbName = os.Getenv("CELL_DBNAME")
	c.DbPass = os.Getenv("CELL_DBPASS")
	c.DbPort = os.Getenv("CELL_DBPORT")
	c.DbProt = os.Getenv("CELL_DBPROT")
	c.DbUser = os.Getenv("CELL_DBUSER")
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
