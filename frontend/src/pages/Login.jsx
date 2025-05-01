import React, { useState } from 'react';
import './Login.css';

const Login = () => {
    const [email, setEmail] = useState('');
    const [senha, setSenha] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault();
        console.log('Email:', email);
        console.log('Senha:', senha);
    };

    return (
        <div className="login">
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
                        placeholder="Senha"
                        value={senha}
                        onChange={(e) => setSenha(e.target.value)}
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
