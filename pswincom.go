package pswincom

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	defaultBaseURL = "https://simple.pswin.com"
	userAgent      = "simple-pswincom-go"
)

// A Client manages communication with the PSWincom Simple API.
type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	Sender     string
	Credential Credential

	client *http.Client
}

// Credential manages credentials
type Credential struct {
	Username string
	Password string
}

// NewClient returns a new PSWincom Simple API client. If a nil httpClient is
// provided, a new http.Client will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(username, password, sender string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	baseURL, _ := url.Parse(defaultBaseURL)
	credential := Credential{username, password}
	c := &Client{client: httpClient, UserAgent: userAgent, BaseURL: baseURL, Sender: sender, Credential: credential}

	return c
}

func NewClientFromEnv(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	baseURL, _ := url.Parse(defaultBaseURL)
	credential := Credential{os.Getenv("PSWINCOM_USER"), os.Getenv("PSWINCOM_PASSWD")}
	sender := os.Getenv("PSWINCOM_SENDER")

	c := &Client{
		client:     httpClient,
		UserAgent:  userAgent,
		BaseURL:    baseURL,
		Sender:     sender,
		Credential: credential,
	}

	return c
}

// EncodeBody creates a 'application/x-www-form-urlencoded' string
func (c Client) EncodeBody(message *Message) (string, error) {

	b := url.Values{}
	b.Set("USER", c.Credential.Username)
	b.Set("PW", c.Credential.Password)
	b.Set("SND", c.Sender)
	b.Set("RCV", message.Recipient)
	b.Set("TXT", message.Body)

	if message.Replace {
		b.Set("REPLACE", "1")
	}

	return url.QueryUnescape(b.Encode())
}

// SendMessage handles sending text messages
func (c *Client) SendMessage(message *Message) error {

	rb, err := c.EncodeBody(message)
	if err != nil {
		return fmt.Errorf("Could not encode body: %v", err)
	}

	req, err := http.NewRequest("POST", c.BaseURL.String(), strings.NewReader(rb))
	if err != nil {
		return fmt.Errorf("Could not setup request: %v", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("Could not execute request: %v", err)
	}

	defer resp.Body.Close()

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return fmt.Errorf("Status Code indicates failure: %d (%s)", resp.StatusCode, resp.Status)
	}

	return nil
}
