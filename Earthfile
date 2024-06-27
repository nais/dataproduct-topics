VERSION 0.6

ARG GO_VERSION=1.22.4

FROM busybox
LABEL org.opencontainers.image.source = "https://github.com/$EARTHLY_GIT_PROJECT_NAME"

build:
    FROM golang:${GO_VERSION}-alpine3.20
    ENV CGO_ENABLED=0

    WORKDIR /app

    COPY go.mod go.sum .
    RUN go mod download

    COPY . .
    RUN go build -o dataproduct-topics ./cmd/dataproduct-topics

    SAVE ARTIFACT dataproduct-topics

    ARG EARTHLY_GIT_PROJECT_NAME
    ARG IMAGE=ghcr.io/$EARTHLY_GIT_PROJECT_NAME
    SAVE IMAGE --push ${IMAGE}-cache:build

tests:
    LOCALLY
    RUN go test ./...

linkerd-await:
    FROM docker.io/curlimages/curl:latest
    ARG LINKERD_AWAIT_VERSION=v0.2.4
    WORKDIR /tmp
    RUN curl -sSLo /tmp/linkerd-await https://github.com/linkerd/linkerd-await/releases/download/release%2F${LINKERD_AWAIT_VERSION}/linkerd-await-${LINKERD_AWAIT_VERSION}-amd64 && \
        chmod 755 /tmp/linkerd-await
    SAVE ARTIFACT linkerd-await

docker:
    FROM alpine:3.20

    WORKDIR /app

    COPY +linkerd-await/linkerd-await .
    COPY +build/dataproduct-topics .
    COPY start.sh .

    ENTRYPOINT [ "/app/linkerd-await", "--shutdown", "--" ]
    CMD ["/app/start.sh"]

    ARG EARTHLY_GIT_SHORT_HASH
    ARG IMAGE_TAG=$EARTHLY_GIT_SHORT_HASH
    ARG EARTHLY_GIT_PROJECT_NAME
    ARG IMAGE=$EARTHLY_GIT_PROJECT_NAME
    SAVE IMAGE --push ${IMAGE}:${IMAGE_TAG} ${IMAGE}:latest
