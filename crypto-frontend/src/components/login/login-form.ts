import { createComponent, ref } from "@vue/composition-api";
import { useState } from "@u3u/vue-hooks";
const { email, required } = require("@vuelidate/validators");
const useVuelidate = require("@vuelidate/core").default;
import store from "@/store";
import router from "@/router";
import { AuthCredentials } from "@/store/login";

export default createComponent({
  setup() {
    const storeState = {
      ...useState("login", ["loading", "error"])
    };

    const email = ref("");
    const password = ref("");

    const $v = useVuelidate(
      {
        email: { required, email, $autoDirty: true },
        password: { required, $autoDirty: true }
      },
      { email, password }
    );

    const login = function(e: Event) {
      e.preventDefault();
      if (!$v.$invalid) {
        const credentials: AuthCredentials = {
          email: email.value,
          password: password.value,
          onLoginSuccess: () => {
            let nextRoute = store.getters["session/savedRoute"];
            if (nextRoute) {
              nextRoute = { name: "events" };
            }
            router.replace(nextRoute);
            store.dispatch("session/clearSavedRoute");
          }
        };
        store.dispatch("login/login", credentials);
      }
    };

    return {
      ...storeState,
      email,
      password,
      login,
      $v
    };
  }
});
