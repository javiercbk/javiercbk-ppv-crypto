const PPVEvent = artifacts.require("PPVEvent");

const BN_ZERO = web3.utils.toBN(0);
const ONE_ETH = web3.utils.toBN(web3.utils.toWei("1", "ether"), 10);
const BN_PPV_PRICE = ONE_ETH;

const _deployPPVContract = function(startDate, endDate, BN_PPV_PRICE) {
  // using new instead of deploy to use a new contract instance on every test run
  return PPVEvent.new(
    Math.floor(startDate.getTime() / 1000),
    Math.floor(endDate.getTime() / 1000),
    BN_PPV_PRICE
  );
};

const _deployValidPPVContract = function(price) {
  const startDate = new Date();
  const endDate = new Date();
  endDate.setDate(startDate.getDate() + 1); // add 1 day
  return _deployPPVContract(startDate, endDate, price);
};

const _depleteAccount = async function(account, beneficiary) {
  let balance = await web3.eth.getBalance(account);
  balance = web3.utils.toBN(balance, 10);
  const tx = {
    from: account,
    to: beneficiary,
    value: balance
  };
  const txResult = await web3.eth.sendTransaction(tx);
  return txResult;
};

const signPayment = async function(amount, invoiceId, contractAddress, signer) {
  web3.utils.soliditySha3;
  const hash = web3.utils.soliditySha3(
    { t: "uint256", v: amount },
    { t: "bytes32", v: invoiceId },
    { t: "address", v: contractAddress }
  );
  // accounts[0] is the owner, it will sign the message if no signerAccount is provided

  const signed = await web3.eth.accounts.sign(
    hash,
    // FIXME: generate private key in a better way
    "0x2679439b7ac1a10a6e380aeb4b8dfba785ac9bfa4dc7f7b5f29dd2b9f12e50fd"
  );
  return signed.signature;
};

