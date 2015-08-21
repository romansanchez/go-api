package main

import (
  "github.com/gorilla/mux"
)

func ApiRouter() *mux.Router {
	router := mux.NewRouter()
	for _,route := range routes {
		router.HandleFunc(route.Pattern, route.HandlerFunc).Methods(route.Method)
	}
	return router
}