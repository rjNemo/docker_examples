FROM gradle:jdk11-alpine as builder
WORKDIR /app/src
ADD . .
RUN ./gradlew installDist

FROM openjdk:19-jdk-alpine

# apt-get upgrade is for vulnerability fix
RUN apk upgrade --no-cache

EXPOSE 8000:8000
COPY --from=builder /app/src/build/install/demo/ /app/
WORKDIR /app/bin
CMD ["./demo"]
