package main 

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() { 
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "go-kafka_kafka_1:9092", // server do kafka
		"client.id": "goapp-consumer",
		"group.id": "goapp-group", // grupo de aplicações que irá consumeir a mensagem, é obrigatório
		"auto.offset.reset": "earliest", 
	}

	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Println("error consumer: " , err.Error())
	}

	topics:= []string{"teste"}
	c.SubscribeTopics(topics, nil)

	for {
		msg, err := c.ReadMessage(-1) // se chegar alguma coisa ele le, -1 deixa conectado

		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		}
	}
}

