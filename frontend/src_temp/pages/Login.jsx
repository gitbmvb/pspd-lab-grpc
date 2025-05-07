import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { loginUser } from '../api'

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
        <form onSubmit={handleSubmit}>
            <h2>Login</h2>
            <input placeholder="Email" value={email} onChange={e => setEmail(e.target.value)} />
            <input placeholder="Password" type="password" value={password} onChange={e => setPassword(e.target.value)} />
            <button type="submit">Login</button>
        </form>
    )
}

export default Login