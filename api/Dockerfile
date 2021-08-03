
FROM golang:alpine as build-env

WORKDIR /usr/src


ENV GOPATH=/go
ENV PATH=/go/bin:$PATH

COPY src/go.mod src/go.sum ./

RUN go mod download
RUN go get -u github.com/cosmtrek/air
# RUN go build main.go
