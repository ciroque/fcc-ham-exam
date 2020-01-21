package main

import (
	"encoding/json"
	"fcc-ham-exam/config"
	"fcc-ham-exam/data/models"
	"fmt"
	"github.com/Sirupsen/logrus"
	"io/ioutil"
)

func main() {
	logrus.Info("Starting...")
	settings, err := config.NewSettings()
	if err != nil {
		logrus.Fatalf("Unable to load settings: %v", err)
	}

	technicianQuestionPool, err := LoadTechnicianQuestions(settings.DataFilePath)
	if err != nil {
		logrus.Fatalf("Unable to load Technician question pool: %v", err)
	}

	logrus.Infof("Technician question pool: %#v", technicianQuestionPool)
}

func LoadTechnicianQuestions(rootPath string) (*models.QuestionPool, error) {
	file, err := ioutil.ReadFile(rootPath + "technician-questions.json")
	if err != nil {
		return nil, fmt.Errorf("unable to load Technician quesions file: %v", err)
	}

	technicianQuestionPool := models.QuestionPool{}

	err = json.Unmarshal([]byte(file), &technicianQuestionPool)
	if err != nil {
		return nil, fmt.Errorf("unable to parse Technician questions: %v", err)
	}

	return &technicianQuestionPool, nil
}
