import Vue, { AsyncComponent, Component } from "vue";
import VueRouter, { RouteConfig } from "vue-router";

Vue.use(VueRouter);

const EventsComponent: AsyncComponent = () =>
  import(/* webpackChunkName: "events" */ "../views/events.vue");

const LoginComponent: AsyncComponent = () =>
  import(/* webpackChunkName: "login" */ "../views/login.vue");

const routes: RouteConfig[] = [
  {
    path: "/",
    name: "events",
    component: EventsComponent
  },
  {
    path: "/login",
    name: "login",
    component: LoginComponent
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
