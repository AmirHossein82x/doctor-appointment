# Build stage
FROM golang:1.21-alpine AS builder
# Install dependencies
RUN apk add --no-cache git make

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN make build

# Migration tools
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest