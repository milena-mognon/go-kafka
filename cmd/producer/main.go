package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	deliveryChan := make(chan kafka.Event)
	producer := NewKafkaProducer()
	Publish("tranferiu", "teste", producer, []byte("tranferencia1"), deliveryChan)

	// joga isso em outra thread para o terminal não ficar travado quando rodar a aplicação
	go DeliveryReport(deliveryChan) // deixa a execução assíncrona
	
	producer.Flush(1000)
}

// Cria um producer
func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "go-kafka_kafka_1:9092", // server do kafka
		"delivery.timeout.ms": "0", // tempo máximo de entrega de uma mensagem // 0 = aguarda sem dar erro
		"acks": "all", // 0 - não precisa receber retorno, 1 - lider deve retornar que recebeu, all - lider e followers precisam confirmar o recebimento
		"enable.idempotence": "true", // padrão é false - permite duplicação. true - garante que chegou exatamente uma vez
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

func DeliveryReport(deliveryChan chan kafka.Event) {
	// "loop infinito" - fica aguardando chegar algo no deliveryChan
	for e:= range deliveryChan {
		switch ev := e.(type)  {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Erro ao enviar")
			} else {
				// para qual partição a mensagem foi enviada
				fmt.Println("Mensagem enviada:", ev.TopicPartition)
				// Ex de Retorno: teste[1]@1 -> topico teste, partição 1, offset 1
			}
		}
	}
}