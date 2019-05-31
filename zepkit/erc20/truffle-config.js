const path = require("path");
require('dotenv').config();
// const mnemonic = process.env.MNENOMIC;
const HDWalletProvider = require("truffle-hdwallet-provider");
const mnemonic = "parent pioneer quick wheat empty travel body coral arrange nasty empty must turtle dismiss extra";
const infura_url = "https://ropsten.infura.io/v3/e768b31d63f34e939640c06edaaf2dc1";

// Create your own key for Production environments (https://infura.io/)
// const INFURA_ID = 'e768b31d63f34e939640c06edaaf2dc1';

module.exports = {
  // See <http://truffleframework.com/docs/advanced/configuration>
  // to customize your Truffle configuration!
  networks: {
    development: {
      host: "ganache",
      port: 8545,
      network_id: "*",
    },
    ropsten: {
      provider: function() {
        return new HDWalletProvider(mnemonic, infura_url)
      },
      network_id: '3',
      gas: 4465030,
      gasPrice: 10000000000,
    },
    kovan: {
      provider: function() {
        return new HDWalletProvider(mnemonic, 'https://kovan.infura.io/v3/' + process.env.INFURA_API_KEY)
      },
      network_id: '42',
      gas: 4465030,
      gasPrice: 10000000000,
    },
    rinkeby: {
      provider: () => new HDWalletProvider(process.env.MNENOMIC, "https://rinkeby.infura.io/v3/" + process.env.INFURA_API_KEY),
      network_id: 4,
      gas: 3000000,
      gasPrice: 10000000000
    },
    main: {
      provider: () => new HDWalletProvider(process.env.MNENOMIC, "https://mainnet.infura.io/v3/" + process.env.INFURA_API_KEY),
      network_id: 1,
      gas: 3000000,
      gasPrice: 10000000000
    }
  }
};
