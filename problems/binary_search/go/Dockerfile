FROM golang:1.22 AS builder
WORKDIR /app
COPY main.go .
RUN go build -o app main.go

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]
