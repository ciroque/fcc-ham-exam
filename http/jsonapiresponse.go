package http

import (
	"fcc-ham-exam/data/models"
)

type JsonApiResponse struct {
	Data *models.FullyQualifiedQuestion	`json:"data"`
}
