package poster

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Post(u string, d map[string]interface{}) (*http.Response, error) {
	jsonString, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", u, bytes.NewBuffer(jsonString))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	return resp, err
}
