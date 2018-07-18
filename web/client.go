package web

import (
	"net/http"
	"encoding/json"
	"bytes"
	"os"
	"strconv"
	"github.com/akrfjmt/moi_internship2018/util"
)

type Client struct {
	Endpoint string
	Token string
	Game  GameResponse
}

func NewClient(endpoint string, token string) *Client {
	return &Client {
		Endpoint: endpoint,
		Token: token,
		Game:  GetEmptyGameResponse(),
	}
}

func (c Client) PostQuestion(question QuestionRequest) QuestionResponse {
	client := &http.Client{}

	jsonStr, err := json.Marshal(question)
	util.Perror(err)

	req, err := http.NewRequest("POST", c.Endpoint + "/games/" + c.Game.Id, bytes.NewBuffer(jsonStr))
	if err != nil {
		os.Exit(1)
	}

	req.Header.Add("X-Api-Version", "2.0")
	req.Header.Add("Authorization", "Bearer " + c.Token)

	resp, err := client.Do(req)
	util.Perror(err)

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var questionResponse QuestionResponse

	err = decoder.Decode(&questionResponse)
	util.Perror(err)

	questionResponse.Level = c.Game.Level
	return questionResponse
}

func (c Client) RequestGame(level int) GameResponse {
	client := &http.Client{}

	req, err := http.NewRequest("GET", c.Endpoint + "/games?level=" + strconv.Itoa(level), nil)
	if err != nil {
		os.Exit(1)
	}

	req.Header.Add("X-Api-Version", "2.0")
	req.Header.Add("Authorization", "Bearer " + c.Token)

	resp, err := client.Do(req)
	util.Perror(err)

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var game GameResponse

	err = decoder.Decode(&game)
	util.Perror(err)

	game.Level = level
	return game
}

func (c *Client) StartGame(level int) {
	c.Game = c.RequestGame(level)
}
