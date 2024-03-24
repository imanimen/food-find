package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Call makes an HTTP request to the provided URL and method with the given
// payload. It returns the deserialized JSON response as a map[string]interface{}
// and any error. The client and request are configured before sending. The
// response body is closed after decoding.
func Call(url string, method string, payload []byte /*authToken string*/) (map[string]interface{}, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Authorization", authToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("Error in request: ", err.Error())
		}
	}(resp.Body)

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
