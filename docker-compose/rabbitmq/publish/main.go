package main

import (
	rabbitmq2 "docker-compose/rabbitmq/rabbitmq"
	"fmt"
	"strconv"
	"time"
)

func main() {
	rabbitmq := rabbitmq2.NewRabbitMQSimple("imoocSimple")
	rabbitmq.PublishSimple("sleep")
	var i int
	for {
		rabbitmq.PublishSimple("Hello" + strconv.Itoa(i))
		fmt.Println("Published", i)
		i++
		time.Sleep(3 * time.Second)
	}
}
