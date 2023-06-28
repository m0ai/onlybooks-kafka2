package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"strings"
	"sync"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Net.MaxOpenRequests = 1
	config.Producer.Idempotent = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	// max in flight requests to kafka broker

	config.Producer.Transaction.ID = "peter-transaction-02"

	//brokers := []string{"peter-kafka01.foo.bar:9092"}
	brokers := "my-cluster-kafka-bootstrap:9092"
	//brokers := "localhost:909"
	targetTopic := "peter-test05"
	log.Println("hello..")

	//producerProvider := newProducerProvider(strings.Split(brokers, ","), config)
	producer, err := sarama.NewSyncProducer(strings.Split(brokers, ","), config)
	if err != nil {
		log.Panicf("Error creating producer: %v", err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Panicf("Error closing producer: %v", err)
		}
	}()


	log.Println("let's begin hhtransaction")
	err = producer.BeginTxn()
	if err != nil {
		log.Panicf("Error Begin transaction: %v", err)
	}

	//log.Println("let's send message")
	//producer.Input() <- &sarama.ProducerMessage{
	//	Topic: targetTopic,
	//	Key:   nil,
	//	Value: sarama.StringEncoder("Apache Kafka is a distributed streaming platform"),
	//}

	// Produce some records in transaction
	//partion, := producer.SendMessage(&sarama.ProducerMessage{Topic: targetTopic, Key: nil, Value: sarama.StringEncoder("test")}

	msg := &sarama.ProducerMessage{
		Topic: targetTopic,
		Key:   nil,
		Value: sarama.StringEncoder("Apache Kafka is a distributed streaming platform"),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Panicf("Failed to send message: %v", err)
	} else {
		log.Println("message sent to partition ", partition, " at offset ", offset)
	}

	err = producer.CommitTxn() // or AbortTxn()
// pool of producers that ensure transactional-id is unique.
type producerProvider struct {
	transactionIdGenerator int32

	producersLock sync.Mutex
	producers     []sarama.AsyncProducer

	producerProvider func() sarama.AsyncProducer
}

func newProducerProvider(brokers []string, producerConfigurationProvider func() *sarama.Config) *producerProvider {
	provider := &producerProvider{}
	provider.producerProvider = func() sarama.AsyncProducer {
		config := producerConfigurationProvider()
		suffix := provider.transactionIdGenerator
		// Append transactionIdGenerator to current config.Producer.Transaction.ID to ensure transaction-id uniqueness.
		if config.Producer.Transaction.ID != "" {
			provider.transactionIdGenerator++
			config.Producer.Transaction.ID = config.Producer.Transaction.ID + "-" + fmt.Sprint(suffix)
		}
		producer, err := sarama.NewAsyncProducer(brokers, config)
		if err != nil {
			return nil
		}
		return producer
	}
	return provider
}
