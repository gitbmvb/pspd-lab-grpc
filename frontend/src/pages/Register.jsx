import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { registerUser } from '../api'
import './Register.css';

const Register = () => {
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [confirmPassword, setConfirmPassword] = useState('')
  const [errorMessage, setErrorMessage] = useState('')
  const navigate = useNavigate()

  const handleSubmit = async (e) => {
    e.preventDefault()
    setErrorMessage('')

    if (password !== confirmPassword) {
      setErrorMessage('As senhas não coincidem.')
      return
    }

    try {
      await registerUser(name, email, password)
      navigate('/login', { state: { success: 'Cadastro realizado com sucesso!' } })
    } catch (err) {
      setErrorMessage('Erro ao criar conta. Verifique os dados e tente novamente.')
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
          <p><span>Crie uma conta</span></p>
          <p>É rápido e fácil!</p>
        </div>
        <form className="login-form" onSubmit={handleSubmit}>
          <input
            type="text"
            placeholder="Seu nome"
            value={name}
            onChange={e => setName(e.target.value)}
            required
          />
          <input
            type="email"
            placeholder="Email"
            value={email}
            onChange={e => setEmail(e.target.value)}
            required
          />
          <input
            type="password"
            placeholder="Senha"
            value={password}
            onChange={e => setPassword(e.target.value)}
            required
          />
          <input
            type="password"
            placeholder="Repetir senha"
            value={confirmPassword}
            onChange={e => setConfirmPassword(e.target.value)}
            required
          />
          <button type="submit">Entrar</button>
        </form>

        {errorMessage && (
          <p className="error-message">{errorMessage}</p>
        )}

        <p className="create-account-link">
          Já possui uma conta? <a href="/login">Faça login</a>.
        </p>
        <p className="bottom-info">UnB/FGA@2025.1</p>
      </div>
    </div>
  )
}

export default Register
