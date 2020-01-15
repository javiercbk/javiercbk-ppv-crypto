export const apiPrefix = "/api/v1";

const USER_TOKEN_KEY = "ppv-user-token";

let token: string | null | undefined = localStorage.getItem(USER_TOKEN_KEY);

export interface GenericAPIResponse<T> {
  code: number;
  message: string;
  error: Boolean;
  version: string;
  data?: T;
}

export const setToken = (tokenStr?: string) => {
  if (tokenStr) {
    localStorage.setItem(USER_TOKEN_KEY, tokenStr);
  } else {
    localStorage.removeItem(USER_TOKEN_KEY);
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
