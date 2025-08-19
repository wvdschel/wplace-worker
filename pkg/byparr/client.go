package byparr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
}

type Option func(*Client) error

func WithHttpClient(h *http.Client) Option {
	return func(c *Client) error {
		c.httpClient = h
		return nil
	}
}

func WithBaseURL(u string) Option {
	return func(c *Client) error {
		c.baseURL = u
		return nil
	}
}

func New(opts ...Option) (*Client, error) {
	c := &Client{
		httpClient: http.DefaultClient,
		baseURL:    "http://localhost:8191/v1",
	}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) GetCookie(ctx context.Context, url string, maxDuration time.Duration) ([]*http.Cookie, error) {
	reqInfo := Request{
		Cmd:        CmdRequestGet,
		URL:        url,
		MaxTimeout: int(maxDuration.Seconds()),
	}

	data, err := json.Marshal(reqInfo)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to send request: %s", resp.Status)
	}

	responseData := &Response{}
	if err := json.NewDecoder(resp.Body).Decode(responseData); err != nil {
		return nil, err
	}

	if responseData.Status != StatusOK || responseData.Message != MessageSuccess {
		return nil, fmt.Errorf("failed get session cookie: %s", responseData.Message)
	}

	res := make([]*http.Cookie, len(responseData.Solution.Cookies))
	for i, cookie := range responseData.Solution.Cookies {
		res[i] = &http.Cookie{
			Name:     cookie.Name,
			Value:    cookie.Value,
			Path:     cookie.Path,
			Domain:   cookie.Domain,
			Expires:  time.Unix(cookie.Expires, 0),
			Secure:   cookie.Secure,
			HttpOnly: cookie.HTTPOnly,
		}
	}

	return res, nil
}
