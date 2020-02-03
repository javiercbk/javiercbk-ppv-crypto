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
  error?: any;
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

const _replaceSubscriptionFactory = (
  eventId: number,
  replacer: (s: SubscriptionIntent) => SubscriptionIntent
) => (sub: SubscriptionIntent): SubscriptionIntent => {
  if (sub.eventId === eventId) {
    return replacer(sub);
  } else {
    return sub;
  }
};

const _subscribeWithETH = async function(
  intent: SubscriptionIntent
): Promise<SubscriptionIntentResult> {};

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
      s.subscriptions = s.subscriptions.map(
        _replaceSubscriptionFactory(payload.eventId, sub => {
          const errorSub = Object.assign({}, sub, {
            state: SubscriptionState.Error
          });
          errorSub.error = payload.error;
          return errorSub;
        })
      );
    },
    subscriptionIntentCompleted: (s, payload: SubscriptionIntentResult) => {
      s.subscriptions = s.subscriptions.map(
        _replaceSubscriptionFactory(payload.eventId, sub => {
          const cancelledSub = Object.assign({}, sub, {
            state: SubscriptionState.Cancelled
          });
          delete cancelledSub.error;
          return cancelledSub;
        })
      );
    },
    subscriptionIntentCancelled: (s, payload: SubscriptionIntent) => {
      s.subscriptions = s.subscriptions.map(
        _replaceSubscriptionFactory(payload.eventId, sub => {
          const cancelledSub = Object.assign({}, sub, {
            state: SubscriptionState.Cancelled
          });
          delete cancelledSub.error;
          return cancelledSub;
        })
      );
    },
    clearOrAddSubscription: (s, payload: SubscriptionIntent) => {
      let found = false;
      s.subscriptions = s.subscriptions.map(
        _replaceSubscriptionFactory(payload.eventId, sub => {
          found = true;
          const clearSub = Object.assign({}, sub, {
            state: SubscriptionState.Subscribing
          });
          delete clearSub.error;
          return clearSub;
        })
      );
      if (!found) {
        s.subscriptions = [
          ...s.subscriptions,
          {
            eventId: payload.eventId,
            currency: payload.currency,
            state: SubscriptionState.Subscribing
          }
        ];
      }
    },
    setSubscriptions: (s, payload: SubscriptionIntent[]) => {
      s.subscriptions = payload;
    }
  }
};

export default eventsModule;
