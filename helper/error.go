package helper

import (
	log "github.com/sirupsen/logrus"
)

type Errors struct {
	Errors []Error `json:"errors"`
}

type Error struct {
	Message string `json:"message"`
}

func ReportError(msg string, err error, v ...interface{}) {
	log.Error(msg, err, v)
}

func ReportFatal(msg string, err error, v ...interface{}) {
	log.Fatal(msg, err, v)
}