contract("PPVEvent", accounts => {
  it("it should allow the owner to change the event price", async () => {
    const newPrice = web3.utils.toBN(web3.utils.toWei("2", "ether"), 10);
    const ppvEventInstance = await _deployValidPPVContract(BN_PPV_PRICE);
    let ppvEventPrice = await ppvEventInstance.ppvEventPrice();
    ppvEventPrice = web3.utils.toBN(ppvEventPrice, 10);
    assert(
      ppvEventPrice,
      BN_PPV_PRICE,
      "price should match the original price"
    );
    const changePriceResponse = await ppvEventInstance.changePrice(newPrice);
    assert(
      changePriceResponse,
      "change price should be a non null transaction"
    );
    assert(
      changePriceResponse.tx,
      "change price should be a non null transaction"
    );
    ppvEventPrice = await ppvEventInstance.ppvEventPrice();
    assert(
      web3.utils.toBN(ppvEventPrice, 10).eq(newPrice),
      "price should match the new price"
    );
    const contractBalance = await web3.eth.getBalance(ppvEventInstance.address);
    assert(
      web3.utils.toBN(contractBalance, 10).eq(BN_ZERO),
      "balance should be zero"
    );
  });

  it("it should deny non contract owners to change the event price", async () => {
    const newPrice = web3.utils.toBN(web3.utils.toWei("2", "ether"), 10);
    const ppvEventInstance = await _deployValidPPVContract(BN_PPV_PRICE);
    // accounts[0] is the contract owner
    const nonOwnerAccount = accounts[1];
    let exceptionThrown = null;
    try {
      await ppvEventInstance.changePrice(newPrice, { from: nonOwnerAccount });
    } catch (e) {
      exceptionThrown = e;
    }
    assert(exceptionThrown !== null, "exception should have been thrown");
    assert.equal(
      exceptionThrown.reason,
      "only the contract owner can call function",
      "exception reason must match"
    );
  });

  it("it should allow non owners to subscribe", async () => {
    const ppvEventInstance = await _deployValidPPVContract(BN_PPV_PRICE);
    const invoiceId = web3.utils.asciiToHex("123");
    const signature = await signPayment(
      BN_PPV_PRICE,
      invoiceId,
      ppvEventInstance.address,
      accounts[0]
    );
    // accounts[0] is the contract owner
    const nonOwnerAccount = accounts[1];
    const subscribeResponse = await ppvEventInstance.subscribe(
      invoiceId,
      signature,
      {
        from: nonOwnerAccount,
        value: BN_PPV_PRICE
      }
    );
    assert(
      subscribeResponse,
      "subscribe response should be a non null transaction"
    );
    assert(
      subscribeResponse.tx,
      "subscribe response should be a non null transaction"
    );
    const contractBalance = await web3.eth.getBalance(ppvEventInstance.address);
    assert(
      web3.utils.toBN(contractBalance, 10).eq(BN_PPV_PRICE),
      "balance should be 1 eth"
    );
  });

  it("it should deny non owners to subscribe more than once", async () => {
    const ppvEventInstance = await _deployValidPPVContract(BN_PPV_PRICE);
    // accounts[0] is the contract owner
    const nonOwnerAccount = accounts[1];
    const invoiceId = web3.utils.asciiToHex("123");
    const signature = await signPayment(
      BN_PPV_PRICE,
      invoiceId,
      ppvEventInstance.address,
      accounts[0]
    );
    const subscribeResponse = await ppvEventInstance.subscribe(
      invoiceId,
      signature,
      {
        from: nonOwnerAccount,
        value: BN_PPV_PRICE
      }
    );
    assert(
      subscribeResponse,
      "subscribe response should be a non null transaction"
    );
    assert(
      subscribeResponse.tx,
      "subscribe response should be a non null transaction"
    );
    let contractBalance = await web3.eth.getBalance(ppvEventInstance.address);
    assert(
      web3.utils.toBN(contractBalance, 10).eq(BN_PPV_PRICE),
      "balance should be zero"
    );
    let exceptionThrown = null;
    try {
      await ppvEventInstance.subscribe(invoiceId, signature, {
        from: nonOwnerAccount,
        value: BN_PPV_PRICE
      });
    } catch (e) {
      exceptionThrown = e;
    }
    assert(exceptionThrown !== null, "exception should have been thrown");
    assert.equal(
      exceptionThrown.reason,
      "the invoice has already been payed",
      "exception reason must match"
    );
    contractBalance = await web3.eth.getBalance(ppvEventInstance.address);
    assert(
      web3.utils.toBN(contractBalance, 10).eq(BN_PPV_PRICE),
      "balance should be zero"
    );
  });

  it("it should deny non owners to subscribe with non matching invoiceId", async () => {
    const ppvEventInstance = await _deployValidPPVContract(BN_PPV_PRICE);
    // accounts[0] is the contract owner
    const nonOwnerAccount = accounts[1];
    const invoiceId = web3.utils.asciiToHex("123");
    const signature = await signPayment(
      BN_PPV_PRICE,
      invoiceId,
      ppvEventInstance.address,
      accounts[0]
    );
    let exceptionThrown = null;
    const otherInvoiceId = web3.utils.asciiToHex("124");
    try {
      await ppvEventInstance.subscribe(otherInvoiceId, signature, {
        from: nonOwnerAccount,
        value: BN_PPV_PRICE
      });
    } catch (e) {
      exceptionThrown = e;
    }
    assert(exceptionThrown !== null, "exception should have been thrown");
    assert.equal(
      exceptionThrown.reason,
      "message was not signed by authorized party",
      "exception reason must match"
    );
    contractBalance = await web3.eth.getBalance(ppvEventInstance.address);
    assert(
      web3.utils.toBN(contractBalance, 10).eq(BN_PPV_PRICE),
      "balance should be zero"
    );
  });

  it("it should deny non owners to subscribe with non matching signer", async () => {
    const ppvEventInstance = await _deployValidPPVContract(BN_PPV_PRICE);
    // accounts[0] is the contract owner
    const nonOwnerAccount = accounts[1];
    const invoiceId = web3.utils.asciiToHex("123");
    const signature = await signPayment(
      BN_PPV_PRICE,
      invoiceId,
      ppvEventInstance.address,
      accounts[1]
    );
    let exceptionThrown = null;
    try {
      await ppvEventInstance.subscribe(invoiceId, signature, {
        from: nonOwnerAccount,
        value: BN_PPV_PRICE
      });
    } catch (e) {
      exceptionThrown = e;
    }
    assert(exceptionThrown !== null, "exception should have been thrown");
    assert.equal(
      exceptionThrown.reason,
      "message was not signed by authorized party",
      "exception reason must match"
    );
    contractBalance = await web3.eth.getBalance(ppvEventInstance.address);
    assert(
      web3.utils.toBN(contractBalance, 10).eq(BN_PPV_PRICE),
      "balance should be zero"
    );
  });

  it("it should deny non owners to subscribe with non matching value", async () => {
    const lesserPrice = web3.utils.toBN(web3.utils.toWei("0.5", "ether"), 10);
    const greaterPrice = web3.utils.toBN(web3.utils.toWei("2", "ether"), 10);
    const ppvEventInstance = await _deployValidPPVContract(BN_PPV_PRICE);
    // accounts[0] is the contract owner
    const nonOwnerAccount = accounts[1];
    let exceptionThrown = null;
    const invoiceId = web3.utils.asciiToHex("123");
    const signature = await signPayment(
      BN_PPV_PRICE,
      invoiceId,
      ppvEventInstance.address,
      accounts[0]
    );
    try {
      await ppvEventInstance.subscribe(invoiceId, signature, {
        from: nonOwnerAccount,
        value: lesserPrice
      });
    } catch (e) {
      exceptionThrown = e;
    }
    assert(exceptionThrown !== null, "exception should have been thrown");
    assert.equal(
      exceptionThrown.reason,
      "value does not match event price",
      "exception reason must match"
    );
    exceptionThrown = null;
    try {
      await ppvEventInstance.subscribe(invoiceId, signature, {
        from: nonOwnerAccount,
        value: greaterPrice
      });
    } catch (e) {
      exceptionThrown = e;
    }
    assert(exceptionThrown !== null, "exception should have been thrown");
    assert.equal(
      exceptionThrown.reason,
      "value does not match event price",
      "exception reason must match"
    );
    const contractBalance = await web3.eth.getBalance(ppvEventInstance.address);
    assert(
      web3.utils.toBN(contractBalance, 10).eq(BN_ZERO),
      "balance should be zero"
    );
  });

  it("it should deny non owners to subscribe without sufficient funds", async () => {
    const ppvEventInstance = await _deployValidPPVContract(BN_PPV_PRICE);
    // accounts[0] is the contract owner
    const nonOwnerAccount = accounts[1];
    const otherNonOwnerAccount = accounts[9];
    await _depleteAccount(otherNonOwnerAccount, nonOwnerAccount);
    const invoiceId = web3.utils.asciiToHex("123");
    const signature = await signPayment(
      BN_PPV_PRICE,
      invoiceId,
      ppvEventInstance.address,
      accounts[0]
    );
    try {
      await ppvEventInstance.subscribe(invoice, signature, {
        from: otherNonOwnerAccount,
        value: BN_PPV_PRICE
      });
    } catch (e) {
      exceptionThrown = e;
    }
    assert(exceptionThrown !== null, "exception should have been thrown");
    assert(
      /sender doesn\'t have enough funds to send tx.*/.test(
        exceptionThrown.message
      ),
      "exception reason must match"
    );
    const contractBalance = await web3.eth.getBalance(ppvEventInstance.address);
    assert(
      web3.utils.toBN(contractBalance, 10).eq(BN_ZERO),
      "balance should be zero"
    );
  });

  it("it should allow a subscribed address to unsubscribing", async () => {
    const startDate = new Date();
    startDate.setDate(startDate.getDate() + 1); // add 1 day
    const endDate = new Date();
    endDate.setDate(startDate.getDate() + 2); // add 2 day
    const ppvEventInstance = await _deployPPVContract(
      startDate,
      endDate,
      BN_PPV_PRICE
    );
    // accounts[0] is the contract owner
    const nonOwnerAccount = accounts[1];
    const invoiceId = web3.utils.asciiToHex("123");
    const signature = await signPayment(
      BN_PPV_PRICE,
      invoiceId,
      ppvEventInstance.address,
      accounts[0]
    );
    const subscribeResponse = await ppvEventInstance.subscribe(
      invoiceId,
      signature,
      {
        from: nonOwnerAccount,
        value: BN_PPV_PRICE
      }
    );
    assert(
      subscribeResponse,
      "subscribe response should be a non null transaction"
    );
    assert(
      subscribeResponse.tx,
      "subscribe response should be a non null transaction"
    );
    let contractBalance = await web3.eth.getBalance(ppvEventInstance.address);
    assert(
      web3.utils.toBN(contractBalance, 10).eq(BN_PPV_PRICE),
      "balance should be 1 eth"
    );
    const unsubscribeResponse = await ppvEventInstance.unsubscribe(invoiceId, {
      from: nonOwnerAccount
    });
    assert(
      unsubscribeResponse,
      "unsubscribe response should be a non null transaction"
    );
    assert(
      unsubscribeResponse.tx,
      "unsubscribe response should be a non null transaction"
    );
    contractBalance = await web3.eth.getBalance(ppvEventInstance.address);
    assert(
      web3.utils.toBN(contractBalance, 10).eq(BN_ZERO),
      "balance should be zero"
    );
  });

  it("it should do nothing when unsubscribing with no subscription", async () => {
    const startDate = new Date();
    startDate.setDate(startDate.getDate() + 1); // add 1 day
    const endDate = new Date();
    endDate.setDate(startDate.getDate() + 2); // add 2 day
    const ppvEventInstance = await _deployPPVContract(
      startDate,
      endDate,
      BN_PPV_PRICE
    );
    // accounts[0] is the contract owner
    const nonOwnerAccount = accounts[1];
    const invoiceId = web3.utils.asciiToHex("123");
    const signature = await signPayment(
      BN_PPV_PRICE,
      invoiceId,
      ppvEventInstance.address,
      accounts[0]
    );
    const subscribeResponse = await ppvEventInstance.subscribe(
      invoiceId,
      signature,
      {
        from: nonOwnerAccount,
        value: BN_PPV_PRICE
      }
    );
    assert(
      subscribeResponse,
      "subscribe response should be a non null transaction"
    );
    assert(
      subscribeResponse.tx,
      "subscribe response should be a non null transaction"
    );
    const unexistingInvoiceId = "234";
    let exceptionThrown = null;
    try {
      await ppvEventInstance.unsubscribe(unexistingInvoiceId, {
        from: nonOwnerAccount
      });
    } catch (e) {
      exceptionThrown = e;
    }
    assert(exceptionThrown !== null, "exception should have been thrown");
    assert.equal(
      exceptionThrown.reason,
      "no subscription available",
      "exception reason must match"
    );
  });

  it("it should not allow unsubscribing other sender", async () => {
    const startDate = new Date();
    startDate.setDate(startDate.getDate() + 1); // add 1 day
    const endDate = new Date();
    endDate.setDate(startDate.getDate() + 2); // add 2 day
    const ppvEventInstance = await _deployPPVContract(
      startDate,
      endDate,
      BN_PPV_PRICE
    );
    // accounts[0] is the contract owner
    const nonOwnerAccount = accounts[1];
    const otherNonOwnerAccount = accounts[2];
    const invoiceId = web3.utils.asciiToHex("123");
    const signature = await signPayment(
      BN_PPV_PRICE,
      invoiceId,
      ppvEventInstance.address,
      accounts[0]
    );
    const subscribeResponse = await ppvEventInstance.subscribe(
      invoiceId,
      signature,
      {
        from: nonOwnerAccount,
        value: BN_PPV_PRICE
      }
    );
    assert(
      subscribeResponse,
      "subscribe response should be a non null transaction"
    );
    assert(
      subscribeResponse.tx,
      "subscribe response should be a non null transaction"
    );
    let exceptionThrown = null;
    try {
      await ppvEventInstance.unsubscribe(invoiceId, {
        from: otherNonOwnerAccount
      });
    } catch (e) {
      exceptionThrown = e;
    }
    assert(exceptionThrown !== null, "exception should have been thrown");
    assert.equal(
      exceptionThrown.reason,
      "subscriber does not match sender",
      "exception reason must match"
    );
  });

  it("it should allow the contract owner end the event", async () => {
    const startDate = new Date();
    startDate.setDate(startDate.getDate() - 2); // add 1 day
    const endDate = new Date();
    endDate.setDate(startDate.getDate() - 1); // add 2 day
    const ppvEventInstance = await _deployPPVContract(
      startDate,
      endDate,
      BN_PPV_PRICE
    );
    // accounts[0] is the contract owner
    const nonOwnerAccount = accounts[1];
    const otherAddress = accounts[2];
    const invoiceId = web3.utils.asciiToHex("123");
    const signature = await signPayment(
      BN_PPV_PRICE,
      invoiceId,
      ppvEventInstance.address,
      accounts[0]
    );
    const subscribeResponse = await ppvEventInstance.subscribe(
      invoiceId,
      signature,
      {
        from: nonOwnerAccount,
        value: BN_PPV_PRICE
      }
    );
    assert(
      subscribeResponse,
      "subscribe response should be a non null transaction"
    );
    assert(
      subscribeResponse.tx,
      "subscribe response should be a non null transaction"
    );
    let contractBalance = await web3.eth.getBalance(ppvEventInstance.address);
    assert(
      web3.utils.toBN(contractBalance, 10).eq(BN_PPV_PRICE),
      "balance should be 1 eth"
    );
    let otherAddressBalance = await web3.eth.getBalance(otherAddress);
    otherAddressBalance = web3.utils.toBN(otherAddressBalance, 10);
    await ppvEventInstance.eventEnd(otherAddress);
    contractBalance = await web3.eth.getBalance(ppvEventInstance.address);
    assert(
      web3.utils.toBN(contractBalance, 10).eq(BN_ZERO),
      "balance should be zero"
    );
    let newOtherAddressBalance = await web3.eth.getBalance(otherAddress);
    newOtherAddressBalance = web3.utils.toBN(newOtherAddressBalance, 10);
    const balanceDiff = newOtherAddressBalance.sub(otherAddressBalance);
    assert(ONE_ETH.eq(balanceDiff), "balance should have increased by 1 eth");
  });

  it("it should not allow a non owner to end the event", async () => {
    const ppvEventInstance = await _deployValidPPVContract(BN_PPV_PRICE);
    // accounts[0] is the contract owner
    const nonOwnerAccount = accounts[1];
    const otherAddress = accounts[2];
    let exceptionThrown = null;
    try {
      await ppvEventInstance.eventEnd(otherAddress, { from: nonOwnerAccount });
    } catch (e) {
      exceptionThrown = e;
    }
    assert(exceptionThrown !== null, "exception should have been thrown");
    assert.equal(
      exceptionThrown.reason,
      "only the contract owner can call function",
      "exception reason must match"
    );
  });

  it("it should deny the contract owner end the event when the event has not started", async () => {
    const startDate = new Date();
    startDate.setDate(startDate.getDate() + 1); // add 2 day
    const endDate = new Date();
    endDate.setDate(startDate.getDate() + 2); // add 2 day
    const ppvEventInstance = await _deployPPVContract(
      startDate,
      endDate,
      BN_PPV_PRICE
    );
    // accounts[0] is the contract owner
    const otherAddress = accounts[2];
    let exceptionThrown = null;
    try {
      await ppvEventInstance.eventEnd(otherAddress);
    } catch (e) {
      exceptionThrown = e;
    }
    assert(exceptionThrown !== null, "exception should have been thrown");
    assert.equal(
      exceptionThrown.reason,
      "ppv event has not started",
      "exception reason must match"
    );
  });

  it("it should deny a subscriber to unsubscribe if the event has started", async () => {
    const startDate = new Date();
    startDate.setDate(startDate.getDate() - 1); // minus 1 day
    const endDate = new Date();
    endDate.setDate(startDate.getDate() + 2); // add 2 day
    const ppvEventInstance = await _deployPPVContract(
      startDate,
      endDate,
      BN_PPV_PRICE
    );
    // accounts[0] is the contract owner
    const nonOwnerAccount = accounts[1];
    const invoiceId = web3.utils.asciiToHex("123");
    const signature = await signPayment(
      BN_PPV_PRICE,
      invoiceId,
      ppvEventInstance.address,
      accounts[0]
    );
    await ppvEventInstance.subscribe(invoiceId, signature, {
      from: nonOwnerAccount,
      value: BN_PPV_PRICE
    });
    let contractBalance = await web3.eth.getBalance(ppvEventInstance.address);
    assert(
      web3.utils.toBN(contractBalance, 10).eq(BN_PPV_PRICE),
      "balance should be 1 eth"
    );
    let exceptionThrown = null;
    try {
      await ppvEventInstance.unsubscribe({ from: nonOwnerAccount });
    } catch (e) {
      exceptionThrown = e;
    }
    assert(exceptionThrown !== null, "exception should have been thrown");
    assert.equal(
      exceptionThrown.reason,
      "ppv event has already started",
      "exception reason must match"
    );
  });
});
