# Build stage
FROM maven:3.8.5-eclipse-temurin-8-alpine AS builder
WORKDIR /app
COPY . .
RUN mvn clean package

# Run stage
FROM amazoncorretto:8u332-alpine3.15-jre
WORKDIR /app
COPY --from=builder /app/target/userInfo-service-1.0-SNAPSHOT.jar .

EXPOSE 9091
CMD [ "java", "-jar", "userInfo-service-1.0-SNAPSHOT.jar" ]