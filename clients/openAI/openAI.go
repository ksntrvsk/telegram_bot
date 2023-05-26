package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host string) *Client {
	return &Client{
		host:     host,
		basePath: basePath(),
		client:   http.Client{},
	}
}

func (client *Client) CreateImage(prompt string, num int) ([]URL, error) {

	url := url.URL{
		Scheme: "https",
		Host:   client.host,
		Path:   client.basePath,
	}

	body, _ := json.Marshal(map[string]string{
		"prompt": prompt,
	})

	request, err := http.NewRequest(http.MethodPost, url.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, myErr("can't do request", err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "")

	resp, err := client.client.Do(request)
	if err != nil {
		return nil, myErr("can't do request", err)
	}

	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, myErr("can't read body responce", err)
	}

	var response Response

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, myErr("can't decode json", err)
	}

	return response.Data, nil
}

func basePath() string {
	return "v1/images/generations"
}

func myErr(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}
