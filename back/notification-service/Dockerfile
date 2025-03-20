FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o notification-service cmd/main/main.go

FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/notification-service .
COPY .env .env 

EXPOSE 8081

CMD ["sh", "-c", "export $(grep -v '^#' .env | xargs) && ./notification-service"]