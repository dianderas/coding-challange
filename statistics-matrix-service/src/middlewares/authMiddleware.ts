import { Request, Response, NextFunction } from "express";
import { verifyToken } from "../utils/jwtUtils";

export const authenticateJWT = async (req: Request, res: Response, next: NextFunction): Promise<void> => {
  const authHeader = req.header("Authorization");

  if (!authHeader || !authHeader.startsWith("Bearer ")) {
    res.status(401).json({ error: "Acceso no autorizado, token faltante" });
    return;
  }

  const token = authHeader?.split(" ")[1];

  const decoded = verifyToken(token!);
  if (!decoded) {
    res.status(403).json({ error: "Token inválido o expirado" });
    return;
  }

  next();
};