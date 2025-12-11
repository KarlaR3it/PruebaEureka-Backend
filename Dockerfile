# Build stage
FROM golang:alpine AS builder

WORKDIR /app

# Copiar archivos de dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Run stage
FROM alpine:latest

WORKDIR /root/

# Copiar el binario compilado
COPY --from=builder /app/main .
COPY .env .env

# Exponer el puerto (ajusta si usas otro)
EXPOSE 8084

# Comando para ejecutar la aplicación
CMD ["./main"]
