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
    go build -ldflags="-w -s" -o /go/bin/midas-api

FROM scratch
COPY --from=builder /go/bin/midas-api /go/bin/midas-api
ENTRYPOINT ["/go/bin/midas-api"]
