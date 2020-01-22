package stochastic

import (
	"fcc-ham-exam/data/models"
	"math/rand"
	"time"
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

	fullyQualifiedQuestion := &models.FullyQualifiedQuestion{
		SubElementTitle: subElement.Title,
		GroupTitle:      group.Title,
		Question:        question,
	}

	ShuffleAnswers(fullyQualifiedQuestion)

	return fullyQualifiedQuestion
}

func (r *Randomizer) ShuffleAnswers(question *models.FullyQualifiedQuestion) {
	swap := func(i, j int) {
		question.Answers[i],
		question.Answers[j] = question.Answers[i],
		question.Answers[i]
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(question.Answers), swap)
}
