package discordbot

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

func HttpPost(requrl string, data []byte, timeout time.Duration, headerMap map[string]string) (payload []byte, err error) {
	req, err := http.NewRequest("POST", requrl, bytes.NewReader(data))
	if err != nil {
		log.Error().Msgf("Failed to new http request: %v", err.Error())
		return
	}

	for k, v := range headerMap {
		req.Header.Set(k, v)
	}

	client := &http.Client{Timeout: timeout}

	resp, err := client.Do(req)
	if resp == nil {
		return
	}

	defer resp.Body.Close()

	payload, err = ioutil.ReadAll(resp.Body)
	return
}
