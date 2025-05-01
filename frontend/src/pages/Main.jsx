import React, { useContext } from 'react'
import './Main.css'
import { assets } from '../assets/assets'
import { Context } from '../context/Context'

const Main = () => {
    const { onSent, recentPrompt, showResult, loading, resultData, setInput, input } = useContext(Context)

    return (
        <div className='main'>
            <div className="nav">
                <p>PSPD Labs</p>
            </div>
            <div className="main-container">
                {!showResult ?
                    <>
                        <div className="greet">
                            <p><span>Olá, FW Cruz.</span></p>
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
                            {loading ?
                            <div className="loader">
                                <hr />
                                <hr />
                                <hr />
                            </div>
                            :
                            <p dangerouslySetInnerHTML={{ __html: resultData }}></p>
                        }
                        </div>
                    </div>
                }
                <div className="main-bottom">
                    <div className="search-box">
                        <input onChange={(e) => setInput(e.target.value)} value={input} type="text" placeholder='Pergunte alguma coisa' />
                        <div>
                            <img onClick={() => onSent()} src={assets.send_icon} alt="" />
                        </div>
                    </div>
                    <p className="bottom-info">UnB/FGA@2025.1</p>
                </div>
            </div>
        </div>
    )
}

export default Main