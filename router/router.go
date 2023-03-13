package router

import (
	"github.com/dmytrodemianchuk/cz-api/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/peoples", controller.GetMyAllNames).Methods("GET")
	router.HandleFunc("/api/people", controller.CreateName).Methods("POST")
	router.HandleFunc("/api/people/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/people/{id}", controller.DeleteAName).Methods("DELETE")
	router.HandleFunc("/api/deleteall", controller.DeleteAllNames).Methods("DELETE")

	return router
}
