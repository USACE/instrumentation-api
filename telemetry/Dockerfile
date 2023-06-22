FROM golang:1.20-alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /go/src/app

ENV CGO_ENABLED=0

COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -mod=vendor -o /go/bin/midas-telemetry

FROM scratch
COPY --from=builder /go/bin/midas-telemetry /go/bin/midas-telemetry
ENTRYPOINT ["/go/bin/midas-telemetry"]
