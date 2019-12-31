import { createComponent } from "@vue/composition-api";
import { CryptoCurrency, PayPerViewEvent } from "@/models/models";

export default createComponent({
  setup(props: PayPerViewEvent) {
    return {
      ...props,
      cryptoCurrency: CryptoCurrency
    };
  }
});
