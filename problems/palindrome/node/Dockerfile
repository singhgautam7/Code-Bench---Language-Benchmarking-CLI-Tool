FROM node:22-slim
RUN apt-get update && apt-get install -y time && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY main.js .
RUN node --version > /version.txt
CMD ["node", "main.js"]
