package leonidas

import (
	log "github.com/sirupsen/logrus"
)

const ERROR = "ERROR"
const INFO = "INFO"
const DEBUG = "DEBUG"
const WARN = "WARN"

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

func Logging(level string, uutid interface{}, msg string) {
	funName, line, fileName := WhereAmI()
	fields := log.Fields{
		"funcName":   funName,
		"lineNumber": line,
		"uutid":      uutid,
		"fileName":   fileName,
	}
	switch level {
	case "DEBUG":
		log.WithFields(fields).Debug()
	case "INFO":
		log.WithFields(fields).Info(msg)
	case "ERROR":
		log.WithFields(fields).Error(msg)
	}

}
