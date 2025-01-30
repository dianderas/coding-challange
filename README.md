### **1️⃣ Backend Go - API de Factorización QR**

Este proyecto es un sistema de **factorización QR** y cálculo de estadísticas de matrices.  
Utiliza **Go (Fiber)** como backend principal y **Node.js (Express)** para procesar cálculos adicionales.

## Login

Para poder usar la API antes debes pasar por el servicio authenticate

📍 **Base URL:** `http://localhost:8080`

| Método | Endpoint      | Descripción                                           |
| ------ | ------------- | ----------------------------------------------------- |
| `POST` | `/auth/login` | Autenticacion de usuario para recibir token de acceso |

📌 **Ejemplo de Request (`POST /auth/login`):**

```json
{
  "user": "admin",
  "password": "admin"
}
```

📌 Ejemplo de Respuesta (200 OK):

```json
{
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNzM4NDQ0MzExLCJuYW1lIjoiYWRtaW4ifQ.ou4rdZwfw3NciLcOaMFgSpMwnPLqpwJokPOH6xHfDEI"
  },
  "success": true
}
```

Después de obtener el token, debes incluirlo en el header `Authorization` para llamar a otros endpoints:

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

## Full Matrix process

📍 **Base URL:** `http://localhost:8080`

| Método | Endpoint                  | Descripción                                         |
| ------ | ------------------------- | --------------------------------------------------- |
| `POST` | `/api/fullmatrix-process` | Factoriza una matriz y envía datos a `backend-node` |

📌 **Ejemplo de Request (`POST /api/fullmatrix-process`):**

```json
{
  "values": [
    [1, 2],
    [3, 4],
    [5, 6]
  ]
}
```

📌 Ejemplo de Respuesta (200 OK):

```json
{
  "nodeResult": {
    "average": -1.321855754080024,
    "isDiagonal": false,
    "max": 0.8970852271450604,
    "min": -7.437357441610946,
    "totalSum": -13.21855754080024
  },
  "qrResult": {
    "q": [
      [-0.16903085094570325, 0.8970852271450604],
      [-0.50709255283711, 0.2760262237369414],
      [-0.8451542547285166, -0.34503277967117696]
    ],
    "r": [
      [-5.916079783099615, -7.437357441610946],
      [0, 0.8280786712108259]
    ]
  }
}
```

### 2️⃣ Backend Node - API de Procesamiento de Matrices

Alternativamente puedes usar directamente el api de estadisticas. Tambien requiere token de authentication.
📍 **Base URL:** `http://localhost:3000`

| Método | Endpoint                         | Descripción                        |
| ------ | -------------------------------- | ---------------------------------- |
| `POST` | `/api/statistics-matrix-process` | Calcula estadísticas de una matriz |

```json
{
  "q": [
    [0.4472, 0.8944],
    [-0.8944, 0.4472]
  ],
  "r": [
    [2.2361, 3.5777],
    [0, 0.8944]
  ]
}
```

````
📌 Ejemplo de Respuesta (200 OK):
```json
{
  "max": 3.5777,
  "min": 0,
  "average": 1.2688,
  "totalSum": 5.0752,
  "isDiagonal": false
}
````

### 📌 Cómo Ejecutarlo Localmente con Docker Compose

📌 1️⃣ Clonar el Repositorio

```
git clone https://github.com/dianderas/coding-challange.git
cd coding-challange

```

📌 2️⃣ Configurar las Variables de Entorno (.env)

```
JWT_SECRET=supersecretkey
STATISTICS_SERVICE_URL=http://backend-node:3000

```

📌 3️⃣ Ejecutar docker-compose

```
docker-compose up -d --build
```

✅ Esto levantará backend-go en 8080 y backend-node en 3000.

📌 4️⃣ Verificar los Contenedores Corriendo

```
docker ps
```

📌 5️⃣ Probar la API

```
curl -X POST http://localhost:8080/api/fullmatrix-process \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer TU_TOKEN_AQUI" \
  -d '{
    "values": [[1, 2], [3, 4]]
  }'
```

### 📌 Pruebas Automatizadas

📌 Ejecutar pruebas unitarias en Go:

```
go test ./internal/... -v
```

📌 Ejecutar pruebas unitarias en Node.js:

```
npm test
```

📌 Ejecutar pruebas de integración en Go:

```
go test ./internal/handlers -v
```
