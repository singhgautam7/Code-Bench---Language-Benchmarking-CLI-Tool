FROM eclipse-temurin:22-jdk
RUN apt-get update && apt-get install -y time && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY Main.java .
RUN java -version 2>&1 | head -n 1 > /version.txt
RUN javac Main.java
CMD ["java", "Main"]
