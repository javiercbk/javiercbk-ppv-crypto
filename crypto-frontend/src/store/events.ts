import { Module } from "vuex";
import moment from "moment";
import {
  PayPerViewEvent,
  Payment,
  strToCryptoCurrency,
  strToPaymentStatus
} from "@/models/models";
import { PaymentDTO, PayPerViewEventDTO } from "@/models/events";
import {
  fetchOptinallyAuthenticated,
  GenericAPIResponse
} from "@/lib/http/api";
import { AppRootState } from "@/store";

export enum EventListState {
  Loading,
  Ready,
  Error
}

export interface PayPerViewEventState {
  availableEvents: PayPerViewEvent[];
  subscribedEvents: PayPerViewEvent[];
  errorEvents: Response | any | null;
  eventListState: EventListState;
}

const toPayPerViewEvent = function(event: PayPerViewEventDTO): PayPerViewEvent {
  return {
    id: event.id, //number;
    name: event.name, //string;
    description: event.description, //string;
    eventType: event.eventType, //string;
    start: moment(event.start, moment.ISO_8601, true), //string;
    end: moment(event.end, moment.ISO_8601, true), //string;
    priceBTC: event.priceBTC, //number;
    priceXMR: event.priceXMR, //number;
    priceETH: event.priceETH, //number;
    ethContractAddr: event.ethContractAddr, //string;
    payments:
      event.payments && event.payments.length > 0
        ? event.payments.map(toPayment)
        : [] // PaymentDTO[] | null;
  };
};

const toPayment = function(p: PaymentDTO): Payment {
  return {
    id: p.id, //number;
    userId: p.userId,
    payPerViewEventId: p.payPerViewEventId,
    currency: strToCryptoCurrency(p.currency),
    currencyPaymentId: p.currencyPaymentId,
    amount: p.amount,
    walletAddress: p.walletAddress,
    status: strToPaymentStatus(p.status),
    blockHash: p.blockHash,
    blockNumberHex: p.blockNumberHex,
    txHash: p.txHash,
    txNumberHex: p.txNumberHex,
    cancelledBlockHash: p.cancelledBlockHash,
    cancelledBlockNumberHex: p.cancelledBlockNumberHex,
    cancelledTxHash: p.cancelledTxHash,
    cancelledTxNumberHex: p.cancelledTxNumberHex,
    cancelledAt: p.cancelledAt
      ? moment(p.cancelledAt, moment.ISO_8601, true)
      : null,
    createdAt: p.createdAt ? moment(p.createdAt, moment.ISO_8601, true) : null,
    updatedAt: p.updatedAt ? moment(p.updatedAt, moment.ISO_8601, true) : null
  };
};

const eventsModule: Module<PayPerViewEventState, AppRootState> = {
  namespaced: true,
  state: () => ({
    availableEvents: [],
    subscribedEvents: [],
    errorEvents: null,
    errorSubscribingEvent: null,
    eventListState: EventListState.Ready
  }),
  getters: {
    availableEvents: s => s.availableEvents,
    subscribedEvents: s => s.subscribedEvents,
    errorEvents: s => s.errorEvents,
    eventListState: s => s.eventListState
  },
  actions: {
    retrieveEvents: async ({ commit }) => {
      commit("setAvailableEvents", []);
      commit("setSubscribedEvents", []);
      commit("setErrorEvents", null);
      commit("setEventListState", EventListState.Loading);
      try {
        const response = await fetchOptinallyAuthenticated("events");
        if (response.ok) {
          const responseJSON = (await response.json()) as GenericAPIResponse<
            PayPerViewEventDTO[]
          >;
          const available: PayPerViewEvent[] = [];
          const subscribed: PayPerViewEvent[] = [];
          const events = responseJSON.data || [];
          events.forEach(e => {
            const event = toPayPerViewEvent(e);
            if (e.payments && e.payments.length > 0) {
              subscribed.push(event);
            } else {
              available.push(event);
            }
          });
          commit("setAvailableEvents", available);
          commit("setSubscribedEvents", subscribed);
          commit("setEventListState", EventListState.Ready);
        } else {
          throw response;
        }
      } catch (err) {
        commit("setErrorEvents", err);
        commit("setEventListState", EventListState.Error);
      }
    }
  },
  mutations: {
    setAvailableEvents: (s, payload: PayPerViewEvent[]) => {
      s.availableEvents = payload;
    },
    setSubscribedEvents: (s, payload: PayPerViewEvent[]) => {
      s.subscribedEvents = payload;
    },
    setErrorEvents: (s, payload: any) => {
      s.errorEvents = payload;
    },
    setEventListState: (s, payload: EventListState) => {
      s.eventListState = payload;
    }
  }
};

export default eventsModule;
