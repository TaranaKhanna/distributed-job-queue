import { Router } from "express";
import { createJobController, getJobByIdController } from "../controllers/jobs.controller.js";

const router = Router();

router.post('/', createJobController);

router.get('/:id', getJobByIdController);

export default router;