package kavenegar

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	apiBaseURL = "https://api.kavenegar.com/"
	apiVersion = "v1"
	apiFormat  = "json"
	version    = "1.0.0"
)
//Return ...
type Return struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

//ReturnError ...
type ReturnError struct {
	*Return `json:"return"`
}

//Client ...
type Client struct {
	client  *http.Client
	APIKey  string
	BaseURL *url.URL
}

//NewClient ...
func NewClient(apikey string) *Client {
	baseURL, _ := url.Parse(apiBaseURL)
	c := &Client{
		client:  http.DefaultClient,
		BaseURL: baseURL,
		APIKey:  apikey,
	}
	return c
}

//EndPoint ...
func (c *Client) EndPoint(parts ...string) *url.URL {
	up := []string{apiVersion, c.APIKey}
	up = append(up, parts...)
	u, _ := url.Parse(strings.Join(up, "/"))
	u.Path = fmt.Sprintf("/%s.%s", u.Path, apiFormat)
	return u
}

//Execute ...
func (c *Client) Execute(urlStr string, b url.Values, v interface{}) error {
	body := strings.NewReader(b.Encode())
	ul, _ := url.Parse(urlStr)
	u := c.BaseURL.ResolveReference(ul)
	req, _ := http.NewRequest("POST", u.String(), body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Charset", "utf-8")
	resp, err := c.client.Do(req)
	if err != nil {
		return &HTTPError{
			Status:  resp.StatusCode,
			Message: resp.Status,
			Err:     err,
		}
	}
	defer resp.Body.Close()
	if 200 != resp.StatusCode {
		re := new(ReturnError)
		err = json.NewDecoder(resp.Body).Decode(&re)
		if err != nil {
			return &HTTPError{
				Status:  resp.StatusCode,
				Message: resp.Status,
			}
		}
		return &APIError{
			Status:  re.Return.Status,
			Message: re.Return.Message,
		}
	}
	_ = json.NewDecoder(resp.Body).Decode(&v)
	return nil
}
