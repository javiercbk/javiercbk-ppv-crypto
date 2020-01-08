import { AbilityBuilder, Ability } from "casl";
import { User } from "@/models/models";

// Alternatively this data can be retrieved from server
export default function defineAbilitiesFor(user: User) {
  const { rules, can } = AbilityBuilder.extract();

  user.permissions.forEach(({ resource, access }) => {
    can(resource, access);
  });

  return new Ability(rules);
}
