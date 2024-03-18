FROM golang:1.22.1-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o guilds-service ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/guilds-service .

EXPOSE 8080

CMD ["./guilds-service"]
