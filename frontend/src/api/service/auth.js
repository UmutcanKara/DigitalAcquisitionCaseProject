import {authRequest} from "../config/request.js";

export class AuthService {
    static async login(data){
        return await authRequest({
            url: "login",
            method: "POST",
            options: data
        })
    }

    static async register(data) {
        return await authRequest({
            url: "register",
            method: "POST",
            options: data
        })
    }

    static async getUser(usernameQuery){
        return await authRequest({
            url: `protected/?username=${usernameQuery}`,
            method: "GET"
        })
    }

    static async logout() {
        return await authRequest({
            url: "protected/logout",
            method: "POST"
        })
    }
}