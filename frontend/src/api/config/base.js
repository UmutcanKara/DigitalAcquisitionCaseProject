import axios from "axios";

export const AUTH_URL = "http://127.0.0.1:8080";
export const WEATHER_URL = "http://127.0.0.1:8081";

export const reqClient = (method, base, url, options) => {
    switch (method) {
        case "GET":
            return axios.get(`${base}/${url}`, {
                ...options,
                headers: {
                    Accept: "*/*",
                    "Content-Type": "*/*",
                },
                withCredentials: true,

            });
        case "POST":
            return axios.post(`${base}/${url}`, options, {
                headers: {
                    Accept: "application/json",
                    "Content-Type": "application/json",
                },
                withCredentials: true
            });
        case "PUT":
            return axios.put(`${base}/${url}`, options, {
                headers: {
                    Accept: "application/json",
                    "Content-Type": "application/json",
                },
                withCredentials: true
            });
    }
};