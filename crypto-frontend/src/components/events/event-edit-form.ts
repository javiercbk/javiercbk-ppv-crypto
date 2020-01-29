import {
  createComponent,
  ref,
  Ref,
  computed,
  watch
} from "@vue/composition-api";
const { required, minValue, numeric } = require("@vuelidate/validators");
const useVuelidate = require("@vuelidate/core").default;
import { useState, useActions } from "@u3u/vue-hooks";
import { datetimeFormatter, datetimeParser } from "@/lib/date/date";
import { isStoreEventReady } from "@/store/event";
import router from "@/router";
import { PayPerViewEvent } from "@/models/models";
import { PayPerViewEventProspect } from "@/models/events";
import { EventFormState, EVENT_PARAM_NAME } from "@/store/event";

const minUnity = minValue(1);

interface EventCreateFormProps {
  eventId?: number;
}

export default createComponent({
  props: {
    eventId: Number
  },
  setup(props: EventCreateFormProps) {
    const state = {
      ...useState("event", [
        "event",
        "eventError",
        "eventFormState",
        "estimatedPrice"
      ])
    };

    const actions = {
      ...useActions("event", ["saveEvent", "estimateEventPrice"])
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
      initStart = event.start.toDate();
      initEnd = event.end.toDate();
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

    const largerThanStart = minValue(start.value);

    const $v = useVuelidate(
      {
        name: { required, $autoDirty: true },
        description: { required, $autoDirty: true },
        eventType: { required, $autoDirty: true },
        start: { required, $autoDirty: true },
        end: { required, largerThanStart, $autoDirty: true },
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
      isReady: computed(() => {
        const eventFormState = (state.eventFormState as Ref<EventFormState>)
          .value;
        return isStoreEventReady(eventFormState);
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

    watch(() => {
      const estimatedPrice = state.estimatedPrice.value;
      if (estimatedPrice.satoshi) {
        priceBTC.value = estimatedPrice.satoshi;
      }
      if (estimatedPrice.piconero) {
        priceXMR.value = estimatedPrice.piconero;
      }
      if (estimatedPrice.wei) {
        priceETH.value = estimatedPrice.wei;
      }
    });

    const estimateCryptoValues = async function(e: Event) {
      e.preventDefault();
      if (computedProps.isReady.value) {
        actions.estimateEventPrice(usDollars.value);
      }
    };

    const createEvent = async function(e: Event) {
      e.preventDefault();
      if (!$v.$invalid && computedProps.isReady.value) {
        const event = (state.event as Ref<PayPerViewEvent | null>).value;
        const eventId = event ? event.id : undefined;
        const payPerViewProspect: PayPerViewEventProspect = {
          id: eventId,
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
        await actions.saveEvent(payPerViewProspect);
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
