FROM golang:latest

WORKDIR /go/src/app
COPY . .

RUN chmod 777 /go/src/app/server

ENTRYPOINT ["./server"]