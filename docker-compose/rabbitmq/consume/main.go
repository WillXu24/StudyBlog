package main

import (
	rabbitmq2 "docker-compose/rabbitmq/rabbitmq"
)

func main() {
	rabbitmq := rabbitmq2.NewRabbitMQSimple("imoocSimple")
	rabbitmq.ConsumeSimple()
}
