import { useState } from 'react';
import './App.css';

function App() {
    const [inputText, setInputText] = useState('');
    const [outputText, setOutputText] = useState('');
    const [keyword, setKeyword] = useState('');

    const handleInputChange = (e) => {
        const value = e.target.value.replace(/[^a-zA-Z]/g, '');
        setInputText(value.toUpperCase());
    };

    const handleKeywordChange = (e) => {
        const value = e.target.value.replace(/[^a-zA-Z]/g, '');
        setKeyword(value.toUpperCase());
    };

    const handleEncrypt = async () => {
        if (!inputText || !keyword) return;
        const result = await window.go.main.App.Encrypt(inputText, keyword);
        setOutputText(result);
    };

    const handleDecrypt = async () => {
        if (!inputText || !keyword) return;
        const result = await window.go.main.App.Decrypt(inputText, keyword);
        setOutputText(result);
    };

    const handleReadFile = async () => {
        const content = await window.go.main.App.OpenFile();

        if (content === "") {
            return;
        }

        if (content.startsWith('Error:')) {
            alert(content);
            return;
        }

        const cleanText = content.replace(/[^a-zA-Z]/g, '').toUpperCase();
        setInputText(cleanText);
    };

    const handleSaveFile = async () => {
        if (outputText) await window.go.main.App.SaveFile(outputText);
    };

    const handleClear = () => {
        setInputText('');
        setOutputText('');
        setKeyword('');
    };

    return (
        <div className="app">
            <h1>🔐 COLUMNAR CIPHER</h1>

            <div className="row">
                <input
                    type="text"
                    value={keyword}
                    onChange={handleKeywordChange}
                    placeholder="KEYWORD"
                    className="keyword"
                    maxLength={20}
                />
            </div>

            <div className="row">
        <textarea
            value={inputText}
            onChange={handleInputChange}
            placeholder="INPUT TEXT"
            className="textarea"
            rows={5}
        />
            </div>

            <div className="buttons">
                <button onClick={handleEncrypt} className="btn encrypt">ENCRYPT</button>
                <button onClick={handleDecrypt} className="btn decrypt">DECRYPT</button>
                <button onClick={handleReadFile} className="btn file">📂 OPEN</button>
                <button onClick={handleSaveFile} className="btn save">💾 SAVE</button>
                <button onClick={handleClear} className="btn clear">🗑️ CLEAR</button>
            </div>

            <div className="row">
        <textarea
            value={outputText}
            readOnly
            placeholder="OUTPUT TEXT"
            className="textarea output"
            rows={5}
        />
            </div>
        </div>
    );
}

export default App;