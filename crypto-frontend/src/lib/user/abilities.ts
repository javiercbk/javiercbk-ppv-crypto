import { computed, Ref } from "@vue/composition-api";
import { AbilityBuilder, Ability } from "casl";
import { User } from "@/models/models";

export const defineAbilitiesFor = function(user: User | null) {
  const { rules, can } = AbilityBuilder.extract();
  if (user && user.permissions) {
    user.permissions.forEach(({ resource, access }) => {
      can(access, resource);
    });
  }

  return new Ability(rules);
};

export const userHasAccessComputed = (
  userRef: Ref<User | null>,
  access: string,
  resource: string
) =>
  computed(() => {
    const user = userRef.value;
    if (user && user.ability) {
      return user.ability.can(access, resource);
    }
    return false;
  });
