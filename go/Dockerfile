FROM golang:1.16-alpine AS build

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/dist .

FROM alpine

COPY --from=build /app/out/dist .
CMD ./dist
