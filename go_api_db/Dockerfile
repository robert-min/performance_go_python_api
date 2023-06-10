FROM golang:1.19.4

ARG EXPOSE_PORT=8080
EXPOSE $EXPOSE_PORT/tcp

USER root


COPY ./server /go/src/server
ARG MOUDLE_PATH=/go/src/server
WORKDIR ${MOUDLE_PATH}

SHELL ["/bin/bash", "-c"]

RUN go mod download && go mod tidy
RUN go build -o server

EXPOSE ${EXPOSE_PORT}

CMD ["./server"]