FROM golang:1.8.3

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get clean

WORKDIR /go/src/ires
COPY ./ext /go/src/ires

# pkg管理用
RUN go get -u github.com/golang/dep/cmd/dep

