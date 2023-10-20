# options: dev | prod
ARG BUILD_ENV=prod
ARG BASE_IMAGE=scratch
ARG GO_VERSION=1.20

FROM golang:${GO_VERSION}-alpine AS builder

ARG GOARCH=amd64
ARG BUILD_ENV

RUN apk update && apk add --no-cache git
WORKDIR /go/src/app

ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download

COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    GOFLAGS=$([ "$BUILD_ENV" = "dev" ] && echo "-gcflags=all=-N -gcflags=all=-l" || echo "-ldflags=-w -ldflags=-s") \
    GOOS=linux GOARCH=${GOARCH} \
    go build -ldflags="-w -s" -o /go/bin ./...

FROM golang:${GO_VERSION}-alpine AS build_dev
ONBUILD RUN go install github.com/cosmtrek/air@latest
ONBUILD RUN go install github.com/go-delve/delve/cmd/dlv@latest

FROM $BASE_IMAGE AS build_prod

FROM build_${BUILD_ENV} AS core
COPY --from=builder /go/bin/core /go/bin/midas-core
ENTRYPOINT ["/go/bin/midas-core"]

FROM build_${BUILD_ENV} AS telemetry
COPY --from=builder /go/bin/telemetry /go/bin/midas-telemetry
ENTRYPOINT ["/go/bin/midas-telemetry"]

FROM build_${BUILD_ENV} AS alert
COPY --from=builder /go/bin/alert /go/bin/midas-alert
ENTRYPOINT ["/go/bin/midas-alert"]

FROM build_${BUILD_ENV} AS dcs-loader
COPY --from=builder /go/bin/dcs-loader /go/bin/midas-dcs-loader
ENTRYPOINT ["/go/bin/midas-dcs-loader"]
