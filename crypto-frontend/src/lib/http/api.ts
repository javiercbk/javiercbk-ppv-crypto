let token: string = "";

export const apiPrefix = "/api/v1";

export interface GenericAPIResponse<T> {
  code: number;
  message: string;
  error: Boolean;
  version: string;
  data?: T;
}

export const setToken = (tokenStr: string) => {
  token = tokenStr;
};

export const fetchAuthenticated = function(
  endpoint: string,
  init?: RequestInit
): Promise<Response> {
  const requestInit = Object.assign({}, init);
  if (token.length > 0) {
    requestInit.headers = Object.assign({}, requestInit.headers, {
      Authorization: `Bearer ${token}`
    });
  }
  return fetch(`${apiPrefix}/${endpoint}`, requestInit);
};
