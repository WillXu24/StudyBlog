package main

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

var uri = "localhost:6379"
var psw = "123456"

func main() {
	c, err := redis.Dial("tcp", uri, redis.DialPassword(psw))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "key", "value")
	if err != nil {
		log.Println("[Error]", err)
		return
	}

	value, err := redis.String(c.Do("GET", "key"))
	if err != nil {
		log.Println("[Error]", err)
		return
	}

	fmt.Println("[success] key :", value)
}
