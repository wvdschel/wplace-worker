package cloudbuster

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	baseURL string
	client  httpDoer
}

type httpDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type Request struct {
	Link  string `json:"link"`
	Proxy string `json:"proxy,omitempty"`
}

type Response struct {
	Token   string `json:"token"`
	Cookies []struct {
		Name         string  `json:"name"`
		Value        string  `json:"value"`
		Domain       string  `json:"domain"`
		Path         string  `json:"path"`
		Expires      float64 `json:"expires"`
		Size         int     `json:"size"`
		HTTPOnly     bool    `json:"httpOnly"`
		Secure       bool    `json:"secure"`
		Session      bool    `json:"session"`
		SameSite     string  `json:"sameSite"`
		Priority     string  `json:"priority"`
		SameParty    bool    `json:"sameParty"`
		SourceScheme string  `json:"sourceScheme"`
		SourcePort   int     `json:"sourcePort"`
		PartitionKey struct {
			TopLevelSite         string `json:"topLevelSite"`
			HasCrossSiteAncestor bool   `json:"hasCrossSiteAncestor"`
		} `json:"partitionKey"`
	} `json:"cookies"`
}

func NewClient(baseURL string, client httpDoer) *Client {
	return &Client{
		baseURL: baseURL,
		client:  client,
	}
}

func (c *Client) GetToken(link string, proxy string) (string, []*http.Cookie, error) {
	reqPayload := Request{
		Link:  link,
		Proxy: proxy,
	}

	reqBody, err := json.Marshal(reqPayload)
	if err != nil {
		return "", nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/token", c.baseURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return "", nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", nil, fmt.Errorf("failed to get CF clearance: %d", resp.StatusCode)
	}

	respBody := Response{}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return "", nil, err
	}

	cookies := make([]*http.Cookie, len(respBody.Cookies))
	for i, c := range respBody.Cookies {
		cookies[i] = &http.Cookie{
			Name:     c.Name,
			Value:    c.Value,
			Path:     c.Path,
			Domain:   c.Domain,
			Expires:  time.Now().Add(time.Duration(c.Expires) * time.Second),
			HttpOnly: c.HTTPOnly,
			Secure:   c.Secure,
		}
	}

	return respBody.Token, cookies, nil
}
