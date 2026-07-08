import "dotenv/config";
import express from "express";
import jobsRouter from "./routes/jobs.js";

const app = express();

app.use(express.json());
app.use("/jobs", jobsRouter);

const PORT = process.env.PORT || 3000;

app.listen(PORT, () => {
    console.log(`Server running on port: ${PORT}`);
});