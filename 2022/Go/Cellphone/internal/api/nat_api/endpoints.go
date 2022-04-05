package nat_api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	pb "cellphone/protos/go"
)

func (self *NativeApiService) handleProvider(w http.ResponseWriter, r *http.Request) {
	splits := getSplits(r)

	switch len(splits) {
	case 2:
		// Provider/{id}
		self.getProviderById(w, r, splits)
	case 3:
		// Provider/{id}/count
		self.getProviderCount(w, r, splits)
	default:
		badRequest(&w)
	}
}

func (self *NativeApiService) handleCellphone(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		self.getCellphone(w, r)
	case "POST":
		self.bulkInsert(w, r)
	default:
		badRequest(&w)
	}
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
		badRequest(&w)
	}
}

func (self *NativeApiService) getCellphone(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (self *NativeApiService) bulkInsert(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// [POST] /Cellphone
func (self *NativeApiService) fetchSingle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		badRequest(&w)
		return
	}

	splits := getSplits(r)

	if len(splits) < 2 || splits[1] == "" {
		badRequest(&w)
		return
	}

	id, err := strconv.Atoi(splits[1])

	if err != nil {
		serverError(&w, err)
		return
	}

	cellphone, err := self.repo.Cellphone.FetchSingle(id)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	jsonContent, err := json.Marshal(cellphone)

	if err != nil {
		serverError(&w, err)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonContent)
}

// [GET] /Provider/{id}
func (self *NativeApiService) getProviderById(w http.ResponseWriter, r *http.Request, splits []string) {
	if r.Method != "GET" {
		badRequest(&w)
		return
	}
}

// [GET] /Provider/{id}/count
func (self *NativeApiService) getProviderCount(w http.ResponseWriter, r *http.Request, splits []string) {
	// TODO
}

// [GET] /Provider/ByName/{name}
func (self *NativeApiService) getProviderByName(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		badRequest(&w)
		return
	}

	paths := getSplits(r)

	if len(paths) == 0 {
		badRequest(&w)
		return
	}

	if len(paths) < 3 || paths[2] == "" {
		badRequest(&w)
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
		return
	}

	jsonContent, err := json.Marshal(provider)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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

	var provider pb.Provider

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
