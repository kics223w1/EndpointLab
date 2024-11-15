# Build stage
FROM golang:1.23.3-alpine AS builder
WORKDIR /app
COPY . .

# Build specifically for linux/amd64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o endpointlab main.go

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/endpointlab .

EXPOSE 8080
ENTRYPOINT ["/app/endpointlab"]