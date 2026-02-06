FROM golang:tip-trixie

COPY auth-service /go/service
COPY api-view-helpers /go/api-view-helpers

WORKDIR /go/service

RUN go build -o compiled_app

CMD ["./compiled_app"]
