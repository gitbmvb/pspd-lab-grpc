import React, { useState } from 'react'
import './Sidebar.css'
import { assets } from '../../assets/assets'

const Sidebar = () => {

    const [extended, setExtended] = useState(false)

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
                <div className="bottom-item recent-entry">
                    <img src={assets.question_icon} alt="" />
                    {extended?<p>Ajuda</p>:null}
                </div>
                <div className="bottom-item recent-entry">
                    <img src={assets.history_icon} alt="" />
                    {extended?<p>Histórico</p>:null}
                </div>
                <div className="bottom-item recent-entry">
                    <img src={assets.setting_icon} alt="" />
                    {extended?<p>Configurações</p>:null}
                </div>
            </div>
        </div>
    )
}

export default Sidebar