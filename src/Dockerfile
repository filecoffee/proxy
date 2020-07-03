FROM golang:latest

LABEL maintainer="file.coffee <hi@file.coffee>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]