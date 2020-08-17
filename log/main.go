package main

import (
	"log/ali"
	"log/mytest"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	log.SetReportCaller(true)
	log.SetLevel(log.InfoLevel)
}
func main() {
	// logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	// logger.Info().Msg("export GOPROXY=https://goproxy.io")

	log.Info("this is a test")
	log.Error("this is an error")
	ali.Alitest()
	mytest.Alitest1()
}
