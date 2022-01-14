package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Coaster struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	InPark       string `json:"in_park"`
	Height       int    `json:"height"`
}

type CoasterHandlers struct {
	sync.Mutex
	store map[string]Coaster
}

func (h *CoasterHandlers) entry(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.RequestURI)
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
}

func (h *CoasterHandlers) get(w http.ResponseWriter, r *http.Request) {
	coasters := make([]Coaster, len(h.store))

	h.Lock()

	i := 0
	for _, coaster := range h.store {
		coasters[i] = coaster
		i++
	}

	h.Unlock()

	jsonBytes, err := json.Marshal(coasters)
	if err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader((http.StatusOK))
	w.Write(jsonBytes)
}

func (h *CoasterHandlers) single(w http.ResponseWriter, r *http.Request) {
	path := r.URL.String()

	parts := strings.Split(path, "/")

	if len(parts) != 3 {
		w.WriteHeader((http.StatusNotFound))
		return
	}

	id := parts[2]

	if id == "random" {
		h.random(w, r)
		return
	}

	h.Lock()

	coaster, ok := h.store[id]

	h.Unlock()

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(coaster)
	if err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader((http.StatusOK))
	w.Write(jsonBytes)
}

func (h *CoasterHandlers) random(w http.ResponseWriter, r *http.Request) {
	ids := make([]string, len(h.store))

	h.Lock()

	i := 0
	for id := range h.store {
		ids[i] = id
		i++
	}

	var target string

	if len(ids) == 0 {
		w.WriteHeader((http.StatusNotFound))
		return
	} else if len(ids) == 1 {
		target = ids[0]
	} else {
		rand.Seed(time.Now().UnixNano())
		target = ids[rand.Intn(len(ids))]
	}

	coaster, ok := h.store[target]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(coaster)
	if err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(err.Error()))
		return
	}

	h.Unlock()

	w.Header().Add("content-type", "application/json")
	w.Header().Add("location", fmt.Sprintf("/coasters/%s", target))
	w.WriteHeader((http.StatusFound))
	w.Write(jsonBytes)
}

func (h *CoasterHandlers) post(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

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

	var coaster Coaster
	err = json.Unmarshal(bodyBytes, &coaster)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	coaster.ID = uuid.New().String()

	h.Lock()

	h.store[coaster.ID] = coaster

	defer h.Unlock()
}

func newCoasterHandlers() *CoasterHandlers {
	return &CoasterHandlers{
		store: map[string]Coaster{},
	}
}

type AdmingPortal struct {
	password string
}

func newAdminPortal() *AdmingPortal {
	password := os.Getenv("ADMIN_PASSWORD")
	if password == "" {
		panic("Required env var ADMIN_PASSWORD not set")
	}

	return &AdmingPortal{password: password}
}

func (a AdmingPortal) handler(w http.ResponseWriter, r *http.Request) {
	_, pass, ok := r.BasicAuth()

	if !ok || pass != a.password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Access Denied"))
		return
	}

	w.Write([]byte("<html><h1>Secret admin portal</h1></html>"))
}

func main() {
	coasterHandlers := newCoasterHandlers()
	adminPortal := newAdminPortal()
	http.HandleFunc("/admin", adminPortal.handler)
	http.HandleFunc("/coasters", coasterHandlers.entry)
	http.HandleFunc("/coasters/", coasterHandlers.single)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
