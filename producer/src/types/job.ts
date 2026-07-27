export interface Job{
    id: string,
    type: string;
    payload: Record<string, unknown>;
    status: "pending" | "processing" | "completed" | "failed";
    createdAt: number;
}

export type CreateJobInput = Pick<Job, "type" | "payload">;