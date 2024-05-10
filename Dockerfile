FROM golang:1.19-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go build -o websocket ./cmd/main.go

LABEL authors="alexander"

CMD ["./websocket"]

