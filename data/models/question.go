package models

type Question struct {
	Answers []Answer	`json:"answers"`
	Number string		`json:"number"`
	Question string		`json:"question"`
}
