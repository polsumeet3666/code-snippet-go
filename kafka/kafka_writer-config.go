package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

// kafka writer
//var writer *kafka.writer

// Configure kafka config
func Configure(kafkaBrokerUrls []string, clientID string, topic string) (w *kafka.Writer, err error) {

	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: clientID,
	}

	config := kafka.WriterConfig{
		Brokers:      kafkaBrokerUrls,
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		Dialer:       dialer,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		//CompressionCodec: snappy.NewCompressionCodec(),
	}

	w = kafka.NewWriter(config)
	//writer = w

	return w, nil
}

func main() {
	kafkaBrokerUrl := strings.Split("localhost:9092", ",")
	kafkaClientID := "1"
	kafkaTopic := "test"

	// connect to kafka
	kafkaProducer, err := Configure(kafkaBrokerUrl, kafkaClientID, kafkaTopic)
	if err != nil {
		log.Fatal(err)
	}
	defer kafkaProducer.Close()

	fmt.Println("connected to kafka")

	// kafka msg
	msg := kafka.Message{
		Key:   []byte("sampleKey"),
		Value: []byte("from go"),
		Time:  time.Now(),
	}

	// produce msg
	err = kafkaProducer.WriteMessages(context.TODO(), msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("msg send to kafka")

}
