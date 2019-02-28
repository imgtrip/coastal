FROM golang:1.10-alpine AS build-env



ADD . /go/src/coastal/
WORKDIR /app
RUN cd /go/src/coastal/cmd/coastal && \
    go build -o coastal && \
    cp coastal /app


FROM alpine

ENV TZ=Asia/Shanghai

RUN apk update && \
    apk add tzdata ca-certificates && \
    update-ca-certificates && \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone && \
    rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=build-env /go/src/coastal/cmd/coastal /app

EXPOSE 9090

ENTRYPOINT ["/app/coastal"]
