FROM golang:tip-trixie

COPY dispatcher-service /go/service
COPY api-view-helpers /go/api-view-helpers

WORKDIR /go/service

RUN go build -o compiled_app

CMD ["./compiled_app"]
