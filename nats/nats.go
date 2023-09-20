package main

import (
	"bytes"
	"fmt"
	"github.com/nats-io/stan.go"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type natsServer struct {
	sc     stan.Conn
	sub    stan.Subscription
	client *http.Client
}

func (s *natsServer) handlerRequest(m *stan.Msg) {
	b := bytes.NewReader(m.Data)

	resp, err := s.client.Post("http://localhost:8080/newOrder", "application/json", b)

	if err != nil {
		fmt.Printf("error %s", err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
}

func (s *natsServer) connect() error {
	s.client = &http.Client{}

	sc, err := stan.Connect("test-cluster", "subscriber", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		return err
	}

	sub, err := sc.Subscribe("addOrder", s.handlerRequest)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	s.sc, s.sub = sc, sub
	return nil
}

func main() {
	server := natsServer{}
	err := server.connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	for {
		<-sigCh
		server.sc.Close()
		server.sub.Close()
		return
	}

}
