FROM golang:1.21-alpine

COPY . /go/src/app

WORKDIR /go/src/app/cmd/port-service

RUN go build -o bin/app main.go

ENV PORT_SERVICE_HTTP_ADDR=:8080

EXPOSE 8080

CMD ["./bin/app"]