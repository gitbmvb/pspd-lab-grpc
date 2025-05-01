import React from 'react'
import './Main.css'
import { assets } from '../assets/assets'

const Main = () => {
    return (
        <div className='main'>
            <div className="nav">
                <p>PSPD Labs</p>
            </div>
            <div className="main-container">
                <div className="greet">
                    <p><span>Olá, FW Cruz.</span></p>
                    <p>Como posso te ajudar hoje?</p>
                </div>
                <div className="main-bottom">
                    <div className="search-box">
                        <input type="text" placeholder='Pergunte alguma coisa' />
                        <div>
                            <img src={assets.send_icon} alt="" />
                        </div>
                    </div>
                    <p className="bottom-info">UnB/FGA@2025.1</p>
                </div>
            </div>
        </div>
    )
}

export default Main