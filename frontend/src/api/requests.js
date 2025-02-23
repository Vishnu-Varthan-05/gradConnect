import axios from "axios";
import {jwtDecode} from "jwt-decode";

const url = "http://localhost:5000/";

export function login(data) {
    return axios.post(url + "login", data)
        .then(res => {
            if (res.data.token) {
                localStorage.setItem("token", res.data.token);
                return res.data; 
            }
        })
        .catch(err => {
            console.log("Error logging in", err);
            throw err; 
        });
}

export function logout() {
    localStorage.clear();
}


export function getUserInfo() {
    const token = localStorage.getItem("token");  
    if (!token) return null;  

    try {
        const decoded = jwtDecode(token);
        return {
            user_id: decoded.userId || null,
            role: decoded.roles[0] || "no",
            fullname: decoded.fullname || "",
        };
    } catch (error) {
        console.error("Invalid token", error);
        return null;
    }
}
