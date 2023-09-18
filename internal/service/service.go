package service

import (
	"errors"
	"github.com/sirupsen/logrus"
	"main/internal/model"
	"main/internal/storage"
)

type OrderService struct {
	Storage storage.Storage
	cache   map[string]model.Order
}

func NewService(storage *storage.Storage) *OrderService {
	newService := new(OrderService)
	newService.Storage = *storage
	return newService
}

func (s *OrderService) CreateOrder(newOrder model.Order) error {
	if _, err := s.cache[newOrder.OrderUID]; err {
		logrus.Print("value exist", s.cache)
		return errors.New("value exist is cache")
	}

	s.cache[newOrder.OrderUID] = newOrder

	err := s.Storage.AddOrder(newOrder)
	if err != nil {
		return err
	}

	return nil
}

func (s *OrderService) GetOrder(orderUID string) (model.Order, error) {
	res, err := s.cache[orderUID]
	if err {
		return res, errors.New("can't find UID on cache")
	}

	return res, nil
}

func NewOrder(db *storage.Storage, orders []model.Order) *OrderService {

	resService := new(OrderService)
	resService.Storage = *db
	resService.cache = make(map[string]model.Order)
	for _, order := range orders {
		resService.cache[order.OrderUID] = order
	}
	return resService
}
