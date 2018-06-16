package helper

import (
	log "github.com/sirupsen/logrus"
)

func ReportError(msg string, err error, v ...interface{}) {
	log.Error(msg, err, v)
}

func ReportFatal(msg string, err error, v ...interface{}) {
	log.Fatal(msg, err, v)
}
