FROM golang:1.19-alpine AS builder
# Install Git
RUN apk update && apk add --no-cache git
# Copy In Source Code
WORKDIR /go/src/app

# Install dependencies
ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download

# Build
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o /go/bin/instrumentation-api

# SCRATCH IMAGE
FROM scratch
COPY --from=builder /go/bin/instrumentation-api /go/bin/instrumentation-api
ENTRYPOINT ["/go/bin/instrumentation-api"]
