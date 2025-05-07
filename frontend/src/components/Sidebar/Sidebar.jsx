import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import './Sidebar.css'
import { assets } from '../../assets/assets'

const Sidebar = () => {
    const [extended, setExtended] = useState(true)
    const navigate = useNavigate()

    const handleLogout = () => {
        localStorage.removeItem('token')
        sessionStorage.removeItem('userEmail')
        navigate('/login')
    }

    return (
        <div className='sidebar'>
            <div className="top">
                <img onClick={() => setExtended(prev => !prev)} className='menu' src={assets.menu_icon} alt="" />
                <div className="new-chat">
                    <img src={assets.plus_icon} alt="" />
                    {extended ? <p>Novo Chat</p> : null}
                </div>
                {extended
                    ? <div className="recent">
                        <p className='recent-title'>Recentes</p>
                        <div className="recent-entry">
                            <img src={assets.message_icon} alt="" />
                            <p>O que é gRPC?</p>
                        </div>
                    </div>
                    : null
                }
            </div>
            <div className="bottom">
                <div className="bottom-item recent-entry" onClick={handleLogout}>
                    <img src={assets.logout_icon} alt="" />
                    {extended ? <p>Sair</p> : null}
                </div>
            </div>
        </div>
    )
}

export default Sidebar