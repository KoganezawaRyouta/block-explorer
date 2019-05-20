FROM node:11.6.0-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh && \
    apk add --no-cache python make g++

RUN git config --global url.https://github.com/.insteadOf git://github.com/
