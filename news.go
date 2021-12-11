package grpc_server

type News struct {
	Id       int    `json:"id"`
	Category string `json:"category"`
	Datetime int    `json:"datetime"`
	Headline string `json:"headline"`
	Image    string `json:"image"`
	Related  string `json:"related"`
	Resource string `json:"resource"`
	Summary  string `json:"summary"`
	Url      string `json:"url"`
}
