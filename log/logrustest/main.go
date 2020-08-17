package main

import (
	"log"
	"net/http"
	"os"

	logger "github.com/sirupsen/logrus"
)

func init() {
	writer, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("openfile failed")
	}
	logger.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logger.SetLevel(logger.InfoLevel)
	logger.SetOutput(writer)
	logger.SetReportCaller(true)
}

func main() {
	simpleHttpGet("www.sogo.com")
	simpleHttpGet("http://www.sogo.com")
}

func simpleHttpGet(url string) {
	for i := 0; i <= 100; i++ {
		logger.Debugf("Trying to hit GET request for %s", url)
		resp, err := http.Get(url)
		if err != nil {
			logger.Errorf("Error fetching URL %s : Error = %s", url, err)
		} else {
			logger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
			resp.Body.Close()
		}
	}

}
