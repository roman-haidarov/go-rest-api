FROM golang:1.21-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o go-rest-api cmd/main.go

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/go-rest-api .

COPY .env .env

ENV PORT=8080

EXPOSE 8080

CMD ["./go-rest-api"]
