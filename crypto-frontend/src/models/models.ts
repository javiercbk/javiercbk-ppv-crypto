import { Ability } from "casl";

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
  id: number;
  currency: CryptoCurrency;
  currencyPaymentId?: string;
  blockHash?: string;
  blockTimestamp?: Date;
  walletAddress?: string;
  amount: number;
  status: PaymentStatus;
}

export interface PayPerViewEvent {
  id: number;
  name: string;
  description: string;
  eventType: string;
  start: Date;
  end: Date;
  priceETH: number;
  priceBTC: number;
  priceXMR: number;
  subscription?: Subscription;
}

export interface Permission {
  resource: string;
  access: string;
}

export interface User {
  id: number;
  firstName: string;
  lastName: string;
  permissions: Permission[];
  ability: Ability;
}
