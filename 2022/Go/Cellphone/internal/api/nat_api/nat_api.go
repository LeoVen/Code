package nat_api

import (
	"cellphone/internal/api/middleware"
	"cellphone/internal/app_config"
	"cellphone/internal/entity"
	"cellphone/internal/repository"
	"cellphone/internal/telemetry"
	"encoding/json"
	"net/http"
	"strings"
)

type NativeApiService struct {
	http.Handler
	repo      *repository.RepositoryService
	telemetry *telemetry.Telemetry
}

func (self *NativeApiService) Start(config app_config.Main) error {
	return http.ListenAndServe(":"+config.Flags["CELL_APIPORT"], self)
}

func NewServer(repo *repository.RepositoryService) *NativeApiService {
	server := &NativeApiService{
		repo:      repo,
		telemetry: nil,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/Cellphone", server.getCellphone)
	mux.HandleFunc("/Cellphone/", server.fetchSingle)

	// delete, post, patch
	mux.HandleFunc("/Provider", server.handleProviderCrud)
	// byId, count
	mux.HandleFunc("/Provider/", server.handleProvider)
	mux.HandleFunc("/Provider/ByName/", server.getProviderByName)

	server.Handler = middleware.AuthMiddleware(mux)

	return server
}

func badRequest(w *http.ResponseWriter) {
	data := entity.ApiError{
		Error: "Bad Request",
	}

	jsonBytes, _ := json.Marshal(&data)

	(*w).WriteHeader(http.StatusBadRequest)
	(*w).Header().Add("content-type", "application/json")
	(*w).Write(jsonBytes)
}

func serverError(w *http.ResponseWriter, err error) {
	// e := errors.NewError(err)

	data := entity.ApiError{
		Error: "Internal Server Error",
	}

	jsonBytes, _ := json.Marshal(&data)

	(*w).WriteHeader(http.StatusInternalServerError)
	(*w).Header().Add("content-type", "application/json")
	(*w).Write(jsonBytes)
}

func notFound(w *http.ResponseWriter) {
	data := entity.ApiError{
		Error: "Not found",
	}

	jsonBytes, _ := json.Marshal(&data)

	(*w).WriteHeader(http.StatusNotFound)
	(*w).Header().Add("content-type", "application/json")
	(*w).Write(jsonBytes)
}

func getSplits(r *http.Request) []string {
	path := r.URL.Path

	if len(path) == 0 {
		return []string{}
	}

	if path[0] == '/' {
		path = path[1:]
	}

	return strings.Split(path, "/")
}
