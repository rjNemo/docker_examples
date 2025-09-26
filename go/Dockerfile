FROM golang:1.25-alpine AS build

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

WORKDIR /app
RUN apk -U upgrade --no-cache

COPY go.* ./
RUN go mod download

COPY . ./
RUN go build -ldflags="-s -w" -o ./out/dist .

FROM scratch

COPY --from=build /app/out/dist /dist

EXPOSE 80
CMD ["./dist"]
