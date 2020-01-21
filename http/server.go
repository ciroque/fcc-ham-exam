package http

import (
	"encoding/json"
	"fcc-ham-exam/config"
	"fcc-ham-exam/data/stochastic"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	AbortChannel chan<- string
	Logger       *logrus.Entry
	Settings     *config.Settings
	Randomizer   *stochastic.Randomizer
}

func (server *Server) Run() {
	http.HandleFunc("/api/questions/technician", server.ServeRandomTechnicianQuestion)
	address := fmt.Sprintf("%s:%d", server.Settings.Host, server.Settings.Port)
	server.Logger.Info("Listening on ", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		server.AbortChannel <- err.Error()
	}
}

func (server *Server) ServeRandomTechnicianQuestion(writer http.ResponseWriter, request *http.Request) {
	question := server.Randomizer.SelectRandomQuestion()
	response, err := json.Marshal(question)
	if err != nil {
		server.Logger.Warnf("Error responding to request %#v", err)
	}

	writer.Header().Add("Content-Type", "application/json")

	_, err = fmt.Fprintf(writer, "%s", response)
	if err != nil {
		server.Logger.Warnf("Error responding to request %#v", err)
	}
}