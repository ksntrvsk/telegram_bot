package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

const (
	getUpdatesMethod   = "getUpdates"
	sendMessageMethod  = "sendMessage"
	sendPhotoMethod    = "sendPhoto"
	sendLocationMethod = "sendLocation"
)

func New(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func (client *Client) Updates(offset int, limit int) ([]Update, error) {

	query := url.Values{}
	query.Add("offset", strconv.Itoa(offset))
	query.Add("limit", strconv.Itoa(limit))

	data, err := client.doRequest(getUpdatesMethod, query)
	if err != nil {
		return nil, myErr("can't get data from request", err)
	}

	var response UpdatesResponce

	if err := json.Unmarshal(data, &response); err != nil {
		return nil, myErr("can't decode json", err)
	}

	return response.Result, nil
}

func (client *Client) SendMessage(chatID int, text string) error {

	urlRequest := url.URL{
		Scheme: "https",
		Host:   client.host,
		Path:   path.Join(client.basePath, sendMessageMethod),
	}

	body, _ := json.Marshal(map[string]string{
		"chat_id": strconv.Itoa(chatID),
		"text":    text,
	})

	response, err := http.Post(
		urlRequest.String(),
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		return myErr("can't get responce", err)
	}

	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		return myErr("can't send message", err)
	}

	return nil
}

func (client *Client) SendPhoto(chatID int, photo string) error {

	query := url.Values{}
	query.Add("chat_id", strconv.Itoa(chatID))
	query.Add("photo", photo)

	_, err := client.doRequest(sendPhotoMethod, query)
	if err != nil {
		return myErr("can't send message", err)
	}

	return nil
}

func (client *Client) SendLocation(chatID int, latitude float64, longitude float64) error {

	query := url.Values{}
	query.Add("chat_id", strconv.Itoa(chatID))
	query.Add("latitude", strconv.FormatFloat(latitude, 'E', -1, 64))
	query.Add("longitude", strconv.FormatFloat(longitude, 'E', -1, 64))

	_, err := client.doRequest(sendLocationMethod, query)
	if err != nil {
		return myErr("can't send message", err)
	}

	return nil
}

func (client *Client) doRequest(methods string, query url.Values) (data []byte, err error) {

	url := url.URL{
		Scheme: "http",
		Host:   client.host,
		Path:   path.Join(client.basePath, methods),
	}

	request, err := http.NewRequest(http.MethodGet, url.String(), nil)

	if err != nil {
		return nil, myErr("can't do request", err)
	}

	request.URL.RawQuery = query.Encode()

	responce, err := client.client.Do(request)
	if err != nil {
		return nil, myErr("can't do request", err)
	}

	defer responce.Body.Close()

	body, err := io.ReadAll(responce.Body)
	if err != nil {
		return nil, myErr("can't read body responce", err)
	}

	return body, nil
}

func newBasePath(token string) string {
	return "bot" + token
}

func myErr(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}

/*func (client *Client) SendMessage(chatID int, text string) error {

	query := url.Values{}
	query.Add("chatID", strconv.Itoa(chatID))
	query.Add("text", text)

	_, err := client.doRequest(sendMessageMethod, query)
	if err != nil {
		return myErr("can't send message", err)
	}

	return nil
}*/
