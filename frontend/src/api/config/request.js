import { reqClient, AUTH_URL, WEATHER_URL } from "./base.js";

export const request = async (base, url, method, options) => {
    const onSuccess = (response) => {
        return response;
    };
    const onError = async (error) => {
        if (error?.response?.status === 401) {
            return "Unauthorized";
        }
        if (error?.response?.status === 403) {
            return "Forbidden";
        }

        return Promise.reject(error?.response || error);
    };

    return reqClient(method, base, url, options).then(onSuccess).catch(onError);
};

export const authRequest = async ({ url, method, options }) =>
    request(AUTH_URL, url, method, options);
export const weatherRequest = async ({ url, method, options }) =>
    request(WEATHER_URL, url, method, options);

