import React, { useState, useEffect } from "react";
import "./App.css";
import translations from "./translations";

function App() {
  const [colors, setColors] = useState([]);
  const [correctColor, setCorrectColor] = useState("");
  const [message, setMessage] = useState("");
  const [language, setLanguage] = useState("ja");
  const [correctCount, setCorrectCount] = useState(0);
  const [questionCount, setQuestionCount] = useState(0);
  const [gameComplete, setGameComplete] = useState(false);

  const fetchNewQuestion = (currentCorrectCount = correctCount) => {
    // 問題数をインクリメント
    const newQuestionCount = questionCount + 1;
    setQuestionCount(newQuestionCount);
    
    // 10問解答したらゲーム完了
    if (newQuestionCount >= 10) {
      setGameComplete(true);
      return;
    }
    
    // 5問正解したら選択肢を4つに増やす
    const optionCount = currentCorrectCount >= 5 ? 4 : 3;
    
    // APIエンドポイントを環境変数から取得
    const apiUrl = process.env.REACT_APP_API_URL || "/api/color";
    
    fetch(`${apiUrl}?count=${optionCount}`)
      .then((res) => res.json())
      .then((data) => {
        setColors(data.options);
        setCorrectColor(data.correct);
        setMessage("");
        setIsCorrect(false);
        setAnswered(false);
      });
  };
  
  const resetGame = () => {
    // グローバル変数もリセット
    window.lastCorrectCount = 0;
    
    setCorrectCount(0);
    setQuestionCount(0);
    setGameComplete(false);
    setMessage("");
    setIsCorrect(false);
    setAnswered(false);
    
    // APIエンドポイントを環境変数から取得
    const apiUrl = process.env.REACT_APP_API_URL || "/api/color";
    
    // fetchNewQuestionを直接呼ぶと問題数が増えるので、APIを直接呼び出す
    const optionCount = 3;
    fetch(`${apiUrl}?count=${optionCount}`)
      .then((res) => res.json())
      .then((data) => {
        setColors(data.options);
        setCorrectColor(data.correct);
      });
  };

  useEffect(() => {
    fetchNewQuestion();
  }, []);

  const [isCorrect, setIsCorrect] = useState(false);

  const [answered, setAnswered] = useState(false);

  const handleGuess = (color) => {
    setAnswered(true);
    if (color === correctColor) {
      setMessage(translations[language].correct);
      setIsCorrect(true);
      const newCount = correctCount + 1;
      setCorrectCount(newCount);
      
      // 正解数を更新した後、ローカル変数に保存しておく
      window.lastCorrectCount = newCount;
    } else {
      setMessage(translations[language].incorrect);
      setIsCorrect(false);
    }
    
    // 最後の問題の場合はここでゲーム完了にする
    if (questionCount >= 9) {
      setGameComplete(true);
    }
  };
  
  const toggleLanguage = () => {
    setLanguage(language === "ja" ? "en" : "ja");
  };

  const colorMap = {
    "赤": "red",
    "青": "blue",
    "黄": "yellow",
    "緑": "green",
    "紫": "purple",
    "オレンジ": "orange"
  };

  return (
    <div className="container">
      <div className="language-selector">
        <button onClick={toggleLanguage} className="language-button">
          {language === "ja" ? "English" : "日本語"}
        </button>
      </div>
      <h1>{translations[language].title}</h1>
      
      {gameComplete ? (
        <div className="game-complete">
          <h2>{translations[language].gameComplete}</h2>
          <p>{translations[language].correctCount.replace('{count}', correctCount)}</p>
          <button 
            onClick={resetGame} 
            className="play-again-button"
          >
            {translations[language].playAgain}
          </button>
        </div>
      ) : (
        <>
          <div 
            className="color-box"
            style={{ backgroundColor: colorMap[correctColor] || correctColor }}
          ></div>
          <div className="button-container">
            {colors.map((color, index) => (
              <button 
                key={index} 
                onClick={() => handleGuess(color)}
                className="color-button"
                disabled={answered}
              >
                {translations[language].colorNames[color] || color}
              </button>
            ))}
          </div>
          <p className="message">{message}</p>
          {answered && !gameComplete && (
            <button 
              onClick={() => fetchNewQuestion(window.lastCorrectCount || correctCount)} 
              className="next-button"
            >
              {translations[language].next}
            </button>
          )}
        </>
      )}
    </div>
  );
}

export default App;