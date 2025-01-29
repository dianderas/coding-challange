import { Router } from "express";
import { handleProcessMatrix } from "../controllers/statisticsMatrix.controller";
import { authenticateJWT } from "../middlewares/authMiddleware";
import { validateMatrixInput } from "../middlewares/validateMatrix";

const router = Router();

router.post("/statistics-matrix-process", authenticateJWT, ...validateMatrixInput, handleProcessMatrix)

export default router;