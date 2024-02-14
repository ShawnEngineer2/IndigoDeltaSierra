package svcclient

import (
	"fmt"
	"indigodeltasierra/svcclient/models"
	"os"

	"sync"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func ProduceSingleKafkaMessage(event *models.EventData, topic string, brokers string) error {
	//This routine produces a single message to the Kafka queue

	events := make([]models.EventData, 1)

	events[0] = *event

	ProduceKafkaMessages(&events, topic, brokers)

	return nil

}

func ProduceKafkaMessages(events *[]models.EventData, topic string, brokers string) error {
	//This function produces to the indicated Topic and Partition via the indicated Broker(s)

	//Uses the Confluent Kafka library to produce messages into a Redpanda topic

	//Create new Producer instance
	hostName, err := os.Hostname()

	if err != nil {
		fmt.Println(err.Error())
	}

	kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"client.id":         hostName,
		"acks":              "all"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	//Produce into the queue synchronously
	var wg sync.WaitGroup

	for _, event := range *events {

		wg.Add(1)
		go func(event models.EventData) {
			defer wg.Done()
			sendKafkaMessage(kafkaProducer, &topic, &event)
		}(event)

	}

	wg.Wait()

	if err != nil {
		fmt.Printf("Failed to produce message: %s\n", err)
		os.Exit(1)
	}

	//fmt.Println("Message(s) Produced")
	kafkaProducer.Flush(15 * 1000)
	//fmt.Println("End Of Program")

	return nil
}

func sendKafkaMessage(producer *kafka.Producer, topic *string, event *models.EventData) error {

	//Create the Kafka message
	delivery_chan := make(chan kafka.Event, 10000)

	kafkaMessage := kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: topic, Partition: int32(event.TargetPartition)},
		Key:            []byte(event.EventKey),
		Value:          []byte(event.EventData)}

	err := producer.Produce(&kafkaMessage, delivery_chan)

	if err != nil {
		fmt.Println(err.Error())
	}

	e := <-delivery_chan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	}
	// else {
	// 	fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
	// 		*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	// }
	close(delivery_chan)

	return nil

}
