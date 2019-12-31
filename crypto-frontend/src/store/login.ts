import { Module } from "vuex";
import { apiPrefix, GenericAPIResponse, setToken } from "@/lib/http/api";
import { AppRootState } from "@/store";

interface JWTToken {
  token: String;
}

export interface LoginState {
  loading: Boolean;
  error: any;
}

export interface AuthCredentials {
  username: String;
  password: String;
}

const eventsModule: Module<LoginState, AppRootState> = {
  namespaced: true,
  state: () => ({
    availableEvents: [],
    subscribedEvents: [],
    loading: true,
    error: null
  }),
  getters: {
    loading: s => s.loading
  },
  actions: {
    login: ({ commit }, credentials: AuthCredentials) => {
      commit("setLoading");
      commit("clearError");
      fetch(`${apiPrefix}/auth`, {
        method: "POST",
        body: JSON.stringify(credentials)
      })
        .then(response => {
          return response
            .json()
            .then((responseJSON: GenericAPIResponse<JWTToken>) => {
              if (responseJSON.data) {
                setToken(responseJSON.data.token);
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
