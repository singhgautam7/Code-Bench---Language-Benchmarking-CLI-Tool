FROM gcc:14-bookworm
RUN apt-get update && apt-get install -y time && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY main.cpp .
RUN g++ --version | head -n 1 > /version.txt
RUN g++ -O2 -std=c++20 main.cpp -o app
CMD ["./app"]
