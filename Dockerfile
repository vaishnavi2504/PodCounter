FROM golang:1.14.2-alpine

RUN mkdir /app

COPY main.go /app/
COPY vendor/ /app/vendor/
COPY go.mod /app/
COPY go.sum /app/

WORKDIR /app

RUN go build -mod=vendor

CMD ["/app/hello"]