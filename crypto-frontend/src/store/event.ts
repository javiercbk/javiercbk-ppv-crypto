import { Module } from "vuex";
import moment from "moment";
import {
  PayPerViewEvent,
  Payment,
  strToCryptoCurrency,
  strToPaymentStatus
} from "@/models/models";
import {
  PaymentDTO,
  PayPerViewEventDTO,
  SubscriptionIntent
} from "@/models/events";
import { fetchAuthenticated, GenericAPIResponse } from "@/lib/http/api";
import { AppRootState } from "@/store";
import { PayPerViewEventProspect } from "@/models/events";
import { CryptoCurrencyValues, usdToCrypto } from "@/lib/cryptocurrency";

export const EVENT_PARAM_NAME = "eventId";

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
  estimatedPrice: CryptoCurrencyValues;
  event: PayPerViewEvent | null;
  errorSubscribingEvent: Response | any | null;
  errorEvent: Response | any | null;
  eventFormState: EventFormState;
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
    eventFormState: EventFormState.Ready
  }),
  getters: {
    event: s => s.event,
    errorSubscribingEvent: s => s.errorSubscribingEvent,
    errorEvent: s => s.errorEvent,
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
    setEstimatedPrice: (s, payload: CryptoCurrencyValues) => {
      s.estimatedPrice = payload;
    },
    setEvent: (s, payload: PayPerViewEvent | null) => {
      s.event = payload;
    },
    setErrorSubscribingEvent: (s, payload: any) => {
      s.errorSubscribingEvent = payload;
    },
    setErrorEvent: (s, payload: any) => {
      s.errorEvent = payload;
    },
    setEventFormState: (s, payload: EventFormState) => {
      s.eventFormState = payload;
    }
  }
};

export default eventsModule;
