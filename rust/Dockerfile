FROM rust:alpine as builder

WORKDIR /app 

COPY . ./

RUN apk -U upgrade --no-cache && \
  apk add --no-cache musl-dev

RUN cargo build --release


FROM alpine:latest

RUN apk -U upgrade --no-cache 

COPY --from=builder app/target/release/example .

EXPOSE 8080

CMD ./example
