package storage

import (
	"database/sql"
	"main/internal/model"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(conn *sql.DB) *Storage {
	storage := &Storage{}
	storage.db = conn
	return storage
}

func (s *Storage) AddOrder(order model.Order) error {
	if _, err := s.db.
}
