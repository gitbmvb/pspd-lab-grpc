import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'

const Dashboard = () => {
    const [user, setUser] = useState(null)
    const navigate = useNavigate()

    useEffect(() => {
        const email = sessionStorage.getItem('userEmail')
        if (!email) {
            navigate('/login')
            return
        }

        const fetchUser = async () => {
            try {
                const res = await fetch(`http://localhost:8080/api/users/${email}`)
                const json = await res.json()

                if (json.status !== 'ok') throw new Error('Failed to load user')

                setUser(json.data)
            } catch (err) {
                console.error('Error loading user:', err)
                navigate('/login')
            }
        }

        fetchUser()
    }, [navigate])

    const handleLogout = () => {
        sessionStorage.removeItem('userEmail')
        navigate('/login')
    }

    if (!user) return <div>Loading...</div>

    return (
        <div>
            <h1>Welcome, {user.name}!</h1>
            <p>Your email: {user.email}</p>
            <button onClick={handleLogout}>Logout</button>
        </div>
    )
}

export default Dashboard
