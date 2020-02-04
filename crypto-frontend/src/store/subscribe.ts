import { Module } from "vuex";
import { fetchAuthenticated, GenericAPIResponse } from "@/lib/http/api";
import { AppRootState } from "@/store";
import { CryptoCurrency } from "@/models/models";
import { subscribeToPPV } from "@/lib/cryptocurrency/eth/connect";

export interface SubscribeAction {
  eventId: number;
  cryptoCurrency: CryptoCurrency;
  address?: string;
}

export enum SubscriptionState {
  Subscribing,
  Error,
  Cancelled,
  AwaitingConfirmation,
  Subscribed
}

export interface SubscriptionIntent {
  eventId: number;
  currency: CryptoCurrency;
  state: SubscriptionState;
  address?: string;
  error?: any;
}

interface SubscriptionInvoiceDTO {
  eventId: number;
  currency: CryptoCurrency;
  address: string;
  expires: string;
}

interface SubscriptionError {
  eventId: number;
  error: any;
}

interface SubscribeState {
  subscriptions: SubscriptionIntent[];
}

interface SubscriptionIntentResult {
  eventId: number;
  transactionId: string;
  ok: boolean;
  cancelled: boolean;
}

const findSubscriptionById = (eventId: number) => (sub: SubscriptionIntent) =>
  sub.eventId === eventId;

const _subscribeWithETH = async function(
  intent: SubscriptionIntent
): Promise<SubscriptionIntentResult> {
  if (intent.address) {
    await subscribeToPPV(intent.address);
    return {
      eventId: intent.eventId,
      transactionId: "dummy-transaction"
    };
  }
};

const eventsModule: Module<SubscribeState, AppRootState> = {
  namespaced: true,
  state: () => ({
    subscriptions: []
  }),
  getters: {
    subscriptions: s => s.subscriptions
  },
  actions: {
    subscribe: async (
      { commit, rootState },
      subscriptionIntent: SubscriptionIntent
    ) => {
      if (rootState.session.user) {
        commit("clearOrAddSubscription", subscriptionIntent);
        let subscriptionResult: SubscriptionIntentResult | undefined;
        try {
          switch (subscriptionIntent.currency) {
            case CryptoCurrency.ETH:
              subscriptionResult = await _subscribeWithETH(subscriptionIntent);
              break;
          }
        } catch (err) {
          commit("setErrorSubscribingEvents", {
            eventId: subscriptionIntent.eventId,
            error: err
          });
        }
        if (subscriptionResult) {
          if (subscriptionResult.ok) {
            commit("subscriptionIntentCompleted", subscriptionResult);
          } else if (subscriptionResult.cancelled) {
            commit("subscriptionIntentCancelled", subscriptionIntent);
          }
        }
      }
    }
  },
  mutations: {
    setErrorSubscribingEvents: (s, payload: SubscriptionError) => {
      const idx = s.subscriptions.findIndex(
        findSubscriptionById(payload.eventId)
      );
      if (idx >= 0) {
        s.subscriptions[idx].error = payload.error;
      }
    },
    subscriptionIntentCompleted: (s, payload: SubscriptionIntentResult) => {
      const idx = s.subscriptions.findIndex(
        findSubscriptionById(payload.eventId)
      );
      if (idx >= 0) {
        const sub = s.subscriptions[idx];
        sub.state = SubscriptionState.Subscribed;
        delete sub.error;
      }
    },
    subscriptionIntentCancelled: (s, payload: SubscriptionIntent) => {
      const idx = s.subscriptions.findIndex(
        findSubscriptionById(payload.eventId)
      );
      if (idx >= 0) {
        const sub = s.subscriptions[idx];
        sub.state = SubscriptionState.Cancelled;
        delete sub.error;
      }
    },
    clearOrAddSubscription: (s, payload: SubscriptionIntent) => {
      const idx = s.subscriptions.findIndex(
        findSubscriptionById(payload.eventId)
      );
      if (idx >= 0) {
        const sub = s.subscriptions[idx];
        sub.state = SubscriptionState.Subscribing;
        delete sub.error;
      } else {
        s.subscriptions.push({
          eventId: payload.eventId,
          currency: payload.currency,
          state: SubscriptionState.Subscribing
        });
      }
    }
  }
};

export default eventsModule;
