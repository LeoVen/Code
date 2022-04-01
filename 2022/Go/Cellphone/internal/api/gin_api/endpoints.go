package gin_api

import (
	"cellphone/internal/entity"
	"cellphone/internal/repository"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func makeGinRoutes(repo *repository.RepositoryService, r *gin.Engine) {

	r.GET("/Cellphone/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		if idParam == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "URL Parameter id is missing"})
			return
		}

		id, err := strconv.Atoi(idParam)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		entity, err := repo.Cellphone.GetById(id)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, entity)
	})

	r.POST("/Cellphone/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		if idParam == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required {id} parameter"})
			return
		}

		id, err := strconv.Atoi(idParam)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		bytes, err := ioutil.ReadAll(ctx.Request.Body)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var values []*entity.Cellphone

		err = json.Unmarshal(bytes, &values)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = repo.Cellphone.BulkInsert(id, values)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.Status(http.StatusOK)
	})

	r.POST("/Cellphone", func(ctx *gin.Context) {
		// TODO
		ctx.Status(http.StatusNotImplemented)
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

	r.GET("/Provider/:id/count", func(ctx *gin.Context) {
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

		count, err := repo.Provider.GetCount(id)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"count": count})
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
		var entity entity.Provider
		if err := ctx.ShouldBindJSON(&entity); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if entity.Id == 0 && entity.Name != "" {
			p, err := repo.Provider.GetByName(entity.Name)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			entity.Id = p.Id
		} else if entity.Id == 0 && entity.Name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "At least one of Id or Name is required"})
			return
		}

		err := repo.Provider.Delete(int(entity.Id))

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, &entity)
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
		ctx.Status(http.StatusNotImplemented)
	})
}
