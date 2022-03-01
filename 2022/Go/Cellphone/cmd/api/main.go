package main

import (
	"cellphone/internal/api"
	"cellphone/internal/app_config"
	"cellphone/internal/entry"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

	if apiEngine == nil {
		panic("Error: API Engine is nil")
	}

	switch engine := apiEngine.(type) {
	case *gin.Engine:
		err = engine.Run()
	case *api.NativeApiService:
		err = http.ListenAndServe(":"+config.Flags["CELL_APIPORT"], nil)
	}

	if err != nil {
		panic(fmt.Sprintf("Error initializing API: %s", err.Error()))
	}
}
