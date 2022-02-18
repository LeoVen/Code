package entry

import (
	"cellphone/internal/api"
	"cellphone/internal/dbconn"
	"cellphone/internal/repository"
	"database/sql"
)

// Returns the db connection, repository, apiEngine and error
func InitializeBackend(dbType int, repoType int, apiType int) (*sql.DB, *repository.RepositoryService, interface{}, error) {
	db, err := dbconn.NewDbConnection(dbType)
	if err != nil {
		return nil, nil, nil, err
	}

	repo, err := repository.Initialize(repoType, db)
	if err != nil {
		return nil, nil, nil, err
	}

	apiEngine := api.MakeRoutes(apiType, repo)

	return db, repo, apiEngine, nil
}
