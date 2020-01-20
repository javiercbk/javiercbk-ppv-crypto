import { Module } from "vuex";
import { PayPerViewEvent, CryptoCurrency } from "@/models/models";
import { fetchAuthenticated, GenericAPIResponse } from "@/lib/http/api";
import { AppRootState } from "@/store";
import { PayPerViewEventProspect } from "@/models/events";

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
  NotFound,
  Saving,
  Saved,
  Created
}

export interface PayPerViewEventState {
  availableEvents: PayPerViewEvent[];
  subscribedEvents: PayPerViewEvent[];
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

const eventsModule: Module<PayPerViewEventState, AppRootState> = {
  namespaced: true,
  state: () => ({
    availableEvents: [],
    subscribedEvents: [],
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
      commit("setEventState", EventFormState.NotFound);
    },
    loadEvent: async ({ commit }, eventId: number) => {
      commit("setEvent", null);
      commit("setErrorEvent", null);
      commit("setEventState", EventFormState.Loading);
      try {
        const response = await fetchAuthenticated(`events/${eventId}`);
        if (response.ok) {
          const responseJSON = (await response.json()) as GenericAPIResponse<
            PayPerViewEvent
          >;
          commit("setEvent", responseJSON.data);
          commit("setEventState", EventFormState.Ready);
        } else if (response.status === 404) {
          commit("setEventState", EventFormState.NotFound);
        } else {
          throw response;
        }
      } catch (e) {
        commit("setErrorEvent", e);
        commit("setEventState", EventFormState.ErrorLoading);
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
      commit("setEventState", EventFormState.Saving);
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
          body: JSON.stringify(payload)
        });
        const responseJSON = (await response.json()) as GenericAPIResponse<
          PayPerViewEvent
        >;
        commit("setEvent", responseJSON.data);
        commit("setEventState", nextSuccessState);
      } catch (e) {
        commit("setErrorEvent", e);
        commit("setEventState", EventFormState.ErrorSaving);
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
      commit("setErrorLoadingEvent", null);
      commit("setErrorSavingEvent", null);
      commit("setEventLoading", false);
    }
  },
  mutations: {
    setAvailableEvents: (s, payload: PayPerViewEvent[]) => {
      s.availableEvents = payload;
    },
    setSubscribedEvents: (s, payload: PayPerViewEvent[]) => {
      s.subscribedEvents = payload;
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
