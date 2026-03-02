FROM python:3.13-slim
RUN apt-get update && apt-get install -y time && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY main.py .
RUN python --version > /version.txt
CMD ["python", "main.py"]
