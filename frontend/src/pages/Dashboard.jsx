import { useEffect, useState, useContext } from 'react'
import { useNavigate } from 'react-router-dom'
import { Context } from '../context/ContextProvider'
import { assets } from '../assets/assets'
import Sidebar from '../components/Sidebar/Sidebar'
import PromptBox from '../components/PromptBox/PromptBox'
import Markdown from 'react-markdown';
import remarkGfm from 'remark-gfm'
import './Dashboard.css'

const Dashboard = () => {
    const [user, setUser] = useState(null)
    const navigate = useNavigate()
    const { onSent, recentPrompt, showResult, loading, resultData, setInput, input } = useContext(Context)

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
        // <div>
        //     <h1>Welcome, {user.name}!</h1>
        //     <p>Your email: {user.email}</p>
        //     <button onClick={handleLogout}>Logout</button>
        // </div>
        <>
            <Sidebar />
            <div className='main'>
                <div className="nav">
                    <p>PSPD Labs</p>
                </div>
                <div className="main-container">
                    {!showResult ?
                        <>
                            <div className="greet">
                                <p><span>Olá, {user.name}.</span></p>
                                <p>Como posso te ajudar hoje?</p>
                            </div>
                        </>
                        :
                        <div className="result">
                            <div className="result-title">
                                <img src={assets.user_icon} alt="" />
                                <p>{recentPrompt}</p>
                            </div>
                            <div className="result-data">
                                <img src={assets.gemini_icon} alt="" />
                                <div className="markdown-container">
                                    {loading ? (
                                        <div className="loader">
                                            <hr />
                                            <hr />
                                            <hr />
                                        </div>
                                    ) : (
                                        <Markdown remarkPlugins={[remarkGfm]}>{resultData}</Markdown>
                                    )}
                                </div>
                            </div>
                        </div>
                    }
                    <div className="main-bottom">
                        <PromptBox
                            onSent={onSent}
                            setInput={setInput}
                            input={input}
                        />
                        <p className="bottom-info">UnB/FGA@2025.1</p>
                    </div>
                </div>
            </div>
        </>
    )
}

export default Dashboard
