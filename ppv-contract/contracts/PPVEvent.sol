pragma solidity >=0.4.25 <0.7.0;

/** @title Pay per view event */
contract PPVEvent {
    struct Subscription {
        uint256 amount;
        address subscriber;
    }
    // The keyword "public" makes variables
    // accessible from other contracts
    address private owner;
    // price in wei
    uint256 public ppvEventPrice; // 1 eth = 10000000000 wei => 0.00000000001 eth = 1 wei ... 15000000000
    uint256 public ppvEventStart; // 01-01-1970 00:00:00.000
    uint256 public ppvEventEnd; // 01-01-1970 00:00:00.000
    mapping(bytes32 => Subscription) private subscriptions;
    // Set to true at the end, disallows any change.
    // By default initialized to `false`.
    bool public ended;

    // Event emmited on ppv event started
    event PPVEventStarted();
    // Event emmited on ppv event end
    event PPVEventEnded();
    // Event emmited on price changed
    event PriceChanged(uint256 ppvEventPrice);
    // Event emmited on new subscription
    event NewSubscription(address subscriptor, uint256 price);
    // Event emmited on new unsubscription
    event NewUnsubscription(address subscriptor, uint256 price);

    modifier notEnded(bool _ended) {
        require(!_ended, "ppv event has ended");
        // Do not forget the "_;"! It will
        // be replaced by the actual function
        // body when the modifier is used.
        _;
    }

    modifier beforeStart(uint256 _time) {
        // Interestingly, now - which is equivalent to block.timestamp -
        // may not be as accurate as one may think. It is up to the miner to pick it,
        // so it could be up to 15 minutes (900 seconds) off as explained in the following formula:
        // parent.timestamp >= block.timestamp <= now + 900 seconds
        // As a consequence, now shouldn’t be used for measuring small time units.
        require(now < _time, "ppv event has already started");
        _;
    }

    modifier afterStart(uint256 _time) {
        require(now >= _time, "ppv event has not started");
        _;
    }

    modifier notPayed(bytes32 invoiceId) {
        Subscription memory sub = subscriptions[invoiceId];
        require(sub.amount == 0, "the invoice has already been payed");
        _;
    }

    modifier onlyOwner() {
        require(
            msg.sender == owner,
            "only the contract owner can call function"
        );
        _;
    }

    modifier sufficientValue(uint256 _price) {
        require(msg.value == _price, "value does not match event price");
        _;
    }

    // Constructor code is only run when the contract
    // is created
    constructor(uint256 _eventStartTime, uint256 _eventEndTime, uint256 _price)
        public
    {
        owner = msg.sender;
        ppvEventStart = _eventStartTime;
        ppvEventEnd = _eventEndTime;
        ppvEventPrice = _price;
    }

    /**
     * @dev Changes the price of the pay per view event.
     * @param _price the new price of the pay per view event.
     */
    function changePrice(uint256 _price) public onlyOwner() notEnded(ended) {
        ppvEventPrice = _price;
        emit PriceChanged(_price);
    }

    /**
     * @dev subscribe to pay per view event.
     */
    function subscribe(bytes32 invoiceId, bytes memory signature)
        public
        payable
        notEnded(ended)
        sufficientValue(ppvEventPrice)
        notPayed(invoiceId)
    {
        // we need to validate that the invoiceId, the price and the contract address
        // matches the signature provided in the "signature" param
        // that way we make sure that our application has already binded the invoiceId
        // with a user.
        bytes32 message = prefixed(
            keccak256(abi.encodePacked(ppvEventPrice, invoiceId, this))
        );
        // the message and the signature must match. Signer must be the owner of the contract
        // although we could have signer addresses (private keys) that do not own the contract
        // but are able to sign invoices. Probably that would be best.
        require(
            recoverSigner(message, signature) == owner,
            "message was not signed by authorized party"
        );
        subscriptions[invoiceId] = Subscription(msg.value, msg.sender);
        emit NewSubscription(msg.sender, msg.value);
    }

    /// unsubscribe from the PPV event ONLY if the event has not started
    function unsubscribe(bytes32 invoiceId)
        public
        notEnded(ended)
        beforeStart(ppvEventStart)
    {
        Subscription memory sub = subscriptions[invoiceId];
        require(sub.amount > 0, "no subscription available");
        require(
            sub.subscriber == msg.sender,
            "subscriber does not match sender"
        );
        // it is important to delete this value before transfering to avoid a re-entry attack.
        delete subscriptions[invoiceId];
        // do not use transfer nor send because:
        // Those contracts will break because their fallback functions used to consume less than 2300 gas
        (bool success, ) = msg.sender.call.value(sub.amount)("");
        require(success, "transfer failed");
        emit NewUnsubscription(msg.sender, sub.amount);
    }

    /// End the PPV Event
    function eventEnd(address payable dest)
        public
        payable
        onlyOwner()
        notEnded(ended)
        afterStart(ppvEventStart)
    {
        ended = true;
        // do not use transfer nor send because:
        // Those contracts will break because their fallback functions used to consume less than 2300 gas
        (bool success, ) = dest.call.value(address(this).balance)("");
        require(success, "transfer failed");
        emit PPVEventEnded();
    }

    /// signature methods.
    function splitSignature(bytes memory sig)
        internal
        pure
        returns (uint8 v, bytes32 r, bytes32 s)
    {
        require(sig.length == 65, "invalid signature length");

        assembly {
            // first 32 bytes, after the length prefix.
            r := mload(add(sig, 32))
            // second 32 bytes.
            s := mload(add(sig, 64))
            // final byte (first byte of the next 32 bytes).
            v := byte(0, mload(add(sig, 96)))
        }

        return (v, r, s);
    }

    function recoverSigner(bytes32 message, bytes memory sig)
        internal
        pure
        returns (address)
    {
        (uint8 v, bytes32 r, bytes32 s) = splitSignature(sig);

        return ecrecover(message, v, r, s);
    }

    /// builds a prefixed hash to mimic the behavior of eth_sign.
    function prefixed(bytes32 hash) internal pure returns (bytes32) {
        return
            keccak256(
                abi.encodePacked("\x19Ethereum Signed Message:\n32", hash)
            );
    }

}
