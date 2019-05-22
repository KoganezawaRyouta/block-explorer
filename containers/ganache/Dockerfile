FROM nodejs

RUN npm install -g ganache-cli
RUN mkdir -p /var/ganache

ENV DOCKER true
ENV DATADIR /var/ganache

ENV MNEMONIC pillows andymori al tomoyuki spitz abc quruli whoops fukurouz air bluehearts highlows
ENV NETWORKID 5777

ENTRYPOINT ["sh", "-c", "ganache-cli --host=0.0.0.0 --db=${DATADIR} --mnemonic=\"${MNEMONIC}\" --networkId=${NETWORKID}"]