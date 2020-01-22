package models

type Group struct {
	Title string		`json:"title"`
	Questions []Question `json:"questions"`
}
