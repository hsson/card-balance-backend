// Copyright (c) 2017 Alexander Håkansson
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package backend

import (
	"errors"
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

type response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

var errorNotFound = errors.New("Not found")

// newRouter creates a router that handles the routes
// specified in the routes.go file
func newRouter() *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(errorHandler(route.HandlerFunc))
	}

	router.NotFoundHandler = errorHandler(notFound)

	return router
}

type errorHandler func(http.ResponseWriter, *http.Request) (interface{}, error)

func (fn errorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res := response{}
	jsonEnc := json.NewEncoder(w)
	data, err := fn(w, r)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		log.Printf("Error: %v", err.Error())
		res.Success = false
		res.Error = err.Error()
		switch err {
		case errorNotFound:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		res.Success = true
		res.Error = ""
		w.WriteHeader(http.StatusOK)
	}
	res.Data = data
	jsonEnc.Encode(res)
}

func notFound(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return nil, errorNotFound
}
