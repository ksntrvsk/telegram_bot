package unsplash

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
		basePath: newBasePath(),
		client:   http.Client{},
	}
}

func (client *Client) Image(token string) (string, error) {

	query := url.Values{}
	query.Add("client_id", token)

	url := url.URL{
		Scheme: "https",
		Host:   client.host,
		Path:   client.basePath,
	}

	request, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return "", myErr("can't do request", err)
	}

	request.URL.RawQuery = query.Encode()

	resp, err := client.client.Do(request)
	if err != nil {
		return "", myErr("can't do request", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", myErr("can't read body responce", err)
	}

	var response Response

	if err := json.Unmarshal(body, &response); err != nil {
		return "", myErr("can't decode json", err)
	}

	return response.URL.Regular, nil
}

func newBasePath() string {
	return "photos/random"
}

func myErr(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}
