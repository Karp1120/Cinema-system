package models

type Cinema struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Category string `json:"category"`
	Halls    int    `json:"halls"`
	Seats    int    `json:"seats"`
	Status   string `json:"status"`
}
