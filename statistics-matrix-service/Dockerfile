# Etapa de construcción
FROM node:18 AS builder
WORKDIR /app

# Copiar archivos de dependencias y paquetes
COPY package*.json ./

# Instalar todas las dependencias (incluye `devDependencies` en esta etapa)
RUN npm install

# Copiar código fuente
COPY . .

# Compilar TypeScript
RUN npm run build

# Etapa de producción
FROM node:18-alpine
WORKDIR /app

# Copiar archivos compilados
COPY --from=builder /app/dist ./dist
COPY package*.json ./

# Asegurar que las dependencias de producción incluyan `mathjs`
RUN npm install --production

# Exponer el puerto del backend Node.js
EXPOSE 3000

# Comando de inicio
CMD ["node", "dist/server.js"]