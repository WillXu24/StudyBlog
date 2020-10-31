package main

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

var uri = "localhost:6379"
var psw = "1234"

func main() {
	c, err := redis.Dial("tcp", uri, redis.DialPassword(psw))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "name", "will")
	if err != nil {
		log.Println("[Error]", err)
		return
	}

	value, err := redis.String(c.Do("GET", "name"))
	if err != nil {
		log.Println("[Error]", err)
		return
	}

	fmt.Println("[success] name:", value)
}
