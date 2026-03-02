FROM ruby:3.3
RUN apt-get update && apt-get install -y time && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY main.rb .
RUN ruby --version > /version.txt
CMD ["ruby", "main.rb"]
