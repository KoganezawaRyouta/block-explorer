FROM node:8.9.4-alpine
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh && \
    apk add --no-cache --virtual .gyp python make g++

RUN git config --global url.https://github.com/frozeman/.insteadOf git://github.com/frozeman/

RUN npm install web3
