package data

type Answer struct {
	Text string			`json:text`
	Correct bool		`json:correct`
}

type Question struct {
	Answers []Answer	`json:answers`
	Number string		`json:number`
	Question string		`json:question`
}

type Group struct {
	Title string		`json:title`
	Questions []Question `json:questions`
}

type SubElements struct {
	Title string		`json:title`
	Groups []Group		`json:groups`
}

type QuestionPool struct {
	SubElements []SubElements
}
