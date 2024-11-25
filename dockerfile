# Use Golang base image
FROM alpinelinux/golang AS builder

# Set work directory inside the container
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .



# Build the Go app
RUN go build -o main ./cmd/service

# Use a smaller base image for the final container
FROM alpine:latest

# Set work directory
WORKDIR /app

# Copy migration files into the container
COPY ./internal/storage/database/migrations/ /app/internal/storage/database/migrations/

# Copy the Go app binary from the builder
COPY --from=builder /app/main .

# Create log folder
RUN mkdir -p /app/log

# Command to run the app
CMD ["./main"]

# setting defaults for environment
ENV MUSIC_ADDRESS=host.docker.internal:8080 \
    DETAILS_API_ADDRESS=host.docker.internal:8081 \
    MUSIC_DATABASE_STRING="pgx://postgres:postgres@postgres:5432/postgres?sslmode=disable" \
    MUSIC_LOG_FILE="./log/log.json"
