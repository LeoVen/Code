package entry

import (
	"cellphone/internal/api"
	config "cellphone/internal/app_config"
	"cellphone/internal/dbconn"
	"cellphone/internal/repository"
	"database/sql"
)

// InitializeBackend returns the db connection, repository, apiEngine and error
// The apiEngine is not initialized
func InitializeBackend(conf config.Main) (*sql.DB, *repository.RepositoryService, api.ApiHandler, error) {
	db, err := dbconn.NewDbConnection(conf)
	if err != nil {
		return nil, nil, nil, err
	}

	repo, err := repository.Initialize(conf, db)
	if err != nil {
		return nil, nil, nil, err
	}

	apiEngine := api.NewServer(conf, repo)

	return db, repo, apiEngine, nil
}
