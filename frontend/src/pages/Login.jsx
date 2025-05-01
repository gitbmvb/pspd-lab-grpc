import React, { useState } from 'react';
import './Login.css';

const Login = () => {
    const [email, setEmail] = useState('');
    const [password, setpassword] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        
        const response = await fetch('http://localhost:8080/api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email, password }),
        });

        if (response.ok) {
            alert('Login bem-sucedido!');
        } else {
            alert('Invalid credentials. Please try again.');
        }
    };

    return (
        <div className="login">
            <video autoPlay loop muted className="login-bg-video">
                <source src="/background.mp4" type="video/mp4" />
                Your browser does not support the video tag.
            </video>
            <div className="nav">
                <p>PSPD Labs</p>
            </div>
            <div className="login-container">
                <div className="login-greet">
                    <p><span>Bem-vindo!</span></p>
                    <p>Faça login para continuar</p>
                </div>
                <form className="login-form" onSubmit={handleSubmit}>
                    <input
                        type="email"
                        placeholder="Email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required
                    />
                    <input
                        type="password"
                        placeholder="password"
                        value={password}
                        onChange={(e) => setpassword(e.target.value)}
                        required
                    />
                    <button type="submit">Entrar</button>
                </form>
                <p className="bottom-info">UnB/FGA@2025.1</p>
            </div>
        </div>
    );
};

export default Login;