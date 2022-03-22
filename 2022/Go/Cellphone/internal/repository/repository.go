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
	// Retrieves a cellphone and deletes it from the database
	FetchSingle(providerId int) (*entity.Cellphone, error)
	// Inserts multiple cellphones to a certain provider
	BulkInsert(providerId int, entities []entity.Cellphone) error
}

type ProviderRepository interface {
	Repository
	// Gets a provider by name
	GetByName(name string) (*entity.Provider, error)
	// Gets how many cellphones a provider has by its ID
	GetCount(id int) (*int, error)
	// Creates a new provider
	Insert(provider *entity.Provider) error
	// Deletes an existing provider
	Delete(id int) error
	// Updates an existing provider
	Update(provider *entity.Provider) error
}

type RepositoryService struct {
	Cellphone CellphoneRepository
	Provider  ProviderRepository
}

func Initialize(conf app_config.Main, db *sql.DB) (*RepositoryService, error) {
	if conf.RepoType == REPO_MOCK {
		log.Println("Starting MOCK Repository")
	} else if conf.RepoType == REPO_SQL {
		log.Println("Starting SQL Repository")
	} else if conf.RepoType == REPO_GORM {
		log.Println("Starting GORM Repository")
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
