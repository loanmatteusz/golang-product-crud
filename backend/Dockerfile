
FROM golang:1.24.2-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o backend ./cmd/api

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/backend .

EXPOSE 3333

CMD ["./backend"]
