package credentials

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	taskmodels "github.com/karman-digital/clickup/models/tasks"
	timetrackingmodels "github.com/karman-digital/clickup/models/timetracking"
)

func (c *Credentials) SendTimeTrackingRequest(method, path string, body []byte, opts ...timetrackingmodels.GetOptions) (*http.Response, error) {
	req, err := retryablehttp.NewRequest(method, fmt.Sprintf("https://api.clickup.com/api/v2%s", path), body)
	if err != nil {
		return nil, err
	}
	if c.apiKey != "" {
		req.Header.Set("Authorization", c.apiKey)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Credentials) SendTaskRequest(method, path string, body []byte, opts ...taskmodels.GetTaskOptions) (*http.Response, error) {
	req, err := retryablehttp.NewRequest(method, fmt.Sprintf("https://api.clickup.com/api/v2%s", path), body)
	if err != nil {
		return nil, err
	}
	if c.apiKey != "" {
		req.Header.Set("Authorization", c.apiKey)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Credentials) SendRequest(method, path string, body []byte) (*http.Response, error) {
	req, err := retryablehttp.NewRequest(method, fmt.Sprintf("https://api.clickup.com/api/v2%s", path), body)
	if err != nil {
		return nil, err
	}
	if c.apiKey != "" {
		req.Header.Set("Authorization", c.apiKey)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
