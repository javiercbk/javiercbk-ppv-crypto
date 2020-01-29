import Web3 from "web3";
const contractABI = require("@/lib/abi/ppv_abi.json");

const ethNetwork: string = process.env.VUE_APP_ETH_NETWORK_URL || "";

const web3 = new Web3(ethNetwork || Web3.givenProvider);

const walletAddressRE = /^0x[0-9a-fA-F]{40}$/;

export const validateWalletAddress = (addr: string): boolean =>
  walletAddressRE.test(addr);

export const subscribeToPPV = async function(address: string) {
  const NameContract = new web3.eth.Contract(contractABI, address);
  NameContract.methods.subscribe("subscribe").send();
};
