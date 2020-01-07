import {
  createComponent,
  onBeforeMount,
  computed,
  Ref,
  ref
} from "@vue/composition-api";
import { useState } from "@u3u/vue-hooks";
import { PayPerViewEvent } from "@/models/models";
import store from "@/store";
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
        "loadingEvents",
        "availableEvents",
        "subscribedEvents"
      ])
    };

    const hasEvents = ref(
      computed(
        () => (state.availableEvents as Ref<PayPerViewEvent[]>).value.length > 0
      )
    );

    return {
      ...state,
      hasEvents
    };
  }
});
