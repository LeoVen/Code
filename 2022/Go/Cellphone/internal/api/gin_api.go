package api

import (
	"cellphone/internal/app_config"
	"cellphone/internal/entity"
	"cellphone/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GinApi struct {
	engine *gin.Engine
}

func (self *GinApi) Start(config app_config.Main) error {
	return self.engine.Run()
}

func MakeGinRoutes(repo *repository.RepositoryService) *GinApi {

	r := gin.Default()

	r.GET("/Cellphone/:id", func(ctx *gin.Context) {
		// TODO
	})

	r.POST("/Cellphone", func(ctx *gin.Context) {
		// TODO
	})

	r.GET("/Provider/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		if idParam == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "URL Parameter id is missing"})
			return
		}

		id, err := strconv.Atoi(idParam)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		entity, err := repo.Provider.GetById(id)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, entity)
	})

	r.GET("/Provider/:id/Count", func(ctx *gin.Context) {
		// TODO
	})

	r.GET("/Provider/ByName/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")

		if name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Query name is missing"})
			return
		}

		entity, err := repo.Provider.GetByName(name)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, entity)
	})

	r.DELETE("/Provider", func(ctx *gin.Context) {
		// TODO
	})

	r.POST("/Provider", func(ctx *gin.Context) {
		var provider entity.Provider
		if err := ctx.ShouldBindJSON(&provider); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := repo.Provider.Insert(&provider); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.Status(http.StatusCreated)
	})

	r.PATCH("/Provider", func(ctx *gin.Context) {
		// TODO
	})

	return &GinApi{r}
}
