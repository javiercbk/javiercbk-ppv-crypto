import Web3 from "web3";

const walletAddressRE = /^0x[0-9a-fA-F]{40}$/;

const connectToEthNetwork = function(networkLocacation: string) {
  // "http://localhost:7545"
  const web3 = new Web3(new Web3.providers.HttpProvider(networkLocacation));
};

export const validateWalletAddress = (addr: string): boolean =>
  walletAddressRE.test(addr);
