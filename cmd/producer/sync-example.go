package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

// Exemplo de cmo trabalhar de forma síncrona

func main() {
	deliveryChan := make(chan kafka.Event)
	producer := NewKafkaProducer()
	Publish("mensagem teste 2", "teste", producer, nil, deliveryChan)

	// essa espera é síncrona, é necessário receber o retorno no canal para continuar a execução
	e := <-deliveryChan
	// mensagem que está dentro do evento (resultado)
	msg := e.(*kafka.Message)
	// imprime o resultado
	if msg.TopicPartition.Error != nil {
		fmt.Println("Erro ao enviar")
	} else {
		// para qual partição a mensagem foi enviada
		fmt.Println("Mensagem enviada:", msg.TopicPartition)
		// Ex de Retorno: teste[1]@1 -> topico teste, partição 1, offset 1
	}

	// Flush serve para dar um tempo para que a mensagem seja enviada,
	// sem isso a aplicaçã go finaliza sem ter enviado
	producer.Flush(1000)
}

// Cria um producer
func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "go-kafka_kafka_1:9092",
	}
	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

// Publica uma mensagem
func Publish(msg string, topic string, producer *kafka.Producer, key[]byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value: []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key: key,
	}
	// o resultado da plucblicação da mensagem é publicado nesse canal (deliveryChan)
	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}