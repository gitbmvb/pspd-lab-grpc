import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { loginUser } from '../api'
import './Login.css';

const Login = () => {
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const navigate = useNavigate()

    const handleSubmit = async (e) => {
        e.preventDefault()
        try {
            const data = await loginUser(email, password)
            localStorage.setItem('token', data.token || 'sample')
            await loginUser(email, password)
            sessionStorage.setItem('userEmail', email)
            alert('Login successful')
            navigate('/dashboard')
        } catch (err) {
            alert(err.message)
        }
    }

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
                        onChange={e => setEmail(e.target.value)}
                        required
                    />
                    <input
                        type="password"
                        placeholder="password"
                        value={password}
                        onChange={e => setPassword(e.target.value)}
                        required
                    />
                    <button type="submit">Entrar</button>
                </form>
                <p className="bottom-info">UnB/FGA@2025.1</p>
            </div>
        </div>
    )
}

export default Login