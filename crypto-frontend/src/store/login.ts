import { Module } from "vuex";
import { apiPrefix, GenericAPIResponse, setToken } from "@/lib/http/api";
import { AppRootState } from "@/store";
import { User } from "@/models/models";
import defineAbilitiesFor from "@/lib/user/abilities";

interface JWTTokenResponse {
  token: string;
  user: User;
}

export interface LoginState {
  loading: Boolean;
  error: any;
}

export interface AuthCredentials {
  email: string;
  password: string;
  onLoginSuccess?: () => void;
}

const eventsModule: Module<LoginState, AppRootState> = {
  namespaced: true,
  state: () => ({
    loading: true,
    error: null
  }),
  getters: {
    loading: s => s.loading,
    error: s => s.error
  },
  actions: {
    login: ({ commit, dispatch }, credentials: AuthCredentials) => {
      commit("setLoading");
      commit("clearError");
      fetch(`${apiPrefix}/auth`, {
        method: "POST",
        body: JSON.stringify(credentials)
      })
        .then(response => {
          return response
            .json()
            .then((responseJSON: GenericAPIResponse<JWTTokenResponse>) => {
              if (responseJSON.data) {
                setToken(responseJSON.data.token);
                const userInSession = responseJSON.data.user;
                userInSession.ability = defineAbilitiesFor(userInSession);
                dispatch("session/setUser", userInSession, { root: true });
                return true;
              }
              return false;
            })
            .then((success: Boolean) => {
              if (success && credentials.onLoginSuccess) {
                credentials.onLoginSuccess();
              }
            });
        })
        .catch(err => {
          commit("setError", err);
        })
        .finally(() => {
          commit("setLoaded");
        });
      commit("setAvailableEvents");
    }
  },
  mutations: {
    setLoading: s => {
      s.loading = true;
    },
    setLoaded: s => {
      s.loading = false;
    },
    setError: (s, payload: any) => {
      s.error = payload;
    },
    clearError: s => {
      s.error = null;
    }
  }
};

export default eventsModule;
