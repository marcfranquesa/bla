FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bla cmd/bla/main.go

FROM alpine:latest

WORKDIR /

COPY --from=builder /app/bla .
COPY --from=builder /app/web/ /web/

CMD ["./bla"]
