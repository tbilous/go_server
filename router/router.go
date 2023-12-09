package router

import (
	"go_server/services"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/posts", services.GetPosts).Methods("GET")
	return router
}
