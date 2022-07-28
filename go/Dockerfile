FROM golang:1.18-alpine AS build

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

WORKDIR /app

COPY go.* .
RUN go mod download

COPY . .
RUN go build -o ./out/dist .

FROM alpine:latest

RUN apk upgrade --no-cache

COPY --from=build /app/out/dist .
CMD ./dist
