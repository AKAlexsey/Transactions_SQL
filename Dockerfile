# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /

# Copy go mod and sum files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main.go .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /main .
COPY --from=builder /src/views ./src/views
COPY --from=builder /src/static ./src/static

# Expose port
EXPOSE 8080

RUN ["ls", "-la"]

# Run the application
CMD ["./main"]
