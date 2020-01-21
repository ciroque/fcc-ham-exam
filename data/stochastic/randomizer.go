package stochastic

import (
	"fcc-ham-exam/data/models"
	"math/rand"
)

type Randomization interface {
	SelectRandomQuestion() *models.Question
}

type Randomizer struct {
	QuestionPool *models.QuestionPool
}

func (r *Randomizer) SelectRandomQuestion() (*models.FullyQualifiedQuestion) {
	subElementCount := len(r.QuestionPool.SubElements)
	subElementIndex := rand.Intn(subElementCount)
	subElement := r.QuestionPool.SubElements[subElementIndex]

	groupCount := len(subElement.Groups)
	groupIndex := rand.Intn(groupCount)
	group := subElement.Groups[groupIndex]

	questionCount := len(group.Questions)
	questionIndex := rand.Intn(questionCount)
	question := group.Questions[questionIndex]

	return &models.FullyQualifiedQuestion{
		SubElementTitle: subElement.Title,
		GroupTitle:      group.Title,
		Question:        question,
	}
}
