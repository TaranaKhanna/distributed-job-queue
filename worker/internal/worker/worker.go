package worker

import(
	"fmt"
	goredis "github.com/redis/go-redis/v9"
	"worker/internal/redis"
)

func ProcessOneJob(client *goredis.Client) {
	result, err := client.BRPop(redis.Ctx, 0, "jobs:queue").Result()

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(result)

}