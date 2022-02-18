package main

import (
	"cellphone/internal/api"
	"cellphone/internal/dbconn"
	"cellphone/internal/entry"
	"cellphone/internal/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const URL_ROOT = "http://localhost:8080/"
const FROM = 1_000_000_000
const TO = 2_000_000_000
const STEP = 100_000

func main() {
	// TODO add application configuration
	db, _, apiEngine, err := entry.InitializeBackend(dbconn.DB_MYSQL, repository.REPO_SQL, api.API_NAT)

	if err != nil {
		fmt.Println("Error initializing backend:\n", err.Error())
		return
	}

	defer db.Close()

	if apiEngine == nil {
		fmt.Println("Error: API is nil:\n", err.Error())
		return
	}

	switch engine := apiEngine.(type) {
	case *gin.Engine:
		err = engine.Run()
	case *api.NativeApiService:
		err = http.ListenAndServe(":8080", nil)
	}

	if err != nil {
		fmt.Println("Error initializing API:\n", err.Error())
		return
	}
}
