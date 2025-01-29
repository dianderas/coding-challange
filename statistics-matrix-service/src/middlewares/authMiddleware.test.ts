import express, { Request, Response } from "express";
import request from "supertest";
import { authenticateJWT } from "./authMiddleware";

// Configurar Express para pruebas
const app = express();
app.use(express.json());

// Ruta de prueba protegida con el middleware
app.get("/protected", authenticateJWT, (req: Request, res: Response) => {
  res.status(200).json({ message: "Acceso permitido" });
});

// Mockear `verifyToken` para controlar la validación de JWT
jest.mock("../utils/jwtUtils", () => ({
  verifyToken: jest.fn((token: string) => {
    if (token === "valid_token") return { id: 1, username: "testuser" };
    return null;
  }),
}));

describe("Middleware authenticateJWT", () => {
  test("Debe rechazar solicitudes sin token con 401 Unauthorized", async () => {
    const response = await request(app).get("/protected");
    expect(response.status).toBe(401);
    expect(response.body.error).toBe("Acceso no autorizado, token faltante");
  });

  test("Debe rechazar solicitudes con token inválido con 403 Forbidden", async () => {
    const response = await request(app)
      .get("/protected")
      .set("Authorization", "Bearer invalid_token");

    expect(response.status).toBe(403);
    expect(response.body.error).toBe("Token inválido o expirado");
  });

  test("Debe aceptar solicitudes con token válido y pasar a `next()`", async () => {
    const response = await request(app)
      .get("/protected")
      .set("Authorization", "Bearer valid_token");

    expect(response.status).toBe(200);
    expect(response.body.message).toBe("Acceso permitido");
  });
});
