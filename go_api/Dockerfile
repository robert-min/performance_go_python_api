FROM golang:1.19.4

RUN mkdir /go/src/app/server

COPY ./server /go/src/app/server
WORKDIR /go/src/app/server

RUN go mod init server && go mod tidy
RUN go build -o server

CMD ["./server"]