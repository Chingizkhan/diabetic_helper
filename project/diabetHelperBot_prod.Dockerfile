FROM golang:1.21.0-alpine3.18 AS diabet-helper-bot-builder
WORKDIR /app
COPY go.mod go.sum ./
COPY proto ./proto
COPY bot ./bot
RUN go mod tidy
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bot/main ./bot
WORKDIR /app/bot
#EXPOSE 1112
#CMD ["./main"]

FROM alpine
WORKDIR /build
COPY --from=diabet-helper-bot-builder /app/bot/main .
COPY --from=diabet-helper-bot-builder /app/bot/.env .
EXPOSE 1112
CMD ["./main"]