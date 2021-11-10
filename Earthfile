ARG GO_VERSION=1.17.2
ARG EARTHLY_GIT_PROJECT_NAME
ARG IMAGE=ghcr.io/$EARTHLY_GIT_PROJECT_NAME

FROM busybox
LABEL org.opencontainers.image.source = "https://github.com/$EARTHLY_GIT_PROJECT_NAME"

build:
    FROM golang:${GO_VERSION}-alpine3.14
    ENV CGO_ENABLED=0

    WORKDIR /app

    COPY go.mod go.sum .
    RUN go mod download

    COPY . .
    RUN go build -o dataproduct-topics ./cmd/dataproduct-topics

    SAVE ARTIFACT dataproduct-topics
    SAVE IMAGE --push ${IMAGE}-cache:build

tests:
    LOCALLY
    RUN go test ./...

docker:
    FROM alpine:3.14
    ARG EARTHLY_GIT_SHORT_HASH
    ARG IMAGE_TAG=$EARTHLY_GIT_SHORT_HASH

    WORKDIR /app

    COPY +build/dataproduct-topics .
    COPY start.sh .

    CMD ["/app/start.sh"]

    SAVE IMAGE --push ${IMAGE}:${IMAGE_TAG}
    SAVE IMAGE --push ${IMAGE}:latest
