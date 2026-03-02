FROM eclipse-temurin:22-jdk

RUN apt-get update && apt-get install -y time && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY Main.java .

RUN javac Main.java

CMD ["java", "Main"]
