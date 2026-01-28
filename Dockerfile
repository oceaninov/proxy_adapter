FROM golang:alpine AS builder

RUN mkdir -p ~/.ssh
RUN --mount=type=secret,id=mysecret cat /run/secrets/mysecret > ~/.ssh/id_rsa
RUN chmod 600 ~/.ssh/id_rsa

RUN mkdir -p /go/src
ADD . /go/src

#Build Source
WORKDIR /go/src

RUN apk add --no-cache --update; \
    apk add git openssh tzdata curl iputils-ping inetutils-telnet;

RUN git config --global url."git@github.com:".insteadOf "https://github.com/"; \
    export GOPRIVATE=github.com/oceaninov; \
    export GONOPROXY=github.com/oceaninov; \
    export GONOSUMDB=github.com/oceaninov; \
    ssh-keyscan github.com >> /root/.ssh/known_hosts; \
    go mod tidy; \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o main-app main.go

#Final Build Image
FROM alpine:latest

RUN apk update
RUN apk add --no-cache iputils-ping
RUN apk add busybox-extras

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /go/src/main-app /app/main-app

WORKDIR /app

RUN mkdir params;

ENTRYPOINT ["/app/main-app"]

