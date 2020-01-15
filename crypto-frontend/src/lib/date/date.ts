import moment from "moment";

export const datetimeFormatter = function(d: Date): string {
  return moment(d).format("YYYY-MM-DDTHH:mm:ss");
};

export const datetimeParser = function(s: string): Date {
  return moment(s, "YYYY-MM-DDTHH:mm:ss", true).toDate();
};
