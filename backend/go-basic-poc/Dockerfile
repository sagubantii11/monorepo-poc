# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

WORKDIR /backend

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

# Build the Go application (CGO disabled)
RUN CGO_ENABLED=0 go build -o go-basic-poc ./cmd/main.go

# Stage 2: Create the final image
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /backend/go-basic-poc ./

# Copy .env file if used
COPY .env ./

# Expose the port
EXPOSE 8080

# Run the Go app
CMD ["./go-basic-poc"]