# syntax=docker/dockerfile:1

# Base Image
FROM golang:1.20

WORKDIR /app

# Install project dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of source code files
COPY . ./

# Build project
RUN CGO_ENABLED=0 GOOS=linux go build -o build ./cmd/... && go build ./internal/...

# Other configs
VOLUME /app/logs

# Run application
CMD ["./build/bot-srv"]