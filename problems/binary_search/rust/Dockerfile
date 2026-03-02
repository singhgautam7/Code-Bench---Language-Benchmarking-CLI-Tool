FROM rust:1.80 AS builder

WORKDIR /app
COPY main.rs .

RUN rustc -O main.rs -o app

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y time && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /app/app .

CMD ["./app"]
