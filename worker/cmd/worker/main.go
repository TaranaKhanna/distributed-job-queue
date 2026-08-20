package main

import (
	"fmt"
	"worker/internal/redis"
	"worker/internal/worker"
)

func main(){
	client := redis.NewClient()

	fmt.Println("Worker Started!!")
	fmt.Println(client)

	worker.ProcessOneJob(client)
}