package kavenegar

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
)

const (
	apiBaseURL = "https://api.kavenegar.com/"
	apiVersion = "v1"
	apiFormat  = "json"
	version    = "0.1.0"
)

type Return struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ReturnError struct {
	*Return `json:"return"`
}

type Client struct {
	BaseClient *http.Client
	apikey     string
	BaseURL    *url.URL
}

func NewClient(apikey string) *Client {
	baseURL, _ := url.Parse(apiBaseURL)
	c := &Client{
		BaseClient: http.DefaultClient,
		BaseURL:    baseURL,
		apikey:     apikey,
	}
	return c
}

func (c *Client) EndPoint(parts ...string) *url.URL {
	up := []string{apiVersion, c.apikey}
	up = append(up, parts...)
	u, _ := url.Parse(strings.Join(up, "/"))
	u.Path = fmt.Sprintf("/%s.%s", u.Path, apiFormat)
	return u
}

func (c *Client) Execute(urlStr string, b url.Values, v interface{}) error {
	body := strings.NewReader(b.Encode())
	ul, _ := url.Parse(urlStr)
	u := c.BaseURL.ResolveReference(ul)
	req, _ := http.NewRequest("POST", u.String(), body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Charset", "utf-8")
	resp, err := c.BaseClient.Do(req)
	if err != nil {
		if err, ok := err.(net.Error); ok {
			return err
		}
		if resp == nil {
			return &HTTPError{
				Status:  http.StatusInternalServerError,
				Message: "nil api response",
				Err:     err,
			}
		}
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
