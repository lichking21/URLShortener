FROM docker.io/library/golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o url-shortener main.go

FROM docker.io/library/alpine:latest
WORKDIR /app
COPY --from=builder /app/url-shortener .
COPY migrations/ ./migrations/
COPY .env .

EXPOSE 8080
CMD ["./url-shortener"]