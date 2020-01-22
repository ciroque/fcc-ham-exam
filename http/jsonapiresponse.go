package http

import (
	"fcc-ham-exam/data/models"
)

type JsonApiResponse struct {
	Data FullyQualifiedQuestion	`json:"data"`
}
