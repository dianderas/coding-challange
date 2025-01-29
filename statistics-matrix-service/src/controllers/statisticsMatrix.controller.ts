import { Request, Response } from "express";
import { processMatrix } from "../services/statisticsMatrix.service";

export const handleProcessMatrix = async (req: Request, res: Response): Promise<void> => {
  try {
    const matrixData = req.body;
    const result = await processMatrix(matrixData);
    res.status(200).json(result);
  } catch (error) {
    res.status(500).json({ error: "Error processing the matrix" });
  }
};