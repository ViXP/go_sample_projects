FROM golang:tip-trixie

COPY frontend /go

CMD ["go", "build", "-o", "compiled_app"]
