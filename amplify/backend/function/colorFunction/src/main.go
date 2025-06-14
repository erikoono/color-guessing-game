package main

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type ColorResponse struct {
	Options []string `json:"options"`
	Correct string   `json:"correct"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// クエリパラメータから選択肢の数を取得
	count := 3
	countParam := request.QueryStringParameters["count"]
	if countParam == "4" {
		count = 4
	}

	// 色のリスト
	allColors := []string{"赤", "青", "黄", "緑", "紫", "オレンジ"}

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

	// JSONに変換
	jsonResponse, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		Body: string(jsonResponse),
	}, nil
}

func main() {
	lambda.Start(handler)
}
