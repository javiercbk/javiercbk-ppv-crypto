import {
  createComponent,
  computed,
  onBeforeMount,
  Ref,
  watch
} from "@vue/composition-api";
import { useState, useActions } from "@u3u/vue-hooks";
import EventEditForm from "@/components/events/event-edit-form.vue";
import NotFound from "@/components/not-found.vue";
import { PayPerViewEvent } from "@/models/models";
import router from "@/router";
import { EventFormState, EVENT_PARAM_NAME } from "@/store/events";
import { Location } from "vue-router";

export default createComponent({
  components: {
    EventEditForm,
    NotFound
  },
  setup() {
    const state = {
      ...useState("events", ["event", "eventFormState"])
    };
    const actions = {
      ...useActions("events", ["clearEvent", "loadEvent", "notFound"])
    };

    onBeforeMount(() => {
      const eventId = router.currentRoute.params[EVENT_PARAM_NAME];
      if (eventId && eventId.length > 0) {
        const eventIdNum = parseInt(eventId, 10);
        if (!isNaN(eventIdNum)) {
          actions.loadEvent(eventIdNum);
        } else {
          actions.notFound(eventIdNum);
        }
      } else {
        actions.clearEvent();
      }
    });

    watch(() => {
      const eventFormState = (state.eventFormState as Ref<EventFormState>)
        .value;
      const event = (state.event as Ref<PayPerViewEvent>).value;
      if (eventFormState === EventFormState.Created && event && event.id) {
        const rawLocation: Location = {
          name: "events-edition",
          params: { eventId: event.id.toString() }
        };
        router.replace(rawLocation);
      }
    });

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
