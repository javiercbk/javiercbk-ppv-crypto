import { Route } from "vue-router";
import { Module } from "vuex";
import { Ability } from "casl";
import { User } from "@/models/models";
import { apiPrefix, GenericAPIResponse } from "@/lib/http/api";
import { AppRootState } from "@/store";
import defineAbilitiesFor from "@/lib/user/abilities";

export interface SessionState {
  user: User;
  loading: Boolean;
  error: any;
  userRequested: Boolean;
  savedRoute: Route | null;
}

const sessionModule: Module<SessionState, AppRootState> = {
  namespaced: true,
  state: () => ({
    user: {
      id: 0,
      username: "",
      firstName: "",
      lastName: "",
      permissions: [],
      ability: new Ability([])
    },
    loading: false,
    error: null,
    userRequested: false,
    savedRoute: null
  }),
  getters: {
    user: s => s.user,
    loading: s => s.loading,
    error: s => s.error,
    savedRoute: s => s.savedRoute,
    userRequested: s => s.userRequested
  },
  actions: {
    retrieveUser: ({ commit }) => {
      commit("setLoading", true);
      commit("setError", null);
      fetch(`${apiPrefix}/users/current`)
        .then(response => {
          return response
            .json()
            .then((responseJSON: GenericAPIResponse<User>) => {
              if (responseJSON.data) {
                const userInSession = responseJSON.data;
                userInSession.ability = defineAbilitiesFor(userInSession);
                commit("setUser", userInSession);
              }
            });
        })
        .catch(err => {
          commit("setError", err);
        })
        .finally(() => {
          commit("setUserRequested", true);
          commit("setLoading", false);
        });
    },
    setUser: ({ commit }, payload: User) => {
      commit("setUser", payload);
    },
    clearSavedRoute: ({ commit }) => {
      commit("setSavedRoute", null);
    },
    saveRoute: ({ commit }, payload: Route) => {
      commit("setSavedRoute", payload);
    }
  },
  mutations: {
    setUser: (s, payload: User) => {
      s.user = payload;
    },
    setUserRequested: (s, payload: Boolean) => {
      s.userRequested = payload;
    },
    setSavedRoute: (s, payload: Route | null) => {
      s.savedRoute = payload;
    },
    setLoading: (s, payload: Boolean) => {
      s.loading = payload;
    },
    setError: (s, payload: any) => {
      s.error = payload;
    }
  }
};

export default sessionModule;
