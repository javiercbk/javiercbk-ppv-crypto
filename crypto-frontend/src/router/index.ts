import Vue, { AsyncComponent } from "vue";
import VueRouter, { RouteConfig, Route, RawLocation } from "vue-router";
import store from "@/store";
import { User } from "@/models/models";

Vue.use(VueRouter);

const EventsComponent: AsyncComponent = () =>
  import(/* webpackChunkName: "events" */ "../views/events.vue");

const LoginComponent: AsyncComponent = () =>
  import(/* webpackChunkName: "login" */ "../views/login.vue");

export const EVENTS_ROUTE_NAME = "events";

const routes: RouteConfig[] = [
  {
    path: "/",
    name: EVENTS_ROUTE_NAME,
    component: EventsComponent,
    meta: {
      public: true
    }
  },
  {
    path: "/events/create",
    name: "events-create",
    component: EventsComponent,
    meta: {
      public: false,
      resource: "Event"
    }
  },
  {
    path: "/login",
    name: "login",
    component: LoginComponent,
    meta: {
      public: true
    }
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

const routerSecurityCheck = (
  user: User,
  to: Route,
  from: Route,
  isFirstTime: Boolean,
  next: (to?: RawLocation | false | ((vm: Vue) => any) | void) => void
) => {
  let nextResolve = from as any;
  if (user) {
    if (to.meta && to.meta.public && to.name !== EVENTS_ROUTE_NAME) {
      nextResolve = {
        name: EVENTS_ROUTE_NAME
      };
    } else if (to.meta.role) {
      if (!user.ability.can("read", to.meta.resource)) {
        // user has not enough privileged to see the view
        // TODO define what to do here
        nextResolve = {
          name: EVENTS_ROUTE_NAME
        };
      }
    }
  } else if (to.meta && !to.meta.public) {
    if (isFirstTime) {
      store.dispatch("session/saveRoute", to);
    }
    nextResolve = {
      name: EVENTS_ROUTE_NAME
    };
  }
  next(nextResolve);
};

router.beforeEach(
  (
    to: Route,
    from: Route,
    next: (to?: RawLocation | false | ((vm: Vue) => any) | void) => void
  ) => {
    const userRequested = store.getters["session/userRequested"];
    if (!userRequested) {
      store.dispatch("session/retrieveUser").then(() => {
        // the first route will be setted here
        const user = store.getters["session/user"];
        routerSecurityCheck(user, to, from, true, next);
      });
      return;
    }
    const user = store.getters["session/user"];
    routerSecurityCheck(user, to, from, false, next);
  }
);

export default router;
