FROM golang:alpine

# Required because go requires gcc to build
RUN apk add build-base

RUN apk add inotify-tools

RUN echo $GOPATH

# Build Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . /clean_web
RUN mkdir /files

ARG VERSION="4.13.0"

RUN set -x \
    && apk add --no-cache git \
    && git clone --branch "v${VERSION}" --depth 1 --single-branch https://github.com/golang-migrate/migrate /tmp/go-migrate

WORKDIR /tmp/go-migrate

RUN set -x \
    && CGO_ENABLED=0 go build -tags 'mysql' -ldflags="-s -w" -o ./migrate ./cmd/migrate \
    && ./migrate -version

RUN cp /tmp/go-migrate/migrate /usr/bin/migrate

WORKDIR /clean_web

ENV GOFLAGS -buildvcs=false

RUN go mod download

CMD sh /clean_web/docker/run.sh