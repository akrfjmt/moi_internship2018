package web

type GameResponse struct {
	Id string `json:"id"`
	Level int
}

func GetEmptyGameResponse() GameResponse {
	return GameResponse{
		Id:"",
		Level: 0,
	}
}
