FROM golang:1.22.1 as builder

RUN apt-get update && \
    apt-get install -y git ca-certificates chromium && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o entrypoint

FROM gcr.io/distroless/base-debian11

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/entrypoint /entrypoint
COPY --from=builder /usr/bin/chromium /usr/bin/chromium
COPY --from=builder /usr/lib/chromium /usr/lib/chromium

ENV CHROME_PATH=/usr/bin/chromium

EXPOSE 8080

ENTRYPOINT ["/entrypoint"]
