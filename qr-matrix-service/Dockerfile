# Usa la imagen oficial de Go
FROM golang:1.23.5 AS builder

# Setea el directorio de trabajo en el contenedor
WORKDIR /app

# Copia los archivos de configuración
COPY go.mod go.sum ./
RUN go mod download

# Copia todo el código fuente
COPY . .

# 🚀 Construye la aplicación desde `cmd/`
RUN GOOS=linux GOARCH=amd64 go build -o main ./cmd

# Segunda etapa: Imagen más ligera para producción
FROM alpine:latest

# Instala certificados raíz
RUN apk --no-cache add ca-certificates libc6-compat

WORKDIR /root/

# Copia el binario desde la etapa anterior
COPY --from=builder /app/main .

# Exponer el puerto
EXPOSE 8080

# Definir variable de entorno (sobreescribible)
ENV JWT_SECRET="default-secret"

RUN chmod +x /root/main

# Comando de inicio
CMD ["./main"]