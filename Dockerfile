# Stage 1 – Build the Go binary
FROM golang:1.23 AS builder

WORKDIR /app

# Copy code
COPY . .

# Ensure Go mod is clean
RUN go mod tidy

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o gosecops main.go

# Stage 2 – Minimal runtime container
FROM alpine:latest

# Copy binary from builder
COPY --from=builder /app/gosecops /gosecops

# Ensure binary is executable
RUN chmod +x /gosecops

EXPOSE 8181

ENTRYPOINT ["/gosecops"]
