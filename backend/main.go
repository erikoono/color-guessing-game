package main

import (
  "encoding/json"
  "math/rand"
  "net/http"
  "time"
)

type ColorResponse struct {
  Options []string `json:"options"`
  Correct string   `json:"correct"`
}

func colorHandler(w http.ResponseWriter, r *http.Request) {
  allColors := []string{"赤", "青", "黄", "緑", "紫", "オレンジ"}
  
  // クエリパラメータから選択肢の数を取得
  count := 3
  countParam := r.URL.Query().Get("count")
  if countParam == "4" {
    count = 4
  }
  
  // 選択肢の数に応じてランダムに色を選択
  rand.Shuffle(len(allColors), func(i, j int) {
    allColors[i], allColors[j] = allColors[j], allColors[i]
  })
  
  colors := allColors[:count]
  correct := colors[rand.Intn(len(colors))]
  
  response := ColorResponse{
    Options: colors,
    Correct: correct,
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func main() {
  // 乱数のシードを初期化
  rand.Seed(time.Now().UnixNano())
  
  http.HandleFunc("/api/color", colorHandler)
  http.ListenAndServe(":8080", nil)
}