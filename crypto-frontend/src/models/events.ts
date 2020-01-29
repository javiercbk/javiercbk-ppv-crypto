import { CryptoCurrency } from "@/models/models";

export interface PayPerViewEventProspect {
  id?: number;
  name: string;
  description: string;
  eventType: string;
  start: Date;
  end: Date;
  priceBTC: number;
  priceXMR: number;
  priceETH: number;
}

export interface SubscriptionIntent {
  eventId: number;
  currency: CryptoCurrency;
  transactionId: string;
}

export interface SubscriptionConfirmed {
  eventId: number;
  currency: CryptoCurrency;
  subscribedOn: Date;
}

export interface PaymentDTO {
  id: number;
  userId?: number | null;
  payPerViewEventId: number;
  currency: string;
  currencyPaymentId?: number | null;
  amount: number | null;
  walletAddress: string | null;
  status: string;
  blockHash?: string | null;
  blockNumberHex?: string | null;
  txHash?: string | null;
  txNumberHex?: string | null;
  cancelledBlockHash?: string | null;
  cancelledBlockNumberHex?: string | null;
  cancelledTxHash?: string | null;
  cancelledTxNumberHex?: string | null;
  cancelledAt?: string | null;
  createdAt?: string | null;
  updatedAt?: string | null;
}

export interface PayPerViewEventDTO {
  id: number;
  name: string;
  description: string;
  eventType: string;
  start: string;
  end: string;
  priceBTC: number;
  priceXMR: number;
  priceETH: number;
  ethContractAddr: string;
  payments?: PaymentDTO[] | null;
}
