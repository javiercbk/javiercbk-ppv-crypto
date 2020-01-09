import {
  createComponent,
  onBeforeMount,
  computed,
  Ref
} from "@vue/composition-api";
import { useState } from "@u3u/vue-hooks";
import { PayPerViewEvent, User } from "@/models/models";
import store from "@/store";
import { EventListState } from "@/store/events";
import EventCard from "@/components/events/event-card.vue";

export default createComponent({
  components: {
    EventCard
  },
  setup() {
    onBeforeMount(() => {
      store.dispatch("events/retrieveEvents");
    });

    const state = {
      ...useState("session", ["user"]),
      ...useState("events", [
        "availableEvents",
        "subscribedEvents",
        "errorEvents",
        "eventListState"
      ])
    };

    const computedProps = {
      hasEvents: computed(
        () =>
          (state.availableEvents as Ref<PayPerViewEvent[]>).value.length > 0 ||
          (state.subscribedEvents as Ref<PayPerViewEvent[]>).value.length > 0
      ),
      hasAvailableEvents: computed(
        () => (state.availableEvents as Ref<PayPerViewEvent[]>).value.length > 0
      ),
      hasSubscribedEventsEvents: computed(
        () =>
          (state.subscribedEvents as Ref<PayPerViewEvent[]>).value.length > 0
      ),
      isLoading: computed(
        () =>
          (state.eventListState as Ref<EventListState>).value ===
          EventListState.Loading
      ),
      hasError: computed(
        () =>
          (state.eventListState as Ref<EventListState>).value ===
          EventListState.Error
      ),
      userCanCreate: computed(() => {
        const user = (state.user as Ref<User | null>).value;
        if (user && user.ability) {
          return user.ability.can("write", "event");
        }
        return false;
      })
    };

    return {
      ...state,
      ...computedProps
    };
  }
});
