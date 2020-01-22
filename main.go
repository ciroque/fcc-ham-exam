package main

import (
	"encoding/json"
	"fcc-ham-exam/config"
	"fcc-ham-exam/data/models"
	"fcc-ham-exam/data/stochastic"
	"fcc-ham-exam/http"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.Info("Starting...")

	abortChannel := make(chan string)
	defer close(abortChannel)

	settings, err := config.NewSettings()
	if err != nil {
		logrus.Fatalf("Unable to load settings: %v", err)
	}

	logrus.Infof("Loaded settings: %v", settings)

	technicianQuestionPool, err := LoadTechnicianQuestions(settings.DataFilePath)
	if err != nil {
		logrus.Fatalf("Unable to load Technician question pool: %v", err)
	}

	randomizer := stochastic.Randomizer{QuestionPool: technicianQuestionPool}

	for i := 0; i < 10; i++ {
		question := randomizer.SelectRandomQuestion()
		logrus.Infof("---> %#v", question)
	}

	server := http.Server{
		AbortChannel: abortChannel,
		Logger:       logrus.NewEntry(logrus.New()),
		Settings:     settings,
		Randomizer:   &randomizer,
	}

	go server.Run()

	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, syscall.SIGTERM)
	signal.Notify(sigTerm, syscall.SIGINT)

	select {
	case <-sigTerm:
		{
			logrus.Info("Exiting per SIGTERM")
		}
	case err := <-abortChannel:
		{
			logrus.Error(err)
		}
	}
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
