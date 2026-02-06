FROM golang:tip-trixie

COPY mailer-service /go
COPY api-view-helpers /api-view-helpers

CMD ["go", "build", "-o", "compiled_app"]
