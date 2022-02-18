package api

import (
	"cellphone/internal/entity"
	"cellphone/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MakeGinRoutes(repo *repository.RepositoryService) interface{} {

	r := gin.Default()

	r.GET("/Provider/ByName/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")

		if name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "URL Parameter name is missing"})
			return
		}

		entity, err := repo.Provider.GetByName(name)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, entity)
	})

	r.GET("/Provider/ById/:id", func(ctx *gin.Context) {
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

	r.POST("/Provider", func(ctx *gin.Context) {
		var provider entity.Provider
		if err := ctx.ShouldBindJSON(&provider); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := repo.Provider.InsertSingle(&provider); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.Status(http.StatusCreated)
	})

	r.GET("/Cellphone/:providerId", func(ctx *gin.Context) {
		providerParam := ctx.Param("providerId")

		if providerParam == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "URL Parameter providerId is missing"})
			return
		}

		providerId, err := strconv.Atoi(providerParam)

		if err != nil {
			ctx.Error(err)
			return
		}

		phone, err := repo.Cellphone.ServeSingleFromProvider(providerId)

		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, phone)
	})

	r.POST("/Cellphone", func(ctx *gin.Context) {
		var cellphone entity.Cellphone
		if err := ctx.ShouldBindJSON(&cellphone); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := repo.Cellphone.InsertSingle(&cellphone); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.Status(http.StatusCreated)
	})

	return r
}
