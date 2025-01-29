import { Request, Response, NextFunction } from "express";
import { body, validationResult } from "express-validator";

// Middleware para validar la estructura de la matriz de entrada
export const validateMatrixInput = [
  body("q").isArray().withMessage("La matriz Q debe ser un array"),
  body("q.*").isArray().withMessage("Cada fila de Q debe ser un array"),
  body("q.*.*").isNumeric().withMessage("Todos los elementos de Q deben ser números"),

  body("r").isArray().withMessage("La matriz R debe ser un array"),
  body("r.*").isArray().withMessage("Cada fila de R debe ser un array"),
  body("r.*.*").isNumeric().withMessage("Todos los elementos de R deben ser números"),

  async (req: Request, res: Response, next: NextFunction): Promise<void> => {
    const errors = validationResult(req);
    if (!errors.isEmpty()) {
      res.status(400).json({ error: "Formato inválido", details: errors.array() });
    }
    next();
  }
];