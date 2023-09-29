FROM golang:1.21.0-alpine3.18 AS diabet-helper-builder
WORKDIR /app
COPY go.mod go.sum ./
COPY proto ./proto
COPY diabetHelper ./diabetHelper
RUN go mod tidy
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./diabetHelper/main ./diabetHelper
WORKDIR /app/diabetHelper
#EXPOSE 1111
#CMD ["./main"]

FROM alpine
WORKDIR /build
COPY --from=diabet-helper-builder /app/diabetHelper/main .
COPY --from=diabet-helper-builder /app/diabetHelper/.env .
COPY --from=diabet-helper-builder /app/diabetHelper/storage/postgres/migration /build/storage/postgres/migration
EXPOSE 1111
CMD ["./main"]