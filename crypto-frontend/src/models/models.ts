export enum CryptoCurrency {
  ETH = "ETH",
  BTC = "BTC",
  XMR = "XMR"
}

export enum PaymentStatus {
  unconfirmed = "unconfirmed",
  weakConfirmation = "weakConfirmation",
  confirmed = "confirmed"
}

export interface Subscription {
  id: Number;
  currency: CryptoCurrency;
  currencyPaymentId?: String;
  blockHash?: String;
  blockTimestamp?: Date;
  walletAddress?: String;
  amount: Number;
  status: PaymentStatus;
}

export interface PayPerViewEvent {
  id: Number;
  name: String;
  description: String;
  eventType: String;
  start: Date;
  end: Date;
  priceETH: Number;
  priceBTC: Number;
  priceXMR: Number;
  subscription?: Subscription;
}

export interface User {
  id: Number;
  username: String;
  firstName: String;
  lastName: String;
}
