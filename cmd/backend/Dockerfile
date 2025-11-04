# Dockerfile
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build argument to specify which app to build
RUN go build -o bin/app ./cmd/backend

# Final stage - using latest Alpine
FROM alpine:3.19

# Install security updates and certificates
RUN apk --no-cache add ca-certificates tzdata && \
    apk upgrade

# Create non-root user for security
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/bin/app .
# Copy the folders from the builder stage
COPY --from=builder /app/spec ./spec

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Port argument with default value
EXPOSE 3020

# Set port as environment variable so the app can use it
ENV PORT=3020

CMD ["./app"]
