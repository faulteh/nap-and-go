# Use the official Golang image as the build environment
FROM golang:1.23-alpine AS builder

# Install dependencies required to run the Makefile
RUN apk add --no-cache make

# Set the working directory
WORKDIR /app

# Copy the source code and Makefile into the container
COPY . .

# Run the Makefile target to build the bot binary
RUN make build

# Use a minimal base image for the final container
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Install required runtime dependencies
RUN apk add --no-cache ca-certificates

# Copy the binary from the builder
COPY --from=builder /app/bin/bot /app/bin/bot

# Set environment variables for database connection
ENV DB_HOST=postgres
ENV DB_PORT=5432
ENV DB_USER=napandgo
ENV DB_PASSWORD=napandgo
ENV DB_NAME=napandgo
ENV DB_SSLMODE=disable

# Expose any required ports (optional)
EXPOSE 8080

# Run the bot binary
CMD ["/app/bin/bot"]
