// src/pages/Register.jsx
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { registerUser } from '../api'

const Register = () => {
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const navigate = useNavigate()

  const handleSubmit = async (e) => {
    e.preventDefault()
    try {
      await registerUser(name, email, password)
      navigate('/login')
    } catch (err) {
      alert(err.message) // or setError(err.message) to show in UI
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      <h2>Register</h2>
      <input placeholder="Name" value={name} onChange={e => setName(e.target.value)} />
      <input placeholder="Email" value={email} onChange={e => setEmail(e.target.value)} />
      <input placeholder="Password" type="password" value={password} onChange={e => setPassword(e.target.value)} />
      <button type="submit">Register</button>
    </form>
  )
}

export default Register
