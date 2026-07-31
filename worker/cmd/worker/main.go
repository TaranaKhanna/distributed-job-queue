package main

import (
	"fmt"
	"worker/worker/internal/redis"
)

func greet(name string) string {
	return "Hello " + name
}

func main(){
	client := redis.NewClient()

	fmt.Println("Worker Started!!")
	fmt.Println(client)

	_= redis.Ctx
}