FROM ubuntu:18.04

ARG USER_NAME

RUN apt-get update && \
    apt-get install -y software-properties-common git wget \
    iproute2 iputils-ping network-manager lsof ufw

WORKDIR /tmp
RUN wget https://bitcoincore.org/bin/bitcoin-core-0.17.1/bitcoin-0.17.1-x86_64-linux-gnu.tar.gz && \
tar -zxvf bitcoin-0.17.1-x86_64-linux-gnu.tar.gz && cd bitcoin-0.17.1/bin/ && cp * /usr/bin/

RUN mkdir -p /home/${USER_NAME}/.bitcoin
RUN groupadd -r ${USER_NAME} && \
    useradd -r -g ${USER_NAME} -d /home/${USER_NAME} -s /sbin/nologin -c "Docker image user" ${USER_NAME}

ENV HOME=/home/${USER_NAME}
USER ${USER_NAME}

WORKDIR /home/${USER_NAME}
