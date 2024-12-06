# Dockerfile
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o file-registry ./cmd/file-registry

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/file-registry .
COPY --from=builder /app/static ./static

EXPOSE 8090

ENTRYPOINT ["./file-registry"]
