import { createComponent, ref, Ref, computed } from "@vue/composition-api";
const { required, minValue, numeric } = require("@vuelidate/validators");
const useVuelidate = require("@vuelidate/core").default;
import { useState } from "@u3u/vue-hooks";
import { datetimeFormatter, datetimeParser } from "@/lib/date/date";
import store from "@/store";
import router from "@/router";
import { ethAddress } from "@/lib/validators";
import { PayPerViewEvent } from "@/models/models";
import { PayPerViewEventProspect } from "@/models/events";
import { EventFormState, EVENT_PARAM_NAME } from "@/store/events";

const minUnity = minValue(1);
const minStart = minValue("start");

interface EventCreateFormProps {
  eventId?: number;
}

const stateAllowingSave: EventFormState[] = [
  EventFormState.Created,
  EventFormState.ErrorSaving,
  EventFormState.Ready,
  EventFormState.Saved
];

const canSave = (state: EventFormState) =>
  stateAllowingSave.indexOf(state) !== -1;

export default createComponent({
  props: {
    eventId: Number
  },
  setup(props: EventCreateFormProps) {
    const state = {
      ...useState("events", ["event", "eventError", "eventFormState"])
    };

    const event = (state.event as Ref<PayPerViewEvent | null>).value;

    let initName = "";
    let initDescription = "";
    let initEventType = "";
    let initStart: Date | null = null;
    let initEnd: Date | null = null;
    let initPriceBTC = 0;
    let initPriceXMR = 0;
    let initPriceETH = 0;
    let initEthContractAddr = "";

    if (event) {
      initName = event.name;
      initDescription = event.description;
      initEventType = event.eventType;
      initStart = event.start;
      initEnd = event.end;
      initPriceBTC = event.priceBTC;
      initPriceXMR = event.priceXMR;
      initPriceETH = event.priceETH;
      initEthContractAddr = event.ethContractAddr;
    }

    const usDollars = ref(0.0);

    const name = ref(initName);
    const description = ref(initDescription);
    const eventType = ref(initEventType);
    const start = ref<Date>(initStart);
    const end = ref<Date>(initEnd);
    const priceBTC = ref(initPriceBTC);
    const priceXMR = ref(initPriceXMR);
    const priceETH = ref(initPriceETH);
    const ethContractAddr = ref(initEthContractAddr);

    const $v = useVuelidate(
      {
        name: { required, $autoDirty: true },
        description: { required, $autoDirty: true },
        eventType: { required, $autoDirty: true },
        start: { required, $autoDirty: true },
        end: { required, minStart, $autoDirty: true },
        priceBTC: { required, numeric, minUnity, $autoDirty: true },
        priceXMR: { required, numeric, minUnity, $autoDirty: true },
        priceETH: { required, numeric, minUnity, $autoDirty: true }
      },
      {
        name,
        description,
        eventType,
        start,
        end,
        priceBTC,
        priceXMR,
        priceETH
      }
    );

    const computedProps = {
      loadingEvent: computed(() => {
        const eventFormState = (state.eventFormState as Ref<EventFormState>)
          .value;
        return (
          eventFormState === EventFormState.Loading || !canSave(eventFormState)
        );
      }),
      saveButtonDisabled: computed(() => {
        const eventFormState = (state.eventFormState as Ref<EventFormState>)
          .value;
        return !canSave(eventFormState);
      }),
      isEdition: computed(() => {
        const eventId = router.currentRoute.params[EVENT_PARAM_NAME];
        const eventFormState = (state.eventFormState as Ref<EventFormState>)
          .value;
        return (
          eventFormState !== EventFormState.NotFound &&
          eventId &&
          eventId.length
        );
      })
    };

    const estimateCryptoValues = async function(e: Event) {
      e.preventDefault();
    };

    const createEvent = async function(e: Event) {
      e.preventDefault();
      const eventFormState = (state.eventFormState as Ref<EventFormState>)
        .value;
      if (!$v.$invalid && canSave(eventFormState)) {
        const payPerViewProspect: PayPerViewEventProspect = {
          name: name.value,
          description: description.value,
          eventType: eventType.value,
          start: start.value as Date,
          end: end.value as Date,
          priceBTC: priceBTC.value,
          priceXMR: priceXMR.value,
          priceETH: priceETH.value
        };
        if (props.eventId) {
          payPerViewProspect.id = props.eventId;
        }
        await store.dispatch("events/saveEvent", payPerViewProspect);
      }
    };

    return {
      ...state,
      ...computedProps,
      name,
      description,
      eventType,
      start,
      end,
      priceBTC,
      priceXMR,
      priceETH,
      usDollars,
      ethContractAddr,
      datetimeFormatter,
      datetimeParser,
      $v,
      createEvent,
      estimateCryptoValues
    };
  }
});
