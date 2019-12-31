import { Module } from "vuex";
import { PayPerViewEvent, CryptoCurrency } from "@/models/models";
import { fetchAuthenticated, GenericAPIResponse } from "@/lib/http/api";
import { AppRootState } from "@/store";

export interface PayPerViewEventState {
  availableEvents: PayPerViewEvent[];
  subscribedEvents: PayPerViewEvent[];
  loadingEvents: Boolean;
  errorFetchingEvents: any;
  subscribingEvent: Boolean;
  errorSubscribingEvent: any;
}

export interface SubscriptionIntent {
  eventId: Number;
  currency: CryptoCurrency;
  transactionId: String;
}

export interface SubscriptionConfirmed {
  eventId: Number;
  currency: CryptoCurrency;
  subscribedOn: Date;
}

const eventsModule: Module<PayPerViewEventState, AppRootState> = {
  namespaced: true,
  state: () => ({
    availableEvents: [],
    subscribedEvents: [],
    loadingEvents: false,
    errorFetchingEvents: null,
    subscribingEvent: false,
    errorSubscribingEvent: null
  }),
  getters: {
    availableEvents: s => s.availableEvents,
    subscribedEvents: s => s.subscribedEvents,
    loadingEvents: s => s.loadingEvents
  },
  actions: {
    retrieveEvents: async ({ commit }) => {
      commit("clearErrorFetchingEvents");
      commit("setEventsLoading");
      try {
        const response = await fetchAuthenticated("events");
        const responseJSON = (await response.json()) as GenericAPIResponse<
          PayPerViewEvent[]
        >;
        const available: PayPerViewEvent[] = [];
        const subscribed: PayPerViewEvent[] = [];
        if (responseJSON.data) {
          responseJSON.data.forEach(e => {
            if (e.subscription) {
              subscribed.push(e);
            } else {
              available.push(e);
            }
          });
        }
        commit("setAvailableEvents", available);
        commit("setSubscribedEvents", subscribed);
      } catch (err) {
        commit("setErrorFetchingEvents", err);
      } finally {
        commit("setEventsLoaded");
      }
    },
    subscribe: async (
      { commit, rootState },
      subscriptionIntent: SubscriptionIntent
    ) => {
      if (rootState.session.user.id !== 0) {
        commit("clearErrorSubscribingEvents");
        commit("setSubscribingEvent");
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
          commit("setSubscribedEvent");
        }
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
    setEventsLoading: s => {
      s.loadingEvents = true;
    },
    setEventsLoaded: s => {
      s.loadingEvents = false;
    },
    setErrorFetchingEvents: (s, payload: any) => {
      s.errorFetchingEvents = payload;
    },
    clearErrorFetchingEvents: s => {
      s.errorFetchingEvents = null;
    },
    clearErrorSubscribingEvents: s => {
      s.errorSubscribingEvent = null;
    },
    setSubscribingEvent: s => {
      s.subscribingEvent = true;
    },
    confirmSubscription: (s, payload: PayPerViewEvent) => {
      const eventIndex = s.availableEvents.findIndex(e => e.id === payload.id);
      if (eventIndex !== -1) {
        const subscribedEvent = s.availableEvents.splice(eventIndex, 1);
        s.subscribedEvents.push(subscribedEvent[0]);
      } else {
        s.subscribedEvents.push(payload);
      }
    },
    setSubscribedEvent: s => {
      s.subscribingEvent = false;
    },
    setErrorSubscribingEvents: (s, payload: any) => {
      s.errorSubscribingEvent = payload;
    }
  }
};

export default eventsModule;
