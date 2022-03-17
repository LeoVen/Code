package api

import (
	"cellphone/internal/app_config"
	"cellphone/internal/entity"
	"cellphone/internal/repository"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type NativeApiService struct {
	repo *repository.RepositoryService
}

func (self *NativeApiService) Start(config app_config.Main) error {
	return http.ListenAndServe(":"+config.Flags["CELL_APIPORT"], nil)
}

func MakeNatRoutes(repo *repository.RepositoryService) *NativeApiService {
	nativeService := NativeApiService{repo}

	http.HandleFunc("/Cellphone", nativeService.getCellphone)
	http.HandleFunc("/Cellphone/", nativeService.fetchSingle)

	// delete, post, patch
	http.HandleFunc("/Provider", nativeService.handleProviderCrud)
	// byId, count
	http.HandleFunc("/Provider/", nativeService.handleProvider)
	http.HandleFunc("/Provider/ByName/", nativeService.getProviderByName)

	return &nativeService
}

func (self *NativeApiService) handleProvider(w http.ResponseWriter, r *http.Request) {
	// TODO both are gets
	// /Provider/{id}/count
	// /Provider/{id}
}

func (self *NativeApiService) handleProviderCrud(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		self.deleteProvider(w, r)
	case "POST":
		self.createProvider(w, r)
	case "PATCH":
		self.updateProvider(w, r)
	default:
		// TODO 400
	}
}

func (self *NativeApiService) getCellphone(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// [POST] /Cellphone
func (self *NativeApiService) fetchSingle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		// TODO err 400
	}

	fullPath := r.URL.Path[:]
	path := r.URL.Path

	if len(path) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if path[0] == '/' {
		path = path[1:]
	}

	paths := strings.Split(path, "/")

	if len(paths) < 2 || paths[1] == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(paths[1])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%s Could not parse id(%s): %s", fullPath, paths[1], err.Error())))
		return
	}

	cellphone, err := self.repo.Cellphone.FetchSingle(id)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte(fmt.Sprintf("%s Could not retrieve entity: %s", fullPath, err.Error())))
		return
	}

	jsonContent, err := json.Marshal(cellphone)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%s Error: %s", fullPath, err.Error())))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonContent)
}

// [GET] /Provider/{id}
func (self *NativeApiService) getProviderById(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// [GET] /Provider/{id}/count
func (self *NativeApiService) getProviderCount(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// [GET] /Provider/ByName/{name}
func (self *NativeApiService) getProviderByName(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		// TODO err 400
	}

	fullPath := r.URL.Path[:]

	path := r.URL.Path

	if len(path) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// This leading slash is optional per the net/http doc, so try to remove it whenever possible
	if path[0] == '/' {
		path = path[1:]
	}

	// Provider, ByName, XXX
	paths := strings.Split(path, "/")

	if len(paths) < 3 || paths[2] == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	name := paths[2]

	provider, err := self.repo.Provider.GetByName(name)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte(fmt.Sprintf("%s Could not retrieve entity: %s", fullPath, err.Error())))
		return
	}

	jsonContent, err := json.Marshal(provider)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%s Error: %s", fullPath, err.Error())))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Header().Add("location", fmt.Sprintf("/Provider/ById/%d", provider.Id))
	w.WriteHeader(http.StatusOK)
	w.Write(jsonContent)
}

// [DELETE] /Provider
func (self *NativeApiService) deleteProvider(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// [POST] /Provider
func (self *NativeApiService) createProvider(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")

	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("Need content-type 'application/json' but got %s", ct)))
		return
	}

	var provider entity.Provider

	err = json.Unmarshal(bytes, &provider)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = self.repo.Provider.Insert(&provider)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

// [PATCH] /Provider
func (self *NativeApiService) updateProvider(w http.ResponseWriter, r *http.Request) {
	// TODO
}
