package api

import (
	"cellphone/internal/api/gin_api"
	"cellphone/internal/api/nat_api"
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

func NewServer(conf app_config.Main, repo *repository.RepositoryService) ApiHandler {
	switch conf.ApiType {
	case API_NAT:
		log.Println("Starting net/http API")
		return nat_api.NewServer(repo)
	case API_GIN:
		log.Println("Starting Gin API")
		return gin_api.NewServer(repo)
	}

	return nil
}
