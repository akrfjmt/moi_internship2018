package web

type QuestionResponse struct {
	Hit int `json:"hit"`
	Blow int `json:"blow"`
	Round int `json:"round"`
	Score int `json:"score,omitempty"`
	Message string `json:"message"`
	Level int
}

func (a QuestionResponse) IsClear() bool {
	return a.Hit == a.Level
}
