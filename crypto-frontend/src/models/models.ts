import { Ability } from "casl";
import { Moment } from "moment";

export const ETH = "ETH";
export const BTC = "BTC";
export const XMR = "XMR";
export const paymentStatusUnconfirmed = "unconfirmed";
export const paymentStatusWeakConfirmation = "weakConfirmation";
export const paymentStatusConfirmed = "confirmed";
export const paymentStatusWeakCancelled = "weakCancelled";
export const paymentStatusCancelled = "cancelled";

export enum CryptoCurrency {
  ETH = "ETH",
  BTC = "BTC",
  XMR = "XMR"
}

export const strToCryptoCurrency = function(str: string): CryptoCurrency {
  if (str === BTC) {
    return CryptoCurrency.BTC;
  }
  if (str === XMR) {
    return CryptoCurrency.XMR;
  }
  return CryptoCurrency.ETH;
};

export enum PaymentStatus {
  unconfirmed = "unconfirmed",
  weakConfirmation = "weakConfirmation",
  confirmed = "confirmed",
  weakCancelled = "weakCancelled",
  cancelled = "cancelled"
}

export const strToPaymentStatus = function(str: string): PaymentStatus {
  switch (str) {
    case paymentStatusWeakConfirmation:
      return PaymentStatus.weakConfirmation;
    case paymentStatusConfirmed:
      return PaymentStatus.confirmed;
    case paymentStatusWeakCancelled:
      return PaymentStatus.weakCancelled;
    case paymentStatusCancelled:
      return PaymentStatus.cancelled;
    default:
      return PaymentStatus.unconfirmed;
  }
};

export interface Payment {
  id: number;
  userId?: number | null;
  payPerViewEventId: number;
  currency: CryptoCurrency;
  currencyPaymentId?: number | null;
  amount: number | null;
  walletAddress: string | null;
  status: PaymentStatus;
  blockHash?: string | null;
  blockNumberHex?: string | null;
  txHash?: string | null;
  txNumberHex?: string | null;
  cancelledBlockHash?: string | null;
  cancelledBlockNumberHex?: string | null;
  cancelledTxHash?: string | null;
  cancelledTxNumberHex?: string | null;
  cancelledAt?: Moment | null;
  createdAt?: Moment | null;
  updatedAt?: Moment | null;
}

export interface PayPerViewEvent {
  id: number;
  name: string;
  description: string;
  eventType: string;
  start: Moment;
  end: Moment;
  priceBTC: number;
  priceXMR: number;
  priceETH: number;
  ethContractAddr: string;
  payments?: Payment[];
}

export interface Permission {
  resource: string;
  access: string;
}

export interface User {
  id: number;
  firstName: string;
  lastName: string;
  expiry: Moment;
  permissions: Permission[];
  ability: Ability;
}
