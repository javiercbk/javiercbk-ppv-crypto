import { Module } from "vuex";
import moment from "moment";
import { apiPrefix, GenericAPIResponse, setToken } from "@/lib/http/api";
import { AppRootState } from "@/store";
import { User } from "@/models/models";
import { defineAbilitiesFor } from "@/lib/user/abilities";

interface JWTTokenResponse {
  token: string;
  user: User;
}

export enum LoginFormState {
  Ready,
  Authenticating,
  Error,
  Authenticated
}

export interface LoginState {
  error: Response | any | null;
  loginFormState: LoginFormState;
}

export interface AuthCredentials {
  email: string;
  password: string;
}

const eventsModule: Module<LoginState, AppRootState> = {
  namespaced: true,
  state: () => ({
    error: null,
    loginFormState: LoginFormState.Ready
  }),
  getters: {
    error: s => s.error,
    loginFormState: s => s.loginFormState
  },
  actions: {
    login: async ({ commit, dispatch }, credentials: AuthCredentials) => {
      commit("setLoginFormState", LoginFormState.Authenticating);
      try {
        const response = await fetch(`${apiPrefix}/auth`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json"
          },
          body: JSON.stringify(credentials)
        });
        if (response.ok) {
          const responseJSON: GenericAPIResponse<JWTTokenResponse> = await response.json();
          if (responseJSON.data) {
            responseJSON.data.user.expiry = moment.utc(
              responseJSON.data.user.expiry
            );
            setToken(responseJSON.data.token);
            const userInSession = responseJSON.data.user;
            userInSession.ability = defineAbilitiesFor(userInSession);
            dispatch("session/setUser", userInSession, { root: true });
            commit("setLoginFormState", LoginFormState.Authenticated);
          }
        } else {
          commit("setError", response);
          commit("setLoginFormState", LoginFormState.Error);
        }
      } catch (e) {
        commit("setError", e);
        commit("setLoginFormState", LoginFormState.Error);
      }
    }
  },
  mutations: {
    setError: (s, payload: any) => {
      s.error = payload;
    },
    setLoginFormState: (s, payload: LoginFormState) => {
      s.loginFormState = payload;
    }
  }
};

export default eventsModule;
