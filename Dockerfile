# workspace (GOPATH) configured at /go
FROM golang:1.20 as builder

#
RUN mkdir -p $GOPATH/src/gitlab.udevs.io/eld/Mini-Tweeter-users-service
WORKDIR $GOPATH/src/gitlab.udevs.io/eld/Mini-Tweeter-users-service

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/Mini-Tweeter-users-service /

FROM alpine
COPY --from=builder Mini-Tweeter-users-service .
RUN mkdir config

ENV ENV_FILE_PATH=/app/.env
RUN apk add --no-cache curl

ENTRYPOINT ["/Mini-Tweeter-users-service"]
