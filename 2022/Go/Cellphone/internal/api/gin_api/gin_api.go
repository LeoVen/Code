package gin_api

import (
	"cellphone/internal/api/middleware"
	"cellphone/internal/app_config"
	"cellphone/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinApi struct {
	http.Handler
	engine *gin.Engine
}

func (self *GinApi) Start(config app_config.Main) error {
	return http.ListenAndServe(":"+config.Flags["CELL_APIPORT"], self)
}

func NewServer(repo *repository.RepositoryService) *GinApi {

	r := gin.Default()

	makeGinRoutes(repo, r)

	server := &GinApi{
		engine:  r,
		Handler: middleware.AuthMiddleware(r),
	}

	return server
}
