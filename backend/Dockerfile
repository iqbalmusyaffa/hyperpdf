# Build stage
FROM golang:1.24-bookworm AS builder

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/server ./cmd/server

# Final runtime stage with Ghostscript and QPDF
FROM debian:bookworm-slim

# Install Ghostscript, QPDF, and essential runtime dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    ghostscript \
    qpdf \
    ca-certificates \
    tzdata \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Create storage directories
RUN mkdir -p /app/storage/uploads /app/storage/outputs

# Copy compiled binary and migrations
COPY --from=builder /app/server /app/server
COPY --from=builder /app/migrations /app/migrations
COPY --from=builder /app/docs /app/docs

# Environment variables
ENV APP_PORT=8080
ENV APP_ENV=production
ENV STORAGE_UPLOAD_DIR=/app/storage/uploads
ENV STORAGE_OUTPUT_DIR=/app/storage/outputs
ENV COMPRESSOR_ENGINE=ghostscript

EXPOSE 8080

CMD ["/app/server"]
