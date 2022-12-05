package main

import (
	"cellphone/internal/app_config"
	"cellphone/internal/entry"
	"fmt"
)

var REQ_ARGS = []string{
	"CELL_DBADDR",
	"CELL_DBNAME",
	"CELL_DBPASS",
	"CELL_DBPROT",
	"CELL_DBUSER",
	"CELL_APIPORT",
	"CELL_APITYPE",
	"CELL_DBTYPE",
	"CELL_REPOTYPE",
}

func main() {
	var required map[string]bool = make(map[string]bool)

	for _, req := range REQ_ARGS {
		required[req] = true
	}

	config, err := app_config.GetConfig(required)

	if err != nil {
		panic(fmt.Sprintf("Error parsing arguments: %s", err.Error()))
	}

	db, _, apiEngine, err := entry.InitializeBackend(config)

	if err != nil {
		panic(fmt.Sprintf("Error initializing backend: %s", err.Error()))
	}

	defer db.Close()

	err = apiEngine.Start(config)

	if err != nil {
		panic(fmt.Sprintf("Error initializing API: %s", err.Error()))
	}
}
