# Stage 1: Build
FROM golang:1.20-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o fleet-ingestion

# Stage 2: Run
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/fleet-ingestion .
CMD ["./fleet-ingestion"]