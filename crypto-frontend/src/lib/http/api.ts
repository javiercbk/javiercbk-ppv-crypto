let token: String = "";

export const apiPrefix = "/api/v1";

export interface GenericAPIResponse<T> {
  code: Number;
  message: String;
  error: Boolean;
  version: String;
  data?: T;
}

export const setToken = (tokenStr: String) => {
  token = tokenStr;
};

export const fetchAuthenticated = function(
  endpoint: String,
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
