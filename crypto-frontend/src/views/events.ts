import { createComponent, onBeforeMount, computed } from "@vue/composition-api";
import { useState } from "@u3u/vue-hooks";
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

    return {
      ...state
    };
  }
});
