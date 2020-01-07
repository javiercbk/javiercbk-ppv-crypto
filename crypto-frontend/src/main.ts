import Vue from "vue";
import VueCompositionApi from "@vue/composition-api";
import hooks from '@u3u/vue-hooks';
import Buefy from "buefy";
import { library } from '@fortawesome/fontawesome-svg-core'
import { fab } from '@fortawesome/free-brands-svg-icons'
import { fas as freeFas } from '@fortawesome/free-solid-svg-icons'
import { far as freeFar } from '@fortawesome/free-regular-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import App from "./app.vue";
import router from "./router";
import store from "./store";


library.add(fab)
library.add(freeFas)
library.add(freeFar)
Vue.config.productionTip = false;
Vue.use(hooks)
Vue.use(VueCompositionApi);
Vue.component('font-awesome-icon', FontAwesomeIcon)
Vue.use(Buefy, {
  defaultIconComponent: FontAwesomeIcon,
  defaultIconPack: 'fas',
});

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
