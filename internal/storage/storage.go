package storage

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
	"main/internal/model"
)

type Storage struct {
	db *pg.DB
}

func NewStorage(conn *pg.DB) *Storage {
	storage := new(Storage)
	storage.db = conn
	return storage
}

func (s *Storage) MigrateStorage() ([]model.Order, error) {
	models := []interface{}{
		(*model.Order)(nil),
	}

	var err error

	for _, m := range models {
		option := orm.CreateTableOptions{IfNotExists: true}
		err = s.db.Model(m).CreateTable(&option)
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}
	}

	orders := make([]model.Order, 0)
	err = s.db.Model(&orders).Select()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return orders, nil
}

func (s *Storage) AddOrder(order model.Order) error {
	if _, err := s.db.Model(&order).Insert(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
