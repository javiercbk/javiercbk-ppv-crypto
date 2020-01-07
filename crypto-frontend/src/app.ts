import { createComponent } from "@vue/composition-api";
import AppHeader from "@/components/header.vue";
import AppFooter from "@/components/footer.vue";

export default createComponent({
  components: {
    AppHeader,
    AppFooter
  }
});
