package api

import (
	"github.com/gorilla/mux"
	"main/internal/handler"
)

func ConfigureRoutes(h *handler.Handler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/order/{order_uid}", h.GetOrder).Methods("GET")
	router.HandleFunc("/newOrder", h.NewOrder).Methods("POST")
	return router
}
