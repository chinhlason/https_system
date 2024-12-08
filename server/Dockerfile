FROM golang:1.23-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

COPY ./ssl /app/ssl

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/ssl ./ssl

EXPOSE 2999

CMD ["./main"]