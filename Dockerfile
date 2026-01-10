FROM golang:alpine AS builder

ARG SSH_PRIVATE_KEY
RUN mkdir -p /go/src
ADD . /go/src

#Build Source
WORKDIR /go/src

RUN apk add --no-cache --update; \
    apk add git openssh; \
    apk add tzdata; \
    mkdir -p /root/.ssh; \
    chmod 600 /root/.ssh; \
    echo "${SSH_PRIVATE_KEY}" | tr ',' '\n' > /root/.ssh/id_rsa; \
    chmod 600 /root/.ssh/id_rsa
RUN cat /root/.ssh/id_rsa


RUN git config --global url."git@github.com:".insteadOf "https://github.com/"; \
    export GOPRIVATE=github.com/oceaninov; \
    export GONOPROXY=github.com/oceaninov; \
    export GONOSUMDB=github.com/oceaninov; \
    ssh-keyscan github.com >> /root/.ssh/known_hosts; \
    go mod tidy; \
    ls /go; \
    ls /usr/local/go; \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o main-app main.go

#Final Build Image
FROM alpine:latest
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /go/src/main-app /app/main-app

WORKDIR /app

RUN mkdir params;

ENTRYPOINT ["/app/main-app"]

