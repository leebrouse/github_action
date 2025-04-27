# syntax=docker/dockerfile:1

FROM golang:1.24.2-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8082

ENTRYPOINT ["./app"]
