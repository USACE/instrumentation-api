ARG BASE_IMAGE=scratch
ARG GO_VERSION=1.23

FROM golang:${GO_VERSION}-bookworm AS builder

ARG GOARCH=amd64
ARG BUILD_TARGET

WORKDIR /go/src/app

ENV CGO_ENABLED=0

COPY go.* ./
RUN go mod download

COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
	--mount=type=cache,target=/go/pkg \
	GOOS=linux GOARCH=${GOARCH} \
	go build -ldflags="-w -s" -o /go/bin/app ./cmd/${BUILD_TARGET}

FROM ${BASE_IMAGE}
ARG BUILD_TAG
ENV BUILD_TAG=${BUILD_TAG}
COPY --from=builder /go/bin/app /go/bin/app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT ["/go/bin/app"]
