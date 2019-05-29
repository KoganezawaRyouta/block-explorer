FROM ubuntu:18.04

RUN apt update
RUN yes | apt install nodejs npm git nano

RUN git config --global url.https://github.com/.insteadOf git://github.com/

# Install truffle
RUN npm install -g truffle@5.0.2

# Install npx
#RUN npm install -g npx\
# Install ganache-cli
#RUN npm install -g ganache-cli@6.3.0
# Unbox zepkit truffle box
#RUN npx truffle unbox zeppelinos/zepkit

# Install latest version of ZOS
RUN npm install -g zos@2.2.0

WORKDIR /zepkit
CMD ["sh"]
