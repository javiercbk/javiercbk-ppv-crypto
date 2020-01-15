import { createComponent, computed, Ref } from "@vue/composition-api";
import { useState } from "@u3u/vue-hooks";
import { User } from "@/models/models";

export default createComponent({
  setup() {
    const state = {
      ...useState("session", ["user"])
    };

    const computedProps = {
      isLoggedIn: computed(() => {
        const user = (state.user as Ref<User | null>).value;
        return user && user.id > 0;
      })
    };

    return {
      ...computedProps
    };
  }
});
