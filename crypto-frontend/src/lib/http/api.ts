import moment from "moment";
import _ from "lodash";
export const apiPrefix = "/api/v1";

const USER_TOKEN_KEY = "ppv-user-token";
const USER_TOKEN_EXPIRY = "ppv-user-token-expiry";

let token: string | null | undefined = localStorage.getItem(USER_TOKEN_KEY);

export interface GenericAPIResponse<T> {
  code: number;
  message: string;
  error: Boolean;
  version: string;
  data?: T;
}

export const setToken = (tokenStr: string | null, expiry?: string) => {
  if (tokenStr) {
    localStorage.setItem(USER_TOKEN_KEY, tokenStr);
    if (expiry) {
      localStorage.setItem(USER_TOKEN_EXPIRY, moment(expiry).format());
    }
  } else {
    localStorage.removeItem(USER_TOKEN_KEY);
    localStorage.removeItem(USER_TOKEN_EXPIRY);
  }
  token = tokenStr;
};

export const fetchAuthenticated = function(
  endpoint: string,
  init?: RequestInit
): Promise<Response> {
  const requestInit = Object.assign({}, init);
  if (token && token.length > 0) {
    requestInit.headers = Object.assign({}, requestInit.headers, {
      Authorization: `Bearer ${token}`
    });
  }
  return fetch(`${apiPrefix}/${endpoint}`, requestInit);
};

export const fetchOptinallyAuthenticated = async function(
  endpoint: string,
  init?: RequestInit
): Promise<Response> {
  let response = await fetchAuthenticated(endpoint, init);
  if (response.status === 401 && init && init.headers) {
    setToken(null);
    const withoutAuth = Object.assign({}, init, {
      headers: _.omit(init.headers, "Authorization")
    });
    response = await fetchAuthenticated(endpoint, withoutAuth);
  }
  return response;
};
