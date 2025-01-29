import express, { Request, Response } from "express";
import request from "supertest";
import { validateMatrixInput } from "./validateMatrix";

const app = express();
app.use(express.json());
app.post("/test-matrix", validateMatrixInput, (req: Request, res: Response) => {
  res.status(200).json({ message: "Validación exitosa" });
});

describe("Middleware validateMatrixInput", () => {
  test("Debe aceptar matrices correctamente formateadas", async () => {
    const response = await request(app)
      .post("/test-matrix")
      .send({
        q: [
          [1, 2],
          [3, 4]
        ],
        r: [
          [5, 6],
          [7, 8]
        ]
      });

    expect(response.status).toBe(200);
    expect(response.body.message).toBe("Validación exitosa");
  });

  test("Debe rechazar si `q` no es un array", async () => {
    const response = await request(app)
      .post("/test-matrix")
      .send({
        q: "no es un array",
        r: [
          [5, 6],
          [7, 8]
        ]
      });

    expect(response.status).toBe(400);
    expect(response.body.error).toBe("Formato inválido");
  });

  test("Debe rechazar si `r` no es un array", async () => {
    const response = await request(app)
      .post("/test-matrix")
      .send({
        q: [
          [1, 2],
          [3, 4]
        ],
        r: "no es un array"
      });

    expect(response.status).toBe(400);
    expect(response.body.error).toBe("Formato inválido");
  });

  test("Debe rechazar si `q` contiene elementos que no son arrays", async () => {
    const response = await request(app)
      .post("/test-matrix")
      .send({
        q: [1, 2, [3, 4]], // 🔹 Elementos que no son arrays
        r: [
          [5, 6],
          [7, 8]
        ]
      });

    expect(response.status).toBe(400);
    expect(response.body.error).toBe("Formato inválido");
  });

  test("Debe rechazar si `r` contiene elementos que no son arrays", async () => {
    const response = await request(app)
      .post("/test-matrix")
      .send({
        q: [
          [1, 2],
          [3, 4]
        ],
        r: [5, [7, 8]] // 🔹 Elementos que no son arrays
      });

    expect(response.status).toBe(400);
    expect(response.body.error).toBe("Formato inválido");
  });

  test("Debe rechazar si `q` contiene valores no numéricos", async () => {
    const response = await request(app)
      .post("/test-matrix")
      .send({
        q: [
          ["a", 2], // 🔹 Valor no numérico
          [3, 4]
        ],
        r: [
          [5, 6],
          [7, 8]
        ]
      });

    expect(response.status).toBe(400);
    expect(response.body.error).toBe("Formato inválido");
  });

  test("Debe rechazar si `r` contiene valores no numéricos", async () => {
    const response = await request(app)
      .post("/test-matrix")
      .send({
        q: [
          [1, 2],
          [3, 4]
        ],
        r: [
          [5, "b"], // 🔹 Valor no numérico
          [7, 8]
        ]
      });

    expect(response.status).toBe(400);
    expect(response.body.error).toBe("Formato inválido");
  });
});
