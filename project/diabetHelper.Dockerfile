FROM golang:1.21.0
WORKDIR /app
COPY go.mod go.sum ./
WORKDIR /app/diabetHelper
RUN go install github.com/cosmtrek/air@latest
