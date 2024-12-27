# Use the official Golang image as the build environment
FROM golang:1.23-alpine AS builder

# Install dependencies required to run the Makefile
RUN apk add --no-cache make

# Set up Go environment variables
ENV GOPATH=/golang
ENV PATH=$GOPATH/bin:$PATH

# Set the working directory
WORKDIR /app

# Copy the source code and Makefile into the container
COPY . .

# Run the Makefile target to build the bot binary
RUN make build

# Use a minimal base image for the final container
FROM alpine:latest AS deploy

# Set the working directory
WORKDIR /app

# Install required runtime dependencies
RUN apk add --no-cache ca-certificates

# Copy the binary from the builder
COPY --from=builder /app/bin/bot /app/bin/bot
COPY --from=builder /app/bin/web /app/bin/web

# Set environment variables for database connection
ENV GOPATH=/golang
ENV PATH=$GOPATH/bin:$PATH

ENV DB_HOST=postgres
ENV DB_PORT=5432
ENV DB_USER=napandgo
ENV DB_PASSWORD=napandgo
ENV DB_NAME=napandgo
ENV DB_SSLMODE=disable

# Expose any required ports (optional)
EXPOSE 8080

# Run the container as a non-root user
USER nobody

# Run the bot binary
CMD ["/app/bin/bot"]

# Development image we actually use go so we can use hot reloading with air
FROM builder AS dev

# Set environment variables for database connection
ENV GOPATH=/golang
ENV PATH=$GOPATH/bin:$PATH

ENV DB_HOST=postgres
ENV DB_PORT=5432
ENV DB_USER=napandgo
ENV DB_PASSWORD=napandgo
ENV DB_NAME=napandgo
ENV DB_SSLMODE=disable

# Install air for Go for hot reloading
RUN apk add --no-cache git
RUN go install github.com/air-verse/air@latest
RUN git config --global --add safe.directory /app

# Expose any required ports (optional)
EXPOSE 8080

# Run the container as a non-root user
USER root

# Run the bot binary
CMD ["air", "-c", ".air-bot.toml"]
