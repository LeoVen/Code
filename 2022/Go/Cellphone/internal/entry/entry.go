package entry

import (
	"cellphone/internal/api"
	config "cellphone/internal/app_config"
	"cellphone/internal/dbconn"
	"cellphone/internal/repository"
	"database/sql"
)

// Returns the db connection, repository, apiEngine and error
func InitializeBackend(conf config.Main) (*sql.DB, *repository.RepositoryService, interface{}, error) {
	db, err := dbconn.NewDbConnection(conf)
	if err != nil {
		return nil, nil, nil, err
	}

	repo, err := repository.Initialize(conf, db)
	if err != nil {
		return nil, nil, nil, err
	}

	apiEngine := api.MakeRoutes(conf, repo)

	return db, repo, apiEngine, nil
}
