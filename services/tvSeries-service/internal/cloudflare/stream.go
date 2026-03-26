package cloudflare

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type StreamClient struct {
	AccountID string
	APIToken  string
}

type createVideoResponse struct {
	Result struct {
		UID string `json:"uid"`
	} `json:"result"`
}

func (c *StreamClient) CreateVideo() (string, error) {
	url := "https://api.cloudflare.com/client/v4/accounts/" +
		c.AccountID + "/stream"

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte("{}")))
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var res createVideoResponse
	json.NewDecoder(resp.Body).Decode(&res)

	return res.Result.UID, nil
}
