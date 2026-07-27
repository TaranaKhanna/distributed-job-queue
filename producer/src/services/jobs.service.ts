import { v4 as uuid } from "uuid";
import redis from "../redis.js";
import { Job, CreateJobInput } from "../types/job.js";

export async function createJob(data: CreateJobInput): Promise<Job> {
    const job: Job = {
        id: uuid(),
        type: data.type,
        status: "pending",
        payload: data.payload,
        createdAt: Date.now()
    }

    await redis.hset(`jobs:${job.id}`, {
        id: job.id,
        type: job.type,
        status: job.status,
        payload: JSON.stringify(job.payload),
        createdAt: job.createdAt.toString()
    })

    await redis.lpush("jobs:queue", job.id);

    return job;
}

export async function getJobById(jobId: string): Promise<Job | null> {
    const jobData = await redis.hgetall(`jobs:${jobId}`);

    if(Object.keys(jobData).length === 0){
        return null;
    }

    return {
        id: jobData.id,
        type: jobData.type,
        payload: JSON.parse(jobData.payload),
        status: jobData.status as Job["status"],
        createdAt: Number(jobData.createdAt)
    };
}