FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install required packages
RUN apk add --no-cache git gcc musl-dev

# Copy go mod files first
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the rest of the code
COPY . .

# Install swag
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.3


# Generate swagger docs
RUN swag init -g cmd/main/main.go --parseDependency --parseInternal

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service cmd/main/main.go


FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY --from=builder /go/bin/swag /usr/local/bin/swag
COPY --from=builder /app/auth-service .
COPY --from=builder /app/docs ./docs
COPY .env .env 

EXPOSE 8080

CMD ["sh", "-c" ,"export $(grep -v '^#' .env | xargs) && ./auth-service"]