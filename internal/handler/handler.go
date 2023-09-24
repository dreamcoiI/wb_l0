package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"main/internal/model"
	"main/internal/service"
	"net/http"
)

type Handler struct {
	service *service.OrderService
}

func NewHandler(service *service.OrderService) *Handler {
	newHandler := new(Handler)
	newHandler.service = service
	return newHandler
}

func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars["order_uid"] == "" {
		WrapError(w, errors.New("missing id"))
		return
	}

	order, err := h.service.GetOrder(vars["order_uid"])
	if err != nil {
		WrapError(w, err)
		return
	}

	var response = map[string]interface{}{
		"result": "OK",
		"data":   order,
	}

	resp, err := json.Marshal(response)
	if err != nil {
		WrapError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		WrapError(w, err)
		return
	}

	WrapOK(w, response)
}

func (h *Handler) NewOrder(w http.ResponseWriter, r *http.Request) {

	var newOrder model.Order

	err := json.NewDecoder(r.Body).Decode(&newOrder)

	if err != nil {
		WrapError(w, err)
		return
	}

	err = h.service.CreateOrder(newOrder)

	if err != nil {
		WrapError(w, err)
		return
	}

	var response = map[string]interface{}{
		"result": "OK",
		"data":   "",
	}

	resp, err := json.Marshal(response)
	if err != nil {
		WrapError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		WrapError(w, err)
		return
	}

	WrapOK(w, response)
}

func WrapOK(w http.ResponseWriter, m map[string]interface{}) {
	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Println(w, string(res))
}

func WrapError(w http.ResponseWriter, err error) {
	WrapErrorStatus(w, err, http.StatusBadRequest)
}

func WrapErrorStatus(w http.ResponseWriter, err error, httpStatus int) {
	var m = map[string]string{
		"result": "Error",
		"data":   err.Error(),
	}
	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(httpStatus)
	fmt.Println(w, string(res))
}
