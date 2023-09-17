package api

import (
	"github.com/gorilla/mux"
	"main/internal/handler"
)

func ConfigureRoutes(h *handler.Handler) *mux.Route {
	router := mux.NewRouter()
	router.HandleFunc("/order/{order_uid}", handler.)
	//return router
}
