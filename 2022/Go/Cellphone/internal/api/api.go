package api

import (
	"cellphone/internal/app_config"
	"cellphone/internal/repository"
	"log"
)

const (
	API_MOCK = iota
	API_NAT
	API_GIN
)

type ApiHandler interface {
	// A function that blocks and starts the API
	Start(config app_config.Main) error
}

func MakeRoutes(conf app_config.Main, repo *repository.RepositoryService) ApiHandler {
	switch conf.ApiType {
	case API_NAT:
		log.Println("Starting net/http API")
		return MakeNatRoutes(repo)
	case API_GIN:
		log.Println("Starting Gin API")
		return MakeGinRoutes(repo)
	}

	return nil
}
