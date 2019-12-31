import { Module } from "vuex";
import { User } from "@/models/models";
import { apiPrefix, GenericAPIResponse } from "@/lib/http/api";
import { AppRootState } from "@/store";

export interface SessionState {
  user: User;
  loading: Boolean;
  error: any;
}

const sessionModule: Module<SessionState, AppRootState> = {
  namespaced: true,
  state: () => ({
    user: {
      id: 0,
      username: "",
      firstName: "",
      lastName: ""
    },
    loading: false,
    error: null
  }),
  getters: {
    user: s => s.user
  },
  actions: {
    retrieveUser: ({ commit }) => {
      commit("setLoading");
      commit("clearError");
      fetch(`${apiPrefix}/users/current`)
        .then(response => {
          return response
            .json()
            .then((responseJSON: GenericAPIResponse<User>) => {
              if (responseJSON.data) {
                commit("setUser", responseJSON.data);
              }
            });
        })
        .catch(err => {
          commit("setError", err);
        })
        .finally(() => {
          commit("setLoaded");
        });
    }
  },
  mutations: {
    setUser: (s, payload: User) => {
      s.user = payload;
    },
    setLoading: s => {
      s.loading = true;
    },
    setLoaded: (s, payload: User) => {
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

export default sessionModule;
