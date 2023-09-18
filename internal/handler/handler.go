package handler

import (
	"encoding/json"
	"fmt"
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
	var orderUID struct {
		OrderUID string `json:"order_uid"`
	}

	err := json.NewDecoder(r.Body).Decode(&orderUID)

	if err != nil {
		WrapError(w, err)
		return
	}

	order, err := h.service.GetOrder(orderUID.OrderUID)
	if err != nil {
		WrapError(w, err)
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   order,
	}

	WrapOK(w, m)
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
