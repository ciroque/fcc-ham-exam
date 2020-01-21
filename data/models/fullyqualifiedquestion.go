package models

type FullyQualifiedQuestion struct {
	SubElementTitle string `json:subElementTitle`
	GroupTitle string `json:groupTitle`
	Question Question `json:question`
}
