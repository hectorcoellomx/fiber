# Establecer la imagen base
FROM golang:1.20-alpine

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar el codigo fuente de la app al contenedor
COPY . .

# Compilar la app
RUN go build -o main .

# Puerto en el que se ejecuta la app
EXPOSE 8080

# Comando para ejecutar la app
CMD ["./main"]