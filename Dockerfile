FROM golang:1.13.2-alpine3.10

ENV GO111MODULE=on

WORKDIR /go/src/app
COPY . .
EXPOSE 9040

RUN apk add make git gcc libc-dev protobuf

RUN go get github.com/unknwon/bra
RUN go get github.com/golang/protobuf/protoc-gen-go
RUN go get github.com/grpc-ecosystem/grpc-health-probe

CMD bra run