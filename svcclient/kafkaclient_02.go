package svcclient

import (
	"context"
	"fmt"
	"indigodeltasierra/svcclient/models"
	"time"

	skafka "github.com/segmentio/kafka-go"
)

func S_produceKafkaMessage(event *models.EventData, topic string, brokers string) error {
	//This function produces to the indicated Topic and Partition via the indicated Broker(s)

	//Uses Segementio's Kafka-Go library to produce messages into a Redpanda topic

	//Create a Segmentio Message instance
	eventMsg := skafka.Message{
		Key:   []byte(event.EventKey),
		Value: []byte(event.EventData),
	}

	//Create new Kafka Connection
	conn, err := skafka.DialLeader(context.Background(), "tcp", brokers, topic, event.TargetPartition)

	if err != nil {
		fmt.Println("failed to dial leader:")
		fmt.Println(err.Error())
	}

	//Set Write Time out
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	//Write the message to the queue
	_, err = conn.WriteMessages(eventMsg)

	//Verify the write
	if err != nil {
		fmt.Println("failed to write messages:")
		fmt.Println(err.Error())
	}

	//Close the Kafka connection
	if err := conn.Close(); err != nil {
		fmt.Println("failed to close writer:")
		fmt.Println(err.Error())
	}

	return nil
}

func GetSkafkaConnection(topic string, brokers string, partitionID int) (*skafka.Conn, error) {
	//This routine returns a Segmentio Kafka Connection to the identified broker, topic, and partition
	conn, err := skafka.DialLeader(context.Background(), "tcp", brokers, topic, partitionID)

	if err != nil {
		return conn, err
	}

	//Set Write Time out for the connection
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	return conn, nil
}
