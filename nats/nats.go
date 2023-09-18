package nats

import (
	"github.com/nats-io/stan.go"
	"net/http"
)

type natsServer struct {
	sc     stan.Conn
	sub    stan.Subscription
	client *http.Client
}

func (s *natsServer) connect() error {
	s.client = &http.Client{}

	_, err := stan.Connect("test-cluster", "subscriber", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		return err
	}

	//sub, err := sc.Subscribe("addOrder", )
	if err != nil {
		return err
	}
	return nil
}
