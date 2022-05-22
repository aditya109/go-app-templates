FROM golang:alpine AS base
LABEL org.opencontainers.image.authors="adikid1996@gmail.com"
LABEL maintainer="daitya96"
LABEL version="1.0"
ARG IS_RUNNING_IN_LOCAL
ENV USE_LOCAL_CONFIG=${IS_RUNNING_IN_LOCAL}
ENV GO111MODULE=on
ENV APP_HOME $GOPATH/src/github.com/aditya109/go-server-template
WORKDIR "${APP_HOME}"

COPY . "${APP_HOME}"
RUN go mod download
RUN go build -o app .
EXPOSE 8000
CMD ["./app"]
