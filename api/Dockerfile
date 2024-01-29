ARG BASE_IMAGE=scratch
ARG GO_VERSION=1.21

FROM golang:${GO_VERSION}-alpine AS builder

ARG GOARCH=amd64
ARG BUILD_ENV

RUN apk update && apk add --no-cache git ca-certificates
WORKDIR /go/src/app

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

ENV CGO_ENABLED=0

RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY --chown=appgroup:appuser go.* ./
RUN go mod download

COPY --chown=appgroup:appuser . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    GOOS=linux GOARCH=${GOARCH} \
    go build -ldflags="-w -s" -o /go/bin ./...

FROM ${BASE_IMAGE} AS core
COPY --from=builder /go/bin/core /go/bin/midas-core
ENTRYPOINT ["/go/bin/midas-core"]

FROM ${BASE_IMAGE} AS telemetry
COPY --from=builder /go/bin/telemetry /go/bin/midas-telemetry
ENTRYPOINT ["/go/bin/midas-telemetry"]

FROM ${BASE_IMAGE} AS alert
COPY --from=builder /go/bin/alert /go/bin/midas-alert
ENTRYPOINT ["/go/bin/midas-alert"]

FROM ${BASE_IMAGE} AS dcs-loader
COPY --from=builder /go/bin/dcs-loader /go/bin/midas-dcs-loader
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT ["/go/bin/midas-dcs-loader"]
