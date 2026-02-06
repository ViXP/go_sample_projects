FROM golang:tip-trixie

COPY frontend /go/service

WORKDIR /go/service

RUN go build -o compiled_app

CMD ["./compiled_app"]
