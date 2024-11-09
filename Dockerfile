# syntax=docker/dockerfile:1

ARG VERSION="docker"

FROM golang:alpine as builder

ARG VERSION
ARG USE_MIRROR

RUN mkdir -p /app

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN if [ "$USE_MIRROR" = "true" ]; then go env -w GOPROXY=https://goproxy.cn,direct; fi

ENV CGO_ENABLED=0
ENV VERSION=${VERSION}

RUN go mod download

COPY . .
RUN GOOS=linux go build -ldflags="-X main.Version=${VERSION}" -o /bin/app .

FROM alpine
LABEL maintainer="mail@dongfg.com"

ARG VERSION
ARG USE_MIRROR

RUN if [ "$USE_MIRROR" = "true" ]; then sed -i "s@https://dl-cdn.alpinelinux.org/@https://repo.huaweicloud.com/@g" /etc/apk/repositories; fi

RUN apk update && \
    apk add --no-cache tzdata

ENV TZ=Asia/Shanghai
ENV VERSION=${VERSION}

COPY --from=builder /bin/app /bin/app

ENTRYPOINT ["/bin/app"]