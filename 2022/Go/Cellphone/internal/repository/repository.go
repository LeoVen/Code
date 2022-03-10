package repository

import (
	"cellphone/internal/app_config"
	"cellphone/internal/entity"
	"database/sql"
	"log"
)

const (
	REPO_MOCK = iota
	REPO_SQL
	REPO_GORM
)

type Repository interface {
	GetById(id int) (interface{}, error)
}

type CellphoneRepository interface {
	Repository
	ServeSingleFromProvider(providerId int) (*entity.Cellphone, error)
	InsertSingle(cellphone *entity.Cellphone) error
	GetAllByProviderName(name string) ([]*entity.Cellphone, error)
	DeleteAllFromProvider(providerId int) (int, error)
}

type ProviderRepository interface {
	Repository
	GetByName(name string) (*entity.Provider, error)
	InsertSingle(provider *entity.Provider) error
}

type RepositoryService struct {
	Cellphone CellphoneRepository
	Provider  ProviderRepository
}

func Initialize(conf app_config.Main, db *sql.DB) (*RepositoryService, error) {
	if conf.RepoType == REPO_SQL {
		log.Println("Starting SQL Repository")
	} else if conf.RepoType == REPO_GORM {
		log.Println("Starting GORM Repository")
	} else if conf.RepoType == REPO_MOCK {
		log.Println("Starting MOCK Repository")
	}

	repo := &RepositoryService{}

	cellphone, err := NewCellphoneRepository(conf.RepoType, db)
	if err != nil {
		return nil, err
	}
	provider, err := NewProviderRepository(conf.RepoType, db)
	if err != nil {
		return nil, err
	}

	repo.Cellphone = cellphone
	repo.Provider = provider

	log.Println("Repository Initialized")

	return repo, nil
}
