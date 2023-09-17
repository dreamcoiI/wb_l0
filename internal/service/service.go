package service

import "main/internal/storage"

type Service struct {
	Storage storage.Storage
}

func NewService(storage *storage.Storage) *Service {
	newService := new(Service)
	newService.Storage = *storage
	return newService
}
