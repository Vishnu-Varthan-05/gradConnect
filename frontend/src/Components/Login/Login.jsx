import styles from "./Login.module.css";
import image from "../../assets/login.png";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { login } from "../../api/requests";

export default function Login() {
    const navigate = useNavigate();
    const [loginForm, setLoginForm] = useState({ email: "", password: "" });

    useEffect(() => {
        const token = localStorage.getItem("token");
        if (token) {
            navigate("/");
        }
    }, [navigate]);

    function handleChange(e) {
        setLoginForm({ ...loginForm, [e.target.name]: e.target.value });
    }

    function handleSubmit(e) {
        e.preventDefault();
        login(loginForm)
        .then(() => {
            navigate("/"); 
        })
        .catch(() => {
            alert("Invalid email or password");  
        });
    }

    return (
        <div className={styles.login}>
            <div className={styles.loginContainer}>
                <img src={image} alt="Login" />
                <form onSubmit={handleSubmit}>
                    <p>Login</p>
                    <input type="email" placeholder="Enter your email" name="email" value={loginForm.email} required onChange={handleChange}/>
                    <input type="password" placeholder="Enter your Password" name="password" value={loginForm.password} required onChange={handleChange}/>
                    <button type="submit">Login</button>
                </form>
            </div>
        </div>
    );
}
