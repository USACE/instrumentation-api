FROM golang:1.20-alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /go/src/app

ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download

COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o /go/bin ./...

FROM scratch as core
COPY --from=builder /go/bin/core /go/bin/midas-core
ENTRYPOINT ["/go/bin/midas-core"]

FROM scratch as telemetry
COPY --from=builder /go/bin/telemetry /go/bin/midas-telemetry
ENTRYPOINT ["/go/bin/midas-telemetry"]

FROM scratch as alert
COPY --from=builder /go/bin/alert /go/bin/midas-alert
ENTRYPOINT ["/go/bin/midas-alert"]
