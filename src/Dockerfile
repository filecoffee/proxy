FROM golang:latest

LABEL maintainer="file.coffee <hi@file.coffeer>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 80

CMD ["./main"]