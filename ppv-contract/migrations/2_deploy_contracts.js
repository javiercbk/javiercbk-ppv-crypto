const PPVEvent = artifacts.require("PPVEvent");

const startDate = new Date();
const endDate = new Date();
endDate.setDate(startDate.getDate() + 1); // add 1 day
const price = 12

module.exports = function(deployer) {
  deployer.deploy(PPVEvent, startDate.getDate(), endDate.getDate(), price);
};
