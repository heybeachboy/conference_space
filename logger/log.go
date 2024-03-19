package logger

import (
	"ConferenceSpace/config"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	logNamePrefix = "ConferenceSpace"
)

var logService *logrus.Logger

func init() {
	logService = logrus.New()
	logService.SetLevel(logrus.InfoLevel)
	logService.SetFormatter(&logrus.TextFormatter{})
	/*file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}*/
	logService.SetOutput(os.Stdout)
	//log.Info("Init log successful")
}

func InfoF(format string, args ...interface{}) {
	logService.Infof(format, args...)
}

func ErrorF(format string, args ...interface{}) {
	logService.Errorf(format, args...)
}

func WaringF(format string, args ...interface{}) {
	logService.Warnf(format, args...)
}

func DebugF(format string, args ...interface{}) {
	if !config.Config.IsDebug() {
		return
	}
	logService.Debugf(format, args...)
}

func Close() {
}
