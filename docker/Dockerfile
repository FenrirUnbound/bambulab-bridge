FROM golang:1.23.5 AS builder

WORKDIR /opt/relay

COPY . /opt/relay

RUN CGO_ENABLED=0 go build -o blb ./cmd/...

FROM golang:1.23.5 AS app

WORKDIR /opt/relay

COPY --from=builder /opt/relay/blb /opt/relay/blb

ENTRYPOINT ["/opt/relay/blb"]