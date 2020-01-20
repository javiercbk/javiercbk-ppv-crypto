import {
  createComponent,
  computed,
  onBeforeMount,
  Ref
} from "@vue/composition-api";
import { useState } from "@u3u/vue-hooks";
import EventEditForm from "@/components/events/event-edit-form.vue";
import NotFound from "@/components/not-found.vue";
import store from "@/store";
import router from "@/router";
import { EventFormState, EVENT_PARAM_NAME } from "@/store/events";

export default createComponent({
  components: {
    EventEditForm,
    NotFound
  },
  setup() {
    onBeforeMount(() => {
      const eventId = router.currentRoute.params[EVENT_PARAM_NAME];
      if (eventId && eventId.length > 0) {
        const eventIdNum = parseInt(eventId, 10);
        if (!isNaN(eventIdNum)) {
          store.dispatch("events/loadEvent", eventIdNum);
        } else {
          store.dispatch("events/notFound", eventIdNum);
        }
      } else {
        store.dispatch("events/clearEvent");
      }
    });

    const state = {
      ...useState("events", ["eventFormState"])
    };

    const computedProps = {
      eventNotFound: computed(() => {
        const eventFormState = (state.eventFormState as Ref<EventFormState>)
          .value;
        return eventFormState === EventFormState.NotFound;
      })
    };

    return {
      ...computedProps
    };
  }
});
