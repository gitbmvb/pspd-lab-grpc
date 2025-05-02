import React from 'react';
import './PromptBox.css'
import { assets } from '../../assets/assets';

const PromptBox = ({ onSent, input, setInput }) => {

  const handleInput = (e) => {
    const textarea = e.target;
    setInput(e.target.value);
    textarea.style.height = "auto";
    textarea.style.height = `${textarea.scrollHeight}px`;
  };

  const handleKeyDown = (e) => {
    if (e.key === 'Enter' && e.shiftKey) {
      return;
    }
    if (e.key === 'Enter') {
      e.preventDefault();
      if (input.trim()) onSent();
    }
  }

  return (
    <div className="prompt-box">
      <textarea
        value={input}
        onChange={(e) => setInput(e.target.value)}
        onKeyDown={handleKeyDown}
        onInput={handleInput}
        placeholder="Pergunte alguma coisa"
        rows="1"
      />
      <div>
      <button
          onClick={onSent}
          aria-label="Enviar"
          disabled={!input.trim()}
      >
          <img src={assets.send_icon} alt="Enviar" />
      </button>
      </div>
    </div>
  );
}

export default PromptBox;
