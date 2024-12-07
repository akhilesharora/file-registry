// hardhat.config.js
require("@nomicfoundation/hardhat-ethers");
require("@nomicfoundation/hardhat-chai-matchers");

// Default hardhat account private key if none provided
const DEFAULT_PRIVATE_KEY = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80";

module.exports = {
  networks: {
    localhost: {
      accounts: [process.env.PRIVATE_KEY || DEFAULT_PRIVATE_KEY]
    }
  },
  solidity: "0.8.28"
};
