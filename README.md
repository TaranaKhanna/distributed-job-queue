### Distributed Job Queue

A mini distributed job queue built with Node.js, TypeScript, Go, and Redis.

## Current Architecture

Client
  ↓
Node.js API
  ↓
Redis
  ├── jobs:queue      → Job IDs
  └── jobs:<job-id>   → Job data
                         ↓
                      Go Worker
                         ↓
              pending → processing → completed
              

---

## Local setup instructions

# 1. Clone the repository

git clone git@github.com:TaranaKhanna/distributed-job-queue.git
cd distributed-job-queue

# 2. Start Redis

From the project root:

docker compose up -d


## 3. Start the producer

Open a new terminal:

cd producer

npm install

Create a `.env` file:

REDIS_URL=redis://localhost:6379
PORT=3000

Start the API:

npm run dev

## 4. Start the Go worker

Open another terminal:

cd worker

go mod download

go run ./cmd/worker

## 5. Create a job / get a job
POST http://localhost:3000/jobs
GET http://localhost:3000/jobs/<job-id>
