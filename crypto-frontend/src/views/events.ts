import {
  createComponent,
  onBeforeMount,
  computed,
  Ref,
  reactive
} from "@vue/composition-api";
import { useState } from "@u3u/vue-hooks";
import { PayPerViewEvent } from "@/models/models";
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
      ...useState("events", [
        "availableEvents",
        "subscribedEvents",
        "errorEvents",
        "eventListState"
      ])
    };

    const computedProps = reactive({
      hasEvents: computed(
        () =>
          (state.availableEvents as Ref<PayPerViewEvent[]>).value.length > 0 ||
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
      )
    });

    return {
      ...state,
      ...computedProps
    };
  }
});
