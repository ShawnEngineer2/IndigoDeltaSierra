package svcclient

import (
	"context"
	"fmt"
	"indigodeltasierra/svcclient/models"
	"time"

	skafka "github.com/segmentio/kafka-go"
)

func S_produceKafkaMessage(event *models.EventData, skafkaConnection *skafka.Conn) error {
	//This function produces to the indicated Topic and Partition via the indicated Broker(s)

	//Uses Segementio's Kafka-Go library to produce messages into a Redpanda topic

	//Create a Segmentio Message instance
	eventMsg := skafka.Message{
		Key:   []byte(event.EventKey),
		Value: []byte(event.EventData),
	}

	//Write the message to the queue
	_, err := skafkaConnection.WriteMessages(eventMsg)

	//Verify the write
	if err != nil {
		fmt.Println("failed to write messages:")
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
	conn.SetWriteDeadline(time.Now().Add(10 * time.Hour)) //Set ridiculous 10 HOUR timeout for the write window (ugh)

	return conn, nil
}
