package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"log"

	"github.com/segmentio/kafka-go"
	_ "github.com/segmentio/kafka-go/snappy"
)

func main() {

	brokers := strings.Split("localhost:9092", ",")
	topic := "test"
	clientID := "1"
	config := kafka.ReaderConfig{
		Brokers:         brokers,
		GroupID:         clientID,
		Topic:           topic,
		MinBytes:        10e3,
		MaxBytes:        10e6,
		MaxWait:         1 * time.Second,
		ReadLagInterval: -1,
	}

	reader := kafka.NewReader(config)
	defer reader.Close()

	m, err := reader.ReadMessage(context.Background())
	if err != nil {
		fmt.Println("error in reading msgs")
		log.Fatal(err)
	}

	//value := m.Value

	fmt.Println(string(m.Value))

}
