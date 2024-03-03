package models

import (
	skafka "github.com/segmentio/kafka-go"
)

//This struct represents a keyed connection record

type S_kafkaConnection struct {
	Key        string
	Topic      string
	Partition  int
	Connection *skafka.Conn
}
