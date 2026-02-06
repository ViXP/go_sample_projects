FROM golang:tip-trixie

COPY broker-service /go
COPY api-view-helpers /api-view-helpers

CMD ["go", "build", "-o", "compiled_app"]
