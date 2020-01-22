package models

type Answer struct {
	Text string			`json:"text"`
	Correct bool		`json:"correct"`
}
