FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o gosecops main.go

FROM alpine:latest
COPY --from=builder /app/gosecops /gosecops
EXPOSE 8080
ENTRYPOINT ["/gosecops"]
