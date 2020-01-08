import { Module } from "vuex";
import { PayPerViewEvent, CryptoCurrency } from "@/models/models";
import { fetchAuthenticated, GenericAPIResponse } from "@/lib/http/api";
import { AppRootState } from "@/store";
import { PayPerViewEventProspect } from "@/models/events";

export interface PayPerViewEventState {
  availableEvents: PayPerViewEvent[];
  subscribedEvents: PayPerViewEvent[];
  event: PayPerViewEvent | null;
  loadingEvent: boolean;
  loadingEvents: boolean;
  eventSaving: boolean;
  subscribingEvent: boolean;
  errorFetchingEvents: any;
  errorSubscribingEvent: any;
  errorSavingEvent: any;
  errorLoadingEvent: any;
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

const eventsModule: Module<PayPerViewEventState, AppRootState> = {
  namespaced: true,
  state: () => ({
    availableEvents: [],
    subscribedEvents: [],
    event: null,
    loadingEvent: false,
    loadingEvents: false,
    eventSaving: false,
    errorFetchingEvents: null,
    subscribingEvent: false,
    errorSubscribingEvent: null,
    errorSavingEvent: null,
    errorLoadingEvent: null
  }),
  getters: {
    availableEvents: s => s.availableEvents,
    subscribedEvents: s => s.subscribedEvents,
    loadingEvents: s => s.loadingEvents
  },
  actions: {
    loadEvent: async ({ commit }, eventID: number) => {
      commit("setEvent", null);
      commit("setErrorLoadingEvent", null);
      commit("setErrorSavingEvent", null);
      commit("setEventLoading", true);
      try {
        const response = await fetchAuthenticated(`events/${eventID}`);
        const responseJSON = (await response.json()) as GenericAPIResponse<
          PayPerViewEvent
        >;
        commit("setEvent", responseJSON.data);
      } catch (e) {
        commit("setErrorLoadingEvent", e);
      } finally {
        commit("setEventLoading", false);
      }
    },
    retrieveEvents: async ({ commit }) => {
      commit("setAvailableEvents", []);
      commit("setSubscribedEvents", []);
      commit("setErrorFetchingEvents", null);
      commit("setEventsLoading", true);
      try {
        const response = await fetchAuthenticated("events");
        const responseJSON = (await response.json()) as GenericAPIResponse<
          PayPerViewEvent[]
        >;
        const available: PayPerViewEvent[] = [];
        const subscribed: PayPerViewEvent[] = [];
        const events = responseJSON.data || [];
        events.forEach(e => {
          if (e.subscription) {
            subscribed.push(e);
          } else {
            available.push(e);
          }
        });
        commit("setAvailableEvents", available);
        commit("setSubscribedEvents", subscribed);
      } catch (err) {
        commit("setErrorFetchingEvents", err);
      } finally {
        commit("setEventsLoading", false);
      }
    },
    createEvent: async ({ commit }, payload: PayPerViewEventProspect) => {
      commit("setEventSaving", true);
      commit("setErrorSavingEvent", null);
      try {
        const response = await fetchAuthenticated("events", {
          method: "POST",
          body: JSON.stringify(payload)
        });
        const responseJSON = (await response.json()) as GenericAPIResponse<
          PayPerViewEvent
        >;
        commit("setEvent", responseJSON.data);
      } catch (e) {
        commit("setErrorSavingEvent", e);
      } finally {
        commit("setEventSaving", false);
      }
    },
    subscribe: async (
      { commit, rootState },
      subscriptionIntent: SubscriptionIntent
    ) => {
      if (rootState.session.user.id !== 0) {
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
    }
  },
  mutations: {
    setEventSaving: (s, payload: boolean) => {
      s.eventSaving = payload;
    },
    setAvailableEvents: (s, payload: PayPerViewEvent[]) => {
      s.availableEvents = payload;
    },
    setSubscribedEvents: (s, payload: PayPerViewEvent[]) => {
      s.subscribedEvents = payload;
    },
    setEventsLoading: (s, payload: boolean) => {
      s.loadingEvents = payload;
    },
    setEvent: (s, payload: PayPerViewEvent | null) => {
      s.event = payload;
    },
    setEventLoading: (s, payload: boolean) => {
      s.loadingEvent = payload;
    },
    setErrorFetchingEvents: (s, payload: any) => {
      s.errorFetchingEvents = payload;
    },
    setErrorSubscribingEvents: (s, payload: any) => {
      s.errorSubscribingEvent = payload;
    },
    setErrorLoadingEvent: (s, payload: any) => {
      s.errorLoadingEvent = payload;
    },
    setErrorSavingEvent: (s, payload: any) => {
      s.errorSavingEvent = payload;
    },
    setSubscribingEvent: (s, payload: boolean) => {
      s.subscribingEvent = payload;
    },
    confirmSubscription: (s, payload: PayPerViewEvent) => {
      const eventIndex = s.availableEvents.findIndex(e => e.id === payload.id);
      if (eventIndex !== -1) {
        const subscribedEvent = s.availableEvents.splice(eventIndex, 1);
        s.subscribedEvents.push(subscribedEvent[0]);
      } else {
        s.subscribedEvents.push(payload);
      }
    }
  }
};

export default eventsModule;
