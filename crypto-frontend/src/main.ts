import Vue from "vue";
import VueCompositionApi from "@vue/composition-api";
import Buefy from "buefy";
import App from "./app.vue";
import router from "./router";
import store from "./store";

Vue.config.productionTip = false;
Vue.use(VueCompositionApi);
Vue.use(Buefy);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
