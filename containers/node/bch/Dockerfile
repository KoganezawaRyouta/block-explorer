FROM ubuntu:18.04

RUN apt-get update && \
    apt-get install -y software-properties-common git wget \
    iproute2 iputils-ping network-manager lsof ufw

ARG USER_NAME

WORKDIR /tmp
RUN wget https://download.bitcoinabc.org/0.18.6/linux/bitcoin-abc-0.18.6-x86_64-linux-gnu.tar.gz && \
tar -zxvf bitcoin-abc-0.18.6-x86_64-linux-gnu.tar.gz && cd bitcoin-abc-0.18.6/bin/ && cp * /usr/bin/

RUN mkdir -p /home/${USER_NAME}/.bitcoin
RUN groupadd -r ${USER_NAME} && \
    useradd -r -g ${USER_NAME} -d /home/${USER} -s /sbin/nologin -c "Docker image user" ${USER_NAME}

ENV HOME=/home/${USER_NAME}
USER ${USER_NAME}

WORKDIR /home/${USER_NAME}