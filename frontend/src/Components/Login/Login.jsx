import styles from "./Login.module.css"
import image from "../../assets/login.png"
import { useState } from "react";

export default function Login(){
    const [loginForm, setLoginForm] = useState({
        email:"",
        password:""
    })
    function handleChange(e) {
        setLoginForm({
            ...loginForm,
            [e.target.name]:e.target.value
        })
    }
    function handleSubmit(e) {
        e.preventDefault();
        console.log(loginForm);
    }
    return (
        <div className={styles.login}>
            <div className={styles.loginContainer}>
                <img src={image} alt="" />
                <form onSubmit={handleSubmit}>
                    <p>Login</p>
                    <input type="email" placeholder="Enter your email" name="email" value={loginForm.email} required onChange={handleChange}/>
                    <input type="password" placeholder="Enter your Password" name="password" value={loginForm.password} required onChange={handleChange}/>
                    <button type="submit" >Login</button>
                </form>
            </div>
        </div>
    );
}