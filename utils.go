package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var httpClientGithub = &http.Client{Timeout: 10 * time.Second}

func getJSONFromHTTPRequest(url string, target interface{}) error {
	r, err := httpClientGithub.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
