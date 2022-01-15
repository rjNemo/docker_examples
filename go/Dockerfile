FROM golang:1.17-alpine AS build

WORKDIR /app

COPY go.* .
RUN go mod download

COPY . .
RUN go build -o ./out/dist .

FROM alpine

COPY --from=build /app/out/dist .
CMD ./dist
