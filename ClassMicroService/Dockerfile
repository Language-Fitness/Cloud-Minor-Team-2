FROM golang:1.21.0-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY .env .env

RUN export $(cat .env | grep -v ^# | xargs)

COPY go.mod go.sum ./

RUN go mod download

COPY / .

EXPOSE 8082
EXPOSE 9093

RUN go build -o main .

CMD ["./main"]