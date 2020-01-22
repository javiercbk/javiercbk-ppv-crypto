import { Module } from "vuex";
import moment from "moment";
import {
  PayPerViewEvent,
  Payment,
  CryptoCurrency,
  strToCryptoCurrency,
  strToPaymentStatus
} from "@/models/models";
import { fetchAuthenticated, GenericAPIResponse } from "@/lib/http/api";
import { AppRootState } from "@/store";
import { PayPerViewEventProspect } from "@/models/events";
import { CryptoCurrencyValues, usdToCrypto } from "@/lib/cryptocurrency";

export const EVENT_PARAM_NAME = "eventId";

export enum EventListState {
  Loading,
  Ready,
  Error
}

export enum EventFormState {
  Loading,
  Ready,
  ErrorLoading,
  ErrorSaving,
  EstimatingPrice,
  NotFound,
  Saving,
  Saved,
  Created
}

const stateAllowingSave: EventFormState[] = [
  EventFormState.Created,
  EventFormState.ErrorSaving,
  EventFormState.Ready,
  EventFormState.Saved
];

export const isStoreEventReady = (state: EventFormState) =>
  stateAllowingSave.indexOf(state) !== -1;

export interface PayPerViewEventState {
  availableEvents: PayPerViewEvent[];
  subscribedEvents: PayPerViewEvent[];
  estimatedPrice: CryptoCurrencyValues;
  event: PayPerViewEvent | null;
  errorEvents: Response | any | null;
  errorSubscribingEvent: Response | any | null;
  errorEvent: Response | any | null;
  eventListState: EventListState;
  eventFormState: EventFormState;
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
    estimatedPrice: {
      satoshi: 0,
      piconero: 0,
      wei: 0
    },
    event: null,
    errorEvents: null,
    errorSubscribingEvent: null,
    errorEvent: null,
    eventListState: EventListState.Ready,
    eventFormState: EventFormState.Ready
  }),
  getters: {
    availableEvents: s => s.availableEvents,
    subscribedEvents: s => s.subscribedEvents,
    event: s => s.event,
    errorEvents: s => s.errorEvents,
    errorSubscribingEvent: s => s.errorSubscribingEvent,
    errorEvent: s => s.errorEvent,
    eventListState: s => s.eventListState,
    eventFormState: s => s.eventFormState
  },
  actions: {
    notFound: ({ commit }) => {
      commit("setEvent", null);
      commit("setEventFormState", EventFormState.NotFound);
    },
    estimateEventPrice: async ({ commit, state }, priceUsDollar: number) => {
      const currentState = state.eventFormState;
      if (isStoreEventReady(currentState)) {
        commit("setEventFormState", EventFormState.EstimatingPrice);
        try {
          const dollarValue = priceUsDollar;
          if (dollarValue > 0) {
            const cryptoCurrencyValues = await usdToCrypto(dollarValue);
            commit("setEstimatedPrice", cryptoCurrencyValues);
          }
        } finally {
          commit("setEventFormState", currentState);
        }
      }
    },
    loadEvent: async ({ commit }, eventId: number) => {
      commit("setEvent", null);
      commit("setErrorEvent", null);
      commit("setEventFormState", EventFormState.Loading);
      try {
        const response = await fetchAuthenticated(`events/${eventId}`);
        if (response.ok) {
          const responseJSON = (await response.json()) as GenericAPIResponse<
            PayPerViewEventDTO
          >;
          commit("setEvent", responseJSON.data);
          commit("setEventFormState", EventFormState.Ready);
        } else if (response.status === 404) {
          commit("setEventFormState", EventFormState.NotFound);
        } else {
          throw response;
        }
      } catch (e) {
        commit("setErrorEvent", e);
        commit("setEventFormState", EventFormState.ErrorLoading);
      }
    },
    retrieveEvents: async ({ commit }) => {
      commit("setAvailableEvents", []);
      commit("setSubscribedEvents", []);
      commit("setErrorEvents", null);
      commit("setEventListState", EventListState.Loading);
      try {
        const response = await fetchAuthenticated("events");
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
    },
    saveEvent: async ({ commit }, payload: PayPerViewEventProspect) => {
      commit("setEventFormState", EventFormState.Saving);
      commit("setErrorEvent", null);
      let nextSuccessState = payload.id
        ? EventFormState.Saved
        : EventFormState.Created;
      try {
        let url = "events";
        let method = "POST";
        if (payload.id) {
          url = `events/${payload.id}`;
          method = "PUT";
        }
        const response = await fetchAuthenticated(url, {
          method,
          headers: {
            "Content-Type": "application/json"
          },
          body: JSON.stringify(payload)
        });
        const responseJSON = (await response.json()) as GenericAPIResponse<
          PayPerViewEvent
        >;
        commit("setEvent", responseJSON.data);
        commit("setEventFormState", nextSuccessState);
      } catch (e) {
        commit("setErrorEvent", e);
        commit("setEventFormState", EventFormState.ErrorSaving);
      }
    },
    subscribe: async (
      { commit, rootState },
      subscriptionIntent: SubscriptionIntent
    ) => {
      if (rootState.session.user) {
        commit("setErrorSubscribingEvents", null);
        commit("setSubscribingEvent", true);
        try {
          const response = await fetchAuthenticated(
            `events/${subscriptionIntent.eventId}`,
            {
              method: "POST",
              body: JSON.stringify({
                currency: subscriptionIntent.currency,
                transactionId: subscriptionIntent.transactionId
              })
            }
          );
          const responseJSON = (await response.json()) as GenericAPIResponse<
            PayPerViewEvent
          >;
          if (responseJSON.data) {
            commit("confirmSubscription", responseJSON.data);
          }
        } catch (err) {
          commit("setErrorSubscribingEvents", err);
        } finally {
          commit("setSubscribingEvent", false);
        }
      }
    },
    clearEvent: ({ commit }) => {
      commit("setEvent", null);
      commit("setEstimatedPrice", {
        satoshi: 0,
        piconero: 0,
        wei: 0
      });
      commit("setErrorEvent", null);
      commit("setEventFormState", EventFormState.Ready);
    }
  },
  mutations: {
    setAvailableEvents: (s, payload: PayPerViewEvent[]) => {
      s.availableEvents = payload;
    },
    setSubscribedEvents: (s, payload: PayPerViewEvent[]) => {
      s.subscribedEvents = payload;
    },
    setEstimatedPrice: (s, payload: CryptoCurrencyValues) => {
      s.estimatedPrice = payload;
    },
    setEvent: (s, payload: PayPerViewEvent | null) => {
      s.event = payload;
    },
    setErrorEvents: (s, payload: any) => {
      s.errorEvents = payload;
    },
    setErrorSubscribingEvent: (s, payload: any) => {
      s.errorSubscribingEvent = payload;
    },
    setErrorEvent: (s, payload: any) => {
      s.errorEvent = payload;
    },
    setEventListState: (s, payload: EventListState) => {
      s.eventListState = payload;
    },
    setEventFormState: (s, payload: EventFormState) => {
      s.eventFormState = payload;
    }
  }
};

export default eventsModule;
