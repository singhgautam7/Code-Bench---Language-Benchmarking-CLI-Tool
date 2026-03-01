FROM gcc:13 AS builder
WORKDIR /app
COPY main.cpp .
RUN g++ -O2 -o app main.cpp

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]
