import { helpers } from "vuelidate/lib/validators";

export const ethAddress = helpers.regex("ethAddress", /^0x[0-9a-fA-F]{40}$/);
