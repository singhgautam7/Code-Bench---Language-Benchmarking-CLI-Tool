FROM golang:1.23 AS builder
WORKDIR /app
COPY main.go .
RUN go version > /version.txt
RUN go build -o app main.go

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y time && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /version.txt /version.txt
COPY --from=builder /app/app .
CMD ["./app"]
