package handler

import "main/internal/service"

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	newHandler := new(Handler)
	newHandler.service = service
	return newHandler
}
