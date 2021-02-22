package pswincom

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const baseURL = "https://simple.pswin.com"

type Client struct {
	Username   string
	Password   string
	SenderName string
}

func NewClient(username, password, senderName string) *Client {
	return &Client{
		Username:   username,
		Password:   password,
		SenderName: senderName,
	}
}

type Message struct {
	Body       string
	Message    string
	To         string
	URL        string
	Replace    bool
	Status     string
	StatusCode int
}

func parseNumber(number string) string {
	if strings.Index(number, "+") == 0 {
		return string(number[1:])
	}

	if strings.Index(number, "00") == 0 {
		return string(number[2:])
	}

	if len(number) == 8 {
		return fmt.Sprintf("47%s", number)
	}

	return number
}

func (c *Client) SendMessage(recipient, message string, replace bool) (*Message, error) {
	httpClient := &http.Client{}

	// requestBody := url.Values{}
	// requestBody.Set("USER", c.Username)
	// requestBody.Set("PW", c.Password)
	// requestBody.Set("SND", c.SenderName)
	// requestBody.Set("RCV", parseNumber(recipient))
	// requestBody.Set("TXT", message)

	// if replace {
	// 	requestBody.Set("REPLACE", "1")
	// }

	// strings.NewReader(requestBody.Encode())

	requestBody := fmt.Sprintf("USER=%s&PW=%s&SND=%s&RCV=%s&TXT=%s", c.Username, c.Password, c.SenderName, parseNumber(recipient), message)

	req, err := http.NewRequest("POST", baseURL, strings.NewReader(requestBody))
	if err != nil {
		return nil, fmt.Errorf("Could not setup request: %v", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Could not execute request: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not convert body to bytes: %v", err)
	}

	return &Message{
		Body:       string(body),
		Message:    message,
		To:         parseNumber(recipient),
		URL:        baseURL,
		Replace:    replace,
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
	}, nil
}
