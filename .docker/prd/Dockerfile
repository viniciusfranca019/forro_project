FROM golang:1.21-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o cmd/main ./cmd

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /server
COPY --from=builder /app/cmd/main /server
COPY --from=builder /app/app/. /server/app
EXPOSE 8080
CMD ["./main"]