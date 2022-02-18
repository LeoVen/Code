package repository

import (
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

func Initialize(repoType int, db *sql.DB) (*RepositoryService, error) {
	if repoType == REPO_SQL {
		log.Println("Starting SQL Repository")
	} else if repoType == REPO_GORM {
		log.Println("Starting GORM Repository")
	} else if repoType == REPO_MOCK {
		log.Println("Starting MOCK Repository")
	}

	repo := &RepositoryService{}

	cellphone, err := NewCellphoneRepository(repoType, db)
	if err != nil {
		return nil, err
	}
	provider, err := NewProviderRepository(repoType, db)
	if err != nil {
		return nil, err
	}

	repo.Cellphone = cellphone
	repo.Provider = provider

	return repo, nil
}
