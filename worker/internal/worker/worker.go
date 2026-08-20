package worker

import (
	"fmt"
	"time"
	"worker/internal/redis"
	"worker/internal/types"

	goredis "github.com/redis/go-redis/v9"
)

func ProcessOneJob(client *goredis.Client) {
	result, err := client.BRPop(redis.Ctx, 0, "jobs:queue").Result()

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	jobID := result[1]
	jobData, err := client.HGetAll(
		redis.Ctx,
		"jobs:"+jobID,
	).Result()

	if err != nil {
		fmt.Println("Error fetching job: ", err)
		return
	}

	job := types.Job{
		ID:        jobData["id"],
		Type:      jobData["type"],
		Payload:   jobData["payload"],
		Status:    jobData["status"],
		CreatedAt: jobData["createdAt"],
	}
	
	fmt.Println("job: ", job)

	_, err = client.HSet(
		redis.Ctx,
		"jobs:"+jobID,
		"status",
		"processing",
	).Result()

	if err != nil {
		fmt.Println("Error updating job status: ", err)
		return
	}

	fmt.Println("Job status updated to processing")

	fmt.Println("Processing job...")
	time.Sleep(5 * time.Second)

	_, err = client.HSet(
		redis.Ctx,
		"jobs:"+jobID,
		"status",
		"compelted",
	).Result()

	if err != nil {
		fmt.Println("Error updating job status: ", err)
		return
	}

	fmt.Println("job status updated to complete")
}
