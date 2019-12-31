import Vue from "vue";
import Vuex from "vuex";

import events, { PayPerViewEventState } from "@/store/events";
import login, { LoginState } from "@/store/login";
import session, { SessionState } from "@/store/session";

Vue.use(Vuex);

export interface AppRootState {
  events: PayPerViewEventState;
  login: LoginState;
  session: SessionState;
}

export default new Vuex.Store<AppRootState>({
  modules: {
    events: events,
    login: login,
    session: session
  }
});
