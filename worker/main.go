package main

import (
	"context"
	"encoding/json"
	"fmt"
	repositoryAdapter "github.com/brendontj/didactic-eureka/adapter/repository"
	"github.com/brendontj/didactic-eureka/core/entity"
	"github.com/brendontj/didactic-eureka/core/repository"
	"github.com/brendontj/didactic-eureka/infrastructure/postgres"
	"github.com/brendontj/didactic-eureka/infrastructure/rabbitmq"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"sync"
)

type Worker struct {
	*rabbitmq.Client
	Repo repository.Repository
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	w := &Worker{}

	rabbitmqClient, err := rabbitmq.NewClient(rabbitmq.Config{
		User: os.Getenv("RABBITMQ_USER"),
		Pass: os.Getenv("RABBITMQ_PASSWORD"),
		Host: os.Getenv("RABBITMQ_HOST"),
		Port: os.Getenv("RABBITMQ_PORT"),
	})
	if err != nil {
		panic(err)
	}

	w.Client = rabbitmqClient
	pgDBInfra, err := postgres.NewDB(postgres.Config{
		User:   os.Getenv("POSTGRES_USER"),
		Pass:   os.Getenv("POSTGRES_PASSWORD"),
		Host:   os.Getenv("POSTGRES_HOST"),
		Port:   os.Getenv("POSTGRES_PORT"),
		DBName: os.Getenv("POSTGRES_DB"),
	})
	if err != nil {
		panic(err)
	}

	w.Repo = repositoryAdapter.NewAdapter(pgDBInfra)
	if err := w.Client.QueueDeclare(os.Getenv("SAVE_CUSTOMER_QUEUE")); err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	numWorkers, err := strconv.Atoi(os.Getenv("NUM_WORKERS"))
	if err != nil {
		panic(err)
	}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			w.consumeMessages()
		}()
	}

	wg.Wait()
}

func (w *Worker) consumeMessages() {
	msgs, err := w.Channel.Consume(
		os.Getenv("SAVE_CUSTOMER_QUEUE"), // queue
		"",                               // consumer
		true,                             // auto-ack
		false,                            // exclusive
		false,                            // no-local
		false,                            // no-wait
		nil,                              // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	for msg := range msgs {
		fmt.Printf("Received a message: %s\n", msg.Body)
		var customer entity.Customer
		if err = json.Unmarshal(msg.Body, &customer); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			continue
		}
		fmt.Println("Saving customer...")
		if err = w.Repo.Save(context.Background(), customer); err != nil {
			log.Printf("Failed to save customer: %v", err)
			continue
		}
	}
}
