package main

import (
	"basic-gin-app/configs"
	"basic-gin-app/internal/routes"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

func init() {
}

func main() {
	configs.InitDb()
	logInitialization()
	var route routes.Routes
	route.SetupRouter()
}

func logInitialization() {
	var logPath = "logs/application.log"
	file, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.FileMode(0666))
	if err == nil {
		wrt := io.MultiWriter(os.Stdout, &lumberjack.Logger{
			Filename: logPath,
			MaxSize:  5,
			MaxAge:   90,
		})
		log.SetOutput(wrt)
		log.Println(file.Name(), "is generated log file")
	} else {
		log.Info("Error:", err)
		log.Info("Failed to log to file, using default stderr")
	}

	log.SetReportCaller(false)
	log.SetLevel(log.DebugLevel)
}

func doStuff() string {
	return "I do stuff!"
}
