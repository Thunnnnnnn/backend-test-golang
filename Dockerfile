FROM golang:1.26.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o api .

FROM debian:bookworm-slim

RUN apt-get update && \
    apt-get install -y ca-certificates && \
    update-ca-certificates


WORKDIR /app

COPY --from=builder /app .

EXPOSE 8080

CMD ["./api"]