import {weatherRequest} from "../config/request.js";

export class WeatherService {
    static async getWeather(city, start_date){
        return await weatherRequest({
            url: `protected/?city=${city}&start_date=${start_date}`,
            method: "GET"
        })
    }

    static async updateWeather(city) {
        return await weatherRequest({
            url: `protected/?city=${city}`,
            method: "PUT"
        })
    }
}