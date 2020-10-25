package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

var URI = ""

func main() {
	c, err := redis.Dial("tcp", URI)
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
