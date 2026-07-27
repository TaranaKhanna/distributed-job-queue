import { Request, Response } from "express";
import { createJob, getJobById } from "../services/jobs.service.js";

type JobParams = {
    id: string;
}

export async function createJobController(req: Request, res: Response) {
    const { type, payload } = req.body;

    const job = await createJob({
        type,
        payload
    });

    res.status(201).json(job);
}

export async function getJobByIdController(req: Request<JobParams>, res: Response) {
    const id = req.params.id;

    const job = await getJobById(id);

    if (!job) {
        return res.status(404).json({
            message: "Job not found"
        });
    }

    res.json(job);

}