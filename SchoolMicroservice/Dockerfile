FROM golang:1.20.0-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY .env .env

RUN export $(cat .env | grep -v ^# | xargs)

COPY go.mod go.sum ./

RUN go mod download

COPY / .

EXPOSE 8083
EXPOSE 9050

RUN go build -o main .

CMD ["./main"]