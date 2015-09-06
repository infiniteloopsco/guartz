package utils

import (
	"io"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/rifflock/lfshook"
)

//LogWriter for logging
var LogWriter io.Writer
var Log *log.Logger

//InitLog system
func InitLog() {
	log.SetOutput(os.Stderr)
	log.SetLevel(log.WarnLevel)
	Log = log.New()
	LogWriter = Log.Writer()
	Log.Formatter = new(log.JSONFormatter)
	Log.Hooks.Add(lfshook.NewHook(lfshook.PathMap{
		log.InfoLevel:  "log/info.log",
		log.ErrorLevel: "log/error.log",
	}))
}

//InitLogTest system
func InitLogTest() {
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)
	Log = log.New()
	LogWriter = Log.Writer()
}
