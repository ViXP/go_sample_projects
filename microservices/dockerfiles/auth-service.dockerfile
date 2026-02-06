FROM golang:tip-trixie

COPY auth-service /go
COPY api-view-helpers /api-view-helpers

CMD ["go", "build", "-o", "compiled_app"]
