FROM golang:1.20.4-alpine3.18

RUN go get -u github.com/swaggo/swag/cmd/swag@1.8

RUN mkdir geek

COPY . /go/src/geek

WORKDIR /go/src/geek/httpserver

RUN go build -o bin/httpserver main.go

ENTRYPOINT [ "./bin/httpserver" ]