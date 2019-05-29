FROM ubuntu:18.04

RUN apt update
RUN yes | apt install nodejs npm git nano

RUN git config --global url.https://github.com/.insteadOf git://github.com/

# Install truffle
RUN npm install -g truffle@5.0.2

# Install latest version of ZOS
RUN npm install -g zos@2.2.0

WORKDIR /go/src/github.com/KoganezawaRyouta/block-explorer/go-dapps
CMD ["sh"]
