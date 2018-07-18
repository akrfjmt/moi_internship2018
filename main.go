package main

import (
	_ "net/http/pprof"
	"github.com/akrfjmt/moi_internship2018/web"
	"github.com/akrfjmt/moi_internship2018/hb"
	"fmt"
	"os"
	"math/rand"
)

func main() {
	endpoint := "https://*/internships/2018"
	token := "token"

	fmt.Println("start")
	hbService := hb.HitBlowService{
		Base: 10,
		Digits: 10,
	}

	fmt.Println("質問作成")
	questions := hbService.CreateAllHitBlowNumbers()

	fmt.Println("シャッフル")
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

	client := web.NewClient(
		endpoint,
		token)

	client.StartGame(10)

	for len(questions) > 0 {
		question := questions[0]
		fmt.Println(question)

		questions = questions[1:]

		q := hbService.CreateQuestionRequest(question)
		questionResponse := client.PostQuestion(q)

		if questionResponse.IsClear() {
			clear()
		}

		hbProduct := hbService.CreateHitBlowProduct(questionResponse)

		questions = hbService.FilterQuestions(question, hbProduct, questions)
	}

	fmt.Println("意味不明")
	os.Exit(1)
}

func clear() {
	fmt.Println("clear!")
	os.Exit(0)
}
