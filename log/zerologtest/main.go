package main

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger zerolog.Logger

func init() {
	writer, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return
	}
	//logger = zerolog.New(writer)
	logger = log.With().Caller().Logger().Output(writer)
	logger.Output(writer)

	logger.Level(zerolog.InfoLevel)
}

func main() {
	simpleHttpGet("www.sogo.com")
	simpleHttpGet("http://www.sogo.com")
}

func simpleHttpGet(url string) {
	for i := 0; i <= 100; i++ {
		logger.Debug().Msgf("Trying to hit GET request for %s", url)
		resp, err := http.Get(url)
		if err != nil {
			logger.Error().Msgf("Error fetching URL %s : Error = %s", url, err)
		} else {
			logger.Info().Msgf("Success! statusCode = %s for URL %s", resp.Status, url)
			resp.Body.Close()
		}
	}

}
