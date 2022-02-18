package api

import (
	"cellphone/internal/repository"
)

const (
	API_NAT = iota
	API_GIN
)

func MakeRoutes(apiType int, repo *repository.RepositoryService) interface{} {
	switch apiType {
	case API_NAT:
		return MakeNatRoutes(repo)
	case API_GIN:
		return MakeGinRoutes(repo)
	}

	return nil
}
