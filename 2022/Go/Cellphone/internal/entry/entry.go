package entry

import (
	"cellphone/internal/api"
	config "cellphone/internal/app_config"
	"cellphone/internal/dbconn"
	"cellphone/internal/repository"
	"cellphone/internal/telemetry"
	"database/sql"
	"os"
)

// InitializeBackend returns the db connection, repository, apiEngine and error
// The apiEngine is not initialized
func InitializeBackend(conf config.Main) (*sql.DB, *repository.RepositoryService, api.ApiHandler, error) {
	tel := telemetry.NewTelemetry(os.Stdout) // TODO hook-up with other services

	db, err := dbconn.NewDbConnection(conf)
	if err != nil {
		return nil, nil, nil, err
	}

	repo, err := repository.Initialize(conf, db)
	if err != nil {
		return nil, nil, nil, err
	}

	apiEngine := api.NewServer(&conf, repo, tel)

	return db, repo, apiEngine, nil
}
