# Etapa 1: Construcción del binario
FROM golang:1.22.1 as builder

# Instala las dependencias necesarias
RUN apt-get update && \
    apt-get install -y git ca-certificates chromium && \
    rm -rf /var/lib/apt/lists/*

# Configura el directorio de trabajo
WORKDIR /app

# Copia el código de la aplicación al contenedor
COPY . .

# Descarga las dependencias
RUN go mod tidy

# Compila el binario, optimizado para contenedores
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o entrypoint

# Etapa 2: Imagen mínima para producción
FROM gcr.io/distroless/base-debian11

# Copia los certificados raíz
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copia el binario desde la etapa de construcción
COPY --from=builder /app/entrypoint /entrypoint

# Copia el navegador Chromium
COPY --from=builder /usr/bin/chromium /usr/bin/chromium
COPY --from=builder /usr/lib/chromium /usr/lib/chromium

# Define la variable de entorno para indicar el path del navegador a chromedp
ENV CHROME_PATH=/usr/bin/chromium

# Define el punto de entrada del contenedor
ENTRYPOINT ["/entrypoint"]

