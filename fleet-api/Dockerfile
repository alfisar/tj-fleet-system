# Stage 1: Build
FROM golang:1.20-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o fleet-api

# Stage 2: Run
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/fleet-api .
EXPOSE 8080
CMD ["./fleet-api"]