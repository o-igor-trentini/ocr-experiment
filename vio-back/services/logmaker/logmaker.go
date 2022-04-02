package logmaker

import (
	"log"
	"vio-back/services/violog"
)

type Service interface {
	InfoLogger(err error, errorMessage string, errorStatus int64) error
	ErrorLogger(err error, errorMessage string, errorStatus int64) error
}

func InfoLogger(message string, args ...interface{}) {
	if violog.InfoLogger != nil {
		violog.InfoLogger.Printf(message, args...)
		return
	}

	log.Default().Printf(message, args...)
}

func ErrorLogger(message string, args ...interface{}) {
	if violog.ErrorLogger != nil {
		violog.ErrorLogger.Printf(message, args...)
		return
	}

	log.Default().Printf(message, args...)
}
