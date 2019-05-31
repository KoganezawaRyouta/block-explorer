FROM nodejs

RUN npm install -g ganache-cli@6.3.0
RUN mkdir -p /var/ganache

ENV DOCKER true
ENV DATADIR /var/ganache

ENV MNEMONIC parent pioneer quick wheat empty travel body coral arrange nasty empty must turtle dismiss extra
ENV NETWORKID 5777

ENTRYPOINT ["sh", "-c", "ganache-cli --host=0.0.0.0 --secure --unlock 0 --unlock 9 --db=${DATADIR} --mnemonic=\"${MNEMONIC}\" --networkId=${NETWORKID}"]