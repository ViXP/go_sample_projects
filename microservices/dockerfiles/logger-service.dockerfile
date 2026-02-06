FROM golang:tip-trixie

COPY logger-service /go
COPY api-view-helpers /api-view-helpers

CMD ["go", "build", "-o", "compiled_app"]
