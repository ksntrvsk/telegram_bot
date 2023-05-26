package advice

import (
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

func (client *Client) Advice() (string, error) {
	url := url.URL{
		Scheme: "https",
		Host:   client.host,
		Path:   client.basePath,
	}

	request, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return "", myErr("can't do request", err)
	}

	resp, err := client.client.Do(request)
	if err != nil {
		return "", myErr("can't do request", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", myErr("can't read body responce", err)
	}

	var response Slip

	if err := json.Unmarshal(body, &response); err != nil {
		return "", myErr("can't decode json", err)
	}

	return response.Advice.Advice, nil
}

func basePath() string {
	return "advice"
}

func myErr(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}
