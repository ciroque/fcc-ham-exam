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

func (r *Randomizer) ShuffleAnswers(fqq *models.FullyQualifiedQuestion) {
	swap := func(i, j int) {
		fqq.Question.Answers[i],
		fqq.Question.Answers[j] = fqq.Question.Answers[i],
		fqq.Question.Answers[i]
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(fqq.Question.Answers), swap)
}
