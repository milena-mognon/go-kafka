# Projeto GO Kafka

Projeto desenvolvido em Golang para testar e aprender a utilizar o Kafka

## Comando

Para subir o projeto execute esse comando na raiz. Tudo que é necessário para que o projeto funcione será configurado

`docker-compose up -d` 

É preciso criar um tópico e um consumer para receber e enviar as mensagens

Acessa um novo termial e executa os seguintes comandos:

Entra no container do kafka

`docker exec -it <nome_container_kafka> bash`

Cria um tópico

`kafka-topics --create --bootstrap-server=localhost:9092 --topic=teste --partitions=3`

Cria um consumer

`kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste`

Acessa outro terminal e executa a aplicação go 

Acessa a aplicação go

`docker exec -it <nome_container_go> bash`

Executa a aplicação (producer)

`go run cmd/producer/main.go`

Acessa outro terminal e executa a aplicação go

Executa a aplicação (consumer)

`go run cmd/consumer/main.go`