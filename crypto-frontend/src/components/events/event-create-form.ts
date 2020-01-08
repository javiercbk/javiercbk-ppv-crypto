import { createComponent, onBeforeMount, ref } from "@vue/composition-api";
const { email, required, minValue, numeric } = require("@vuelidate/validators");
const useVuelidate = require("@vuelidate/core").default;
import { useState } from "@u3u/vue-hooks";
import store from "@/store";
import router from "@/router";
import { ethAddress } from "@/lib/validators";
import { PayPerViewEventProspect } from "@/models/events";

const minUnity = minValue(1);
const minStart = minValue("start");

interface EventCreateFormProps {
  eventId?: number;
}

export default createComponent({
  setup(props: EventCreateFormProps) {
    onBeforeMount(() => {
      if (props.eventId && props.eventId > 0) {
        store.dispatch("events/loadEvent");
      }
    });

    const state = {
      ...useState("events", [
        "loadingEvent",
        "eventSaving",
        "errorSavingEvent",
        "errorLoadingEvent"
      ])
    };
    const name = ref("");
    const description = ref("");
    const eventType = ref("");
    const start = ref<Date>(null);
    const end = ref<Date>(null);
    const priceBTC = ref(0);
    const priceXMR = ref(0);
    const ethContractAddr = ref("");

    const $v = useVuelidate(
      {
        name: { required, $autoDirty: true },
        description: { required, $autoDirty: true },
        eventType: { required, $autoDirty: true },
        start: { required, $autoDirty: true },
        end: { required, minStart, $autoDirty: true },
        priceBTC: { required, numeric, minUnity, $autoDirty: true },
        priceXMR: { required, numeric, minUnity, $autoDirty: true },
        ethContractAddr: { required, ethAddress, $autoDirty: true }
      },
      {
        name,
        description,
        eventType,
        start,
        end,
        priceBTC,
        priceXMR,
        ethContractAddr
      }
    );

    const createEvent = async function(e: Event) {
      e.preventDefault();
      if (!$v.$invalid) {
        const payPerViewProspect: PayPerViewEventProspect = {
          name: name.value,
          description: description.value,
          eventType: eventType.value,
          start: start.value as Date,
          end: end.value as Date,
          priceBTC: priceBTC.value,
          priceXMR: priceXMR.value,
          ethContractAddr: ethContractAddr.value
        };
        if (props.eventId) {
          payPerViewProspect.id = props.eventId;
        }
        await store.dispatch("events/saveEvent", payPerViewProspect);
      }
    };

    return {
      ...state,
      name,
      description,
      eventType,
      start,
      end,
      priceBTC,
      priceXMR,
      ethContractAddr,
      $v,
      createEvent
    };
  }
});
