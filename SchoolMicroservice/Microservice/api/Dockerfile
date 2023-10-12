FROM golang:1.21.0-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY api/go.mod api/go.sum ./

RUN go mod download

COPY api/ .

EXPOSE 8080

RUN go build -o main .

CMD ["./main"]