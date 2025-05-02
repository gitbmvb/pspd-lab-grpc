import React, { useContext } from 'react'
import './Main.css'
import { assets } from '../assets/assets'
import { Context } from '../context/Context'
import PromptBox from '../components/PromptBox/PromptBox'
import Markdown from 'react-markdown';
import remarkGfm from 'remark-gfm'

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
    )
}

export default Main
