# Step 1: Modules caching
FROM golang:1.20.1-alpine3.17 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Build stage
FROM golang:1.22-alpine AS builder
# COPY --from=modules /go/pkg /go/pkg

# Copy source code
COPY . /app
WORKDIR /app
# Copy go mod and sum files
# COPY go.mod go.sum /app/

# Download dependencies
# RUN go mod download

# Build the application
# RUN CGO_ENABLED=0 GOOS=linux go build -o main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
# COPY --from=builder ./main .
COPY --from=builder /app/src/views ./src/views
COPY --from=builder /app/src/static ./src/static
COPY --from=builder /bin/app /app


# Expose port
EXPOSE 8080

# RUN ["ls", "-la"]

# Run the application
CMD ["/app/app"]
