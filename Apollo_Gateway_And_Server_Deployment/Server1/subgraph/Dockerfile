FROM golang:1.21.0-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY go.mod go.sum ./
#COPY subgraph/go.mod subgraph/go.sum ./

RUN go mod download

COPY . .
#COPY subgraph/ .

EXPOSE 8080

RUN go build -o main .

CMD ["./main"]