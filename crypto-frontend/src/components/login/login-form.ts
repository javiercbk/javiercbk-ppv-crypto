import {
  createComponent,
  computed,
  ref,
  Ref,
  watch
} from "@vue/composition-api";
import { Route } from "vue-router";
import { useState, useActions } from "@u3u/vue-hooks";
const { email, required } = require("@vuelidate/validators");
const useVuelidate = require("@vuelidate/core").default;
import router from "@/router";
import { AuthCredentials, LoginFormState } from "@/store/login";

const stateAllowingSubmit: LoginFormState[] = [
  LoginFormState.Ready,
  LoginFormState.Error
];

const canSubmit = (state: LoginFormState) =>
  stateAllowingSubmit.indexOf(state) !== -1;

export default createComponent({
  setup() {
    const state = {
      ...useState("login", ["loginFormState", "error"]),
      ...useState("session", ["savedRoute"])
    };

    const actions = {
      ...useActions("login", ["login"]),
      ...useActions("session", ["clearSavedRoute"])
    };

    const computedProps = {
      hasError: computed(
        () =>
          (state.loginFormState as Ref<LoginFormState>).value ===
          LoginFormState.Error
      ),
      isRequesting: computed(
        () =>
          (state.loginFormState as Ref<LoginFormState>).value ===
          LoginFormState.Authenticating
      )
    };

    watch(() => {
      const loginFormState = (state.loginFormState as Ref<LoginFormState>)
        .value;
      if (loginFormState === LoginFormState.Authenticated) {
        let nextRoute: any = (state.savedRoute as Ref<Route | null>).value;
        if (!nextRoute) {
          nextRoute = { name: "events" };
        }
        router.replace(nextRoute);
        actions.clearSavedRoute();
      }
    });

    const userEmail = ref("");
    const password = ref("");

    const $v = useVuelidate(
      {
        userEmail: { required, email, $autoDirty: true },
        password: { required, $autoDirty: true }
      },
      { userEmail, password }
    );

    const login = function(e: Event) {
      e.preventDefault();
      const loginFormState = (state.loginFormState as Ref<LoginFormState>)
        .value;

      if (!$v.$invalid && canSubmit(loginFormState)) {
        const credentials: AuthCredentials = {
          email: userEmail.value,
          password: password.value
        };
        actions.login(credentials);
      }
    };

    return {
      ...state,
      ...computedProps,
      userEmail,
      password,
      login,
      $v
    };
  }
});
