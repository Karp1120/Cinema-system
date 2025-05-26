package models

type Film struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Genre    string `json:"genre"`
	Studio   string `json:"studio"`

	// Расширяем по требованию задания
	Operator string `json:"operator"`
	Actors   string `json:"actors"`
}
