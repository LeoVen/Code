package gin_api

import (
	"cellphone/internal/api/middleware"
	"cellphone/internal/app_config"
	"cellphone/internal/repository"
	"cellphone/internal/telemetry"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinApi struct {
	http.Handler
	engine *gin.Engine
	tel    *telemetry.Telemetry
}

func (self *GinApi) Start(config app_config.Main) error {
	return http.ListenAndServe(":"+config.Flags["CELL_APIPORT"], self)
}

func NewServer(repo *repository.RepositoryService, tel *telemetry.Telemetry) *GinApi {

	r := gin.Default()

	makeGinRoutes(repo, r)

	server := &GinApi{
		engine:  r,
		tel:     tel,
		Handler: middleware.AuthMiddleware(r),
	}

	return server
}
